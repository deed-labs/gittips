import type { InstallationInfo } from '$lib/types';
import { writable } from 'svelte/store';

export interface LocalStorageData extends InstallationInfo {
	wallet_address: string;
	bot_installation_done: boolean;
}

export const defaultData: LocalStorageData = {
	wallet_address: '',
	bot_installation_done: false,
	installed: false,
	name: '',
	ownerId: 0
};

export const storage = writable<LocalStorageData>(
	localStorage['data'] ? JSON.parse(localStorage['data']) : defaultData
);

storage.subscribe((value) => localStorage.setItem('data', JSON.stringify(value)));
