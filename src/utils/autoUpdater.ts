/*
 * Vencord, a modification for Discord's desktop app
 * Copyright (c) 2022 Vendicated and contributors
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import { Settings } from "@api/Settings";

import { relaunch } from "./native";
import { updateNotifications } from "./updateNotifications";
import { update, UpdateLogger } from "./updater";
import { versionChecker } from "./versionChecker";

export class AutoUpdaterService {
    private updateCheckInterval: NodeJS.Timeout | null = null;
    private isChecking = false;
    private lastCheckTime = 0;
    private notifiedThisSession = false;

    constructor() {
        this.init();
    }

    private init() {
        if (IS_WEB || IS_UPDATER_DISABLED) return;

        // Check for updates on startup if enabled
        if (Settings.autoUpdateCheckOnStartup) {
            setTimeout(() => this.performUpdateCheck(), 5000); // Wait 5 seconds after startup
        }

        // Set up periodic update checking
        this.setupPeriodicChecking();
    }

    private setupPeriodicChecking() {
        if (this.updateCheckInterval) {
            clearInterval(this.updateCheckInterval);
        }

        if (Settings.autoUpdate || Settings.autoUpdateCheckOnStartup) {
            const intervalMs = Settings.autoUpdateInterval * 60 * 1000; // Convert minutes to milliseconds
            this.updateCheckInterval = setInterval(() => {
                this.performUpdateCheck();
            }, intervalMs);

            UpdateLogger.info(`Auto-update checking enabled with ${Settings.autoUpdateInterval} minute interval`);
        }
    }

    public async performUpdateCheck(force = false): Promise<boolean> {
        if (this.isChecking && !force) return false;

        const now = Date.now();
        const minInterval = 5 * 60 * 1000; // Minimum 5 minutes between checks

        if (!force && (now - this.lastCheckTime) < minInterval) {
            UpdateLogger.debug("Skipping update check - too soon since last check");
            return false;
        }

        this.isChecking = true;
        this.lastCheckTime = now;

        try {
            UpdateLogger.info("Checking for updates...");
            const updateInfo = await versionChecker.getUpdateInfo(force);

            if (!updateInfo.available) {
                UpdateLogger.info("No updates available");
                return false;
            }

            if (updateInfo.isNewer) {
                UpdateLogger.info("Local version is newer than remote");
                return false;
            }

            UpdateLogger.info(`Update available! ${updateInfo.changes.length} new commits`);

            if (Settings.autoUpdate) {
                return await this.performAutoUpdate(updateInfo);
            } else {
                updateNotifications.showUpdateAvailable(updateInfo);
                return true;
            }
        } catch (error) {
            UpdateLogger.error("Failed to check for updates:", error);
            const errorMessage = error instanceof Error ? error.message : String(error);
            updateNotifications.showUpdateCheckFailed(errorMessage);
            return false;
        } finally {
            this.isChecking = false;
        }
    }

    private async performAutoUpdate(updateInfo?: any): Promise<boolean> {
        try {
            UpdateLogger.info("Performing automatic update...");

            // Show downloading notification
            if (!Settings.autoUpdateSilent) {
                updateNotifications.showUpdateDownloading();
            }

            // Show installing notification
            setTimeout(() => {
                if (!Settings.autoUpdateSilent) {
                    updateNotifications.showUpdateInstalling();
                }
            }, 1000);

            const success = await update();

            if (success) {
                // Show completion notification
                updateNotifications.showUpdateCompleted(Settings.autoUpdateSilent);

                // Auto-restart if silent mode is enabled
                if (Settings.autoUpdateSilent) {
                    setTimeout(relaunch, 3000); // Wait 3 seconds then restart
                }

                UpdateLogger.info("Auto-update completed successfully");
                return true;
            } else {
                throw new Error("Update failed");
            }
        } catch (error) {
            UpdateLogger.error("Auto-update failed:", error);
            const errorMessage = error instanceof Error ? error.message : String(error);
            updateNotifications.showUpdateFailed(errorMessage);
            return false;
        }
    }

    // Removed - now handled by updateNotifications.showUpdateAvailable()

    public updateSettings() {
        this.setupPeriodicChecking();
        this.notifiedThisSession = false; // Reset notification flag when settings change
    }

    public async forceCheck(): Promise<boolean> {
        const result = await this.performUpdateCheck(true);

        // Show "no updates" notification if manually triggered and no updates found
        if (!result) {
            const updateInfo = await versionChecker.getUpdateInfo(true);
            if (!updateInfo.available && !updateInfo.isNewer) {
                updateNotifications.showNoUpdatesAvailable();
            }
        }

        return result;
    }

    public getStatus() {
        return {
            isChecking: this.isChecking,
            lastCheckTime: this.lastCheckTime,
            intervalMs: Settings.autoUpdateInterval * 60 * 1000,
            nextCheckTime: this.lastCheckTime + (Settings.autoUpdateInterval * 60 * 1000)
        };
    }

    public destroy() {
        if (this.updateCheckInterval) {
            clearInterval(this.updateCheckInterval);
            this.updateCheckInterval = null;
        }
    }
}

// Global instance
export const autoUpdater = new AutoUpdaterService();
