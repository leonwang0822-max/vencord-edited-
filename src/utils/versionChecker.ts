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

import gitHash from "~git-hash";

import { Logger } from "./Logger";

const logger = new Logger("VersionChecker", "cyan");

export interface VersionInfo {
    hash: string;
    author: string;
    message: string;
    timestamp?: number;
    tag?: string;
}

export interface UpdateInfo {
    available: boolean;
    isNewer: boolean;
    currentVersion: string;
    latestVersion?: string;
    changes: VersionInfo[];
    updateSize?: number;
    releaseNotes?: string;
}

export class VersionChecker {
    private static instance: VersionChecker;
    private cachedUpdateInfo: UpdateInfo | null = null;
    private lastCheckTime = 0;
    private readonly cacheTimeout = 5 * 60 * 1000; // 5 minutes

    static getInstance(): VersionChecker {
        if (!VersionChecker.instance) {
            VersionChecker.instance = new VersionChecker();
        }
        return VersionChecker.instance;
    }

    /**
     * Compare two version strings or commit hashes
     */
    private compareVersions(current: string, latest: string): number {
        // If both are semantic versions (x.y.z format)
        const semverRegex = /^(\d+)\.(\d+)\.(\d+)(?:-([\w.-]+))?$/;
        const currentSemver = current.match(semverRegex);
        const latestSemver = latest.match(semverRegex);

        if (currentSemver && latestSemver) {
            const [, currentMajor, currentMinor, currentPatch] = currentSemver.map(Number);
            const [, latestMajor, latestMinor, latestPatch] = latestSemver.map(Number);

            if (currentMajor !== latestMajor) return latestMajor - currentMajor;
            if (currentMinor !== latestMinor) return latestMinor - currentMinor;
            if (currentPatch !== latestPatch) return latestPatch - currentPatch;
            return 0;
        }

        // For commit hashes, we rely on the order from git
        return current === latest ? 0 : 1;
    }

    /**
     * Get detailed update information
     */
    async getUpdateInfo(forceRefresh = false): Promise<UpdateInfo> {
        const now = Date.now();

        // Return cached result if still valid
        if (!forceRefresh && this.cachedUpdateInfo && (now - this.lastCheckTime) < this.cacheTimeout) {
            return this.cachedUpdateInfo;
        }

        try {
            logger.info("Checking for version updates...");

            // Get updates from the native updater
            const changes: VersionInfo[] = await new Promise((resolve, reject) => {
                VencordNative.updater.getUpdates()
                    .then(result => {
                        if (result.ok) {
                            resolve(result.value);
                        } else {
                            reject(result.error);
                        }
                    })
                    .catch(reject);
            });

            const currentVersion = gitHash;
            let isNewer = false;
            let available = changes.length > 0;

            // Check if current version is newer than remote (for dev builds)
            if (!IS_STANDALONE && changes.length > 0) {
                const hasCurrentCommit = changes.some(c => c.hash === currentVersion);
                if (hasCurrentCommit) {
                    isNewer = true;
                    available = false;
                }
            }

            // Get latest version info
            const latestVersion = changes.length > 0 ? changes[0].hash : currentVersion;

            // Generate release notes from changes
            const releaseNotes = this.generateReleaseNotes(changes);

            const updateInfo: UpdateInfo = {
                available,
                isNewer,
                currentVersion,
                latestVersion,
                changes,
                releaseNotes
            };

            // Cache the result
            this.cachedUpdateInfo = updateInfo;
            this.lastCheckTime = now;

            logger.info(`Update check complete: ${available ? "Updates available" : "Up to date"}`);
            return updateInfo;

        } catch (error) {
            logger.error("Failed to get update info:", error);

            // Return a safe fallback
            const fallbackInfo: UpdateInfo = {
                available: false,
                isNewer: false,
                currentVersion: gitHash,
                changes: []
            };

            return fallbackInfo;
        }
    }

    /**
     * Generate formatted release notes from changes
     */
    private generateReleaseNotes(changes: VersionInfo[]): string {
        if (changes.length === 0) return "No changes available.";

        const notes = changes.slice(0, 10).map((change, index) => {
            const shortHash = change.hash.substring(0, 7);
            const author = change.author || "Unknown";
            const message = change.message || "No commit message";

            return `${index + 1}. [${shortHash}] ${message} (by ${author})`;
        }).join("\n");

        const moreChanges = changes.length > 10 ? `\n\n... and ${changes.length - 10} more changes` : "";

        return `Recent Changes:\n${notes}${moreChanges}`;
    }

    /**
     * Check if updates are available (simple boolean check)
     */
    async hasUpdates(forceRefresh = false): Promise<boolean> {
        const info = await this.getUpdateInfo(forceRefresh);
        return info.available;
    }

    /**
     * Get the number of available updates
     */
    async getUpdateCount(forceRefresh = false): Promise<number> {
        const info = await this.getUpdateInfo(forceRefresh);
        return info.changes.length;
    }

    /**
     * Clear cached update information
     */
    clearCache(): void {
        this.cachedUpdateInfo = null;
        this.lastCheckTime = 0;
    }

    /**
     * Get formatted version string for display
     */
    getVersionString(): string {
        const hash = gitHash.substring(0, 7);
        return `${hash}${IS_DEV ? " (dev)" : ""}${IS_STANDALONE ? " (standalone)" : ""}`;
    }
}

// Export singleton instance
export const versionChecker = VersionChecker.getInstance();
