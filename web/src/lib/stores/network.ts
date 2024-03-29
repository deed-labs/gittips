import type { Message } from '$lib/pkg/txs/ton';
import { writable, get, derived, type Readable } from 'svelte/store';
import type { IWallet } from '../pkg/wallet/wallet';

type Wallets = {
	[name: string]: IWallet;
};

type WalletInfo = {
	address: string;
	connected: boolean;
};

export type WalletStore = Readable<WalletInfo> & {
	connectExternal: (onConnected?: () => void) => Promise<string>;
	connectInjected: (onConnected?: () => void) => Promise<void>;
	sendTransaction: (...messages: Message[]) => Promise<void>;
	disconnect: () => Promise<void>;
	available: boolean;
	supported: boolean;

	injected: boolean;
	embedded: boolean;
};

type WalletStores = {
	[name: string]: WalletStore;
};

export type Network = {
	connected: Readable<boolean>;
	address: Readable<string>;
	wallets: WalletStores;
};

const makeWalletStore = (wallet: IWallet): WalletStore => {
	const { subscribe, set } = writable({ address: wallet.address, connected: wallet.connected });

	const connectExternal = async (onConnected?: () => void): Promise<string> => {
		let link = await wallet.connectExternal((address: string) => {
			set({
				address,
				connected: true
			});

			if (onConnected) onConnected();
		});

		return link;
	};

	const connectInjected = async (onConnected?: () => void): Promise<void> => {
		await wallet.connectInjected((address: string) => {
			set({
				address,
				connected: true
			});

			if (onConnected) onConnected();
		});
	};

	const sendTransaction = async (...messages: Message[]) =>
		await wallet.sendTransaction(...messages);

	const disconnect = async () => {
		await wallet.disconnect();
		set({
			address: '',
			connected: false
		});
	};

	return {
		subscribe,
		connectExternal,
		connectInjected,
		sendTransaction,
		disconnect,
		available: wallet.available,
		supported: wallet.supported,
		injected: wallet.injected,
		embedded: wallet.embedded
	};
};

const makeNetworkStore = () => {
	let initialValue = { wallets: {} } as Network;
	const { subscribe, set } = writable(initialValue);

	return {
		subscribe,
		async init(wallets: Wallets) {
			let walletStores: { [name: string]: WalletStore } = {};
			for (let [name, wallet] of Object.entries(wallets)) {
				await wallet.ready;
				walletStores[name] = makeWalletStore(wallet);
			}

			let connected = derived(Object.values(walletStores), (wallets) => {
				return wallets.some((wallet) => wallet.connected);
			});

			let address = derived(Object.values(walletStores), (wallets) => {
				let addresses = wallets.map((wallet) => wallet.address);
				let address = addresses.filter((address) => address && address !== '')[0];

				return address;
			});

			set({ connected, address, wallets: walletStores });
		},
		getConnectedWallet() {
			let network = get(this);
			let wallets = Object.values(network.wallets).filter((w) => get(w).connected);

			return wallets.length === 0 ? null : wallets[0];
		},
		async disconnect() {
			let network = get(this);
			for (let wallet of Object.values(network.wallets)) {
				await wallet.disconnect();
			}
		}
	};
};

export const TON = makeNetworkStore();
