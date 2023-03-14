import { writable } from 'svelte/store';

interface LocalStorageData {
	wallet_address: string;
	bot_installation_done: boolean;
}

export const storage = writable<LocalStorageData>(JSON.parse(localStorage.getItem('data') ?? '{}'));

storage.subscribe((value) => localStorage.setItem('data', JSON.stringify(value)));
