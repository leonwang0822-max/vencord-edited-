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

import { showNotification } from "@api/Notifications";
import { Settings } from "@api/Settings";
import { openUpdaterModal } from "@components/settings/tabs/updater";

import { relaunch } from "./native";
import { UpdateInfo } from "./versionChecker";

export interface NotificationOptions {
    title: string;
    body: string;
    color?: string;
    onClick?: () => void;
    permanent?: boolean;
    icon?: string;
    actions?: Array<{
        label: string;
        action: () => void;
    }>;
}

export class UpdateNotificationManager {
    private static instance: UpdateNotificationManager;
    private notificationHistory: Array<{ timestamp: number; type: string; message: string }> = [];
    private activeNotifications = new Set<string>();

    static getInstance(): UpdateNotificationManager {
        if (!UpdateNotificationManager.instance) {
            UpdateNotificationManager.instance = new UpdateNotificationManager();
        }
        return UpdateNotificationManager.instance;
    }

    /**
     * Show update available notification
     */
    showUpdateAvailable(updateInfo: UpdateInfo): void {
        if (this.activeNotifications.has("update-available")) return;

        const changeCount = updateInfo.changes.length;
        const latestChange = updateInfo.changes[0];

        let body = `${changeCount} new update${changeCount > 1 ? "s" : ""} available`;
        if (latestChange) {
            const shortMessage = latestChange.message.length > 50
                ? latestChange.message.substring(0, 50) + "..."
                : latestChange.message;
            body += `\nLatest: ${shortMessage}`;
        }

        this.showNotification({
            title: "🔄 Vencord Update Available!",
            body,
            color: "var(--yellow-360)",
            onClick: () => {
                this.activeNotifications.delete("update-available");
                openUpdaterModal?.();
            },
            permanent: true
        });

        this.activeNotifications.add("update-available");
        this.logNotification("update-available", body);
    }

    /**
     * Show update downloading notification
     */
    showUpdateDownloading(): void {
        this.showNotification({
            title: "⬇️ Downloading Update",
            body: "Vencord is downloading the latest update...",
            color: "var(--blue-360)",
            permanent: false
        });

        this.logNotification("update-downloading", "Update download started");
    }

    /**
     * Show update installing notification
     */
    showUpdateInstalling(): void {
        this.showNotification({
            title: "⚙️ Installing Update",
            body: "Vencord is installing the update. Please wait...",
            color: "var(--blue-360)",
            permanent: false
        });

        this.logNotification("update-installing", "Update installation started");
    }

    /**
     * Show update completed notification
     */
    showUpdateCompleted(silent = false): void {
        if (silent && Settings.autoUpdateSilent) {
            // Show brief notification before auto-restart
            this.showNotification({
                title: "✅ Update Complete",
                body: "Vencord updated successfully. Restarting in 3 seconds...",
                color: "var(--green-360)",
                permanent: false
            });
        } else {
            this.showNotification({
                title: "✅ Vencord Updated!",
                body: "Update installed successfully. Click to restart Discord.",
                color: "var(--green-360)",
                onClick: relaunch,
                permanent: true
            });
        }

        this.activeNotifications.delete("update-available");
        this.logNotification("update-completed", "Update completed successfully");
    }

    /**
     * Show update failed notification
     */
    showUpdateFailed(error?: string): void {
        const errorMessage = error ? ` Error: ${error}` : "";

        this.showNotification({
            title: "❌ Update Failed",
            body: `Failed to update Vencord.${errorMessage} Click to try manual update.`,
            color: "var(--red-360)",
            onClick: () => {
                this.activeNotifications.delete("update-available");
                openUpdaterModal?.();
            },
            permanent: true
        });

        this.logNotification("update-failed", `Update failed: ${error || "Unknown error"}`);
    }

    /**
     * Show update check failed notification
     */
    showUpdateCheckFailed(error?: string): void {
        if (!Settings.autoUpdateNotification) return; // Only show if notifications are enabled

        this.showNotification({
            title: "⚠️ Update Check Failed",
            body: "Failed to check for updates. Will retry later.",
            color: "var(--orange-360)",
            permanent: false
        });

        this.logNotification("update-check-failed", `Update check failed: ${error || "Unknown error"}`);
    }

    /**
     * Show no updates available notification (only when manually triggered)
     */
    showNoUpdatesAvailable(): void {
        this.showNotification({
            title: "✅ Up to Date",
            body: "You're running the latest version of Vencord!",
            color: "var(--green-360)",
            permanent: false
        });

        this.logNotification("no-updates", "No updates available");
    }

    /**
     * Show update progress notification
     */
    showUpdateProgress(progress: number, stage: string): void {
        const percentage = Math.round(progress * 100);

        this.showNotification({
            title: "🔄 Updating Vencord",
            body: `${stage}... ${percentage}%`,
            color: "var(--blue-360)",
            permanent: false
        });
    }

    /**
     * Show custom notification with enhanced options
     */
    private showNotification(options: NotificationOptions): void {
        showNotification({
            title: options.title,
            body: options.body,
            color: options.color || "var(--brand-500)",
            onClick: options.onClick,
            permanent: options.permanent || false,
            noPersist: false
        });
    }

    /**
     * Log notification for debugging and history
     */
    private logNotification(type: string, message: string): void {
        this.notificationHistory.push({
            timestamp: Date.now(),
            type,
            message
        });

        // Keep only last 50 notifications
        if (this.notificationHistory.length > 50) {
            this.notificationHistory = this.notificationHistory.slice(-50);
        }
    }

    /**
     * Clear active notification flags
     */
    clearActiveNotifications(): void {
        this.activeNotifications.clear();
    }

    /**
     * Get notification history for debugging
     */
    getNotificationHistory(): Array<{ timestamp: number; type: string; message: string }> {
        return [...this.notificationHistory];
    }

    /**
     * Check if a specific notification type is active
     */
    isNotificationActive(type: string): boolean {
        return this.activeNotifications.has(type);
    }

    /**
     * Format update info for display
     */
    formatUpdateInfo(updateInfo: UpdateInfo): string {
        const changes = updateInfo.changes.slice(0, 3); // Show first 3 changes
        const changeList = changes.map(change => {
            const shortHash = change.hash.substring(0, 7);
            const shortMessage = change.message.length > 40
                ? change.message.substring(0, 40) + "..."
                : change.message;
            return `• [${shortHash}] ${shortMessage}`;
        }).join("\n");

        const moreChanges = updateInfo.changes.length > 3
            ? `\n... and ${updateInfo.changes.length - 3} more changes`
            : "";

        return `Recent changes:\n${changeList}${moreChanges}`;
    }
}

// Export singleton instance
export const updateNotifications = UpdateNotificationManager.getInstance();
