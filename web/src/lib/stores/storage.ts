import type { InstallationInfo } from '$lib/types';
import { writable } from 'svelte/store';

interface LocalStorageData {
	wallet_address: string;
	owner: InstallationInfo;
	bot_installation_done: boolean;
}

export const defaultData: LocalStorageData = {
	wallet_address: '',
	bot_installation_done: false,
	owner: { installed: false, name: '', id: 0 }
};

export const storage = writable<LocalStorageData>(
	localStorage['data'] ? JSON.parse(localStorage['data']) : defaultData
);

storage.subscribe((value) => localStorage.setItem('data', JSON.stringify(value)));
