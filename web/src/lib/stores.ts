import { writable, get, derived, type Readable } from 'svelte/store';
import type { IWallet } from './pkg/wallet/wallet';

type Wallets = {
	[name: string]: IWallet;
};

type WalletInfo = {
	address: string;
	connected: boolean;
};

type WalletStore = Readable<WalletInfo> & {
	connectExternal: (onConnected?: () => void) => Promise<string>;
	disconnect: () => void;
	available: boolean;
};

type WalletStores = {
	[name: string]: WalletStore;
};

type Network = {
	connected: Readable<boolean>;
	address: Readable<string>;
	wallets: WalletStores;
};

const makeWalletStore = (wallet: IWallet): WalletStore => {
	const { subscribe, set } = writable({} as WalletInfo);

	const connectExternal = async (onConnected?: () => void) => {
		let link = await wallet.connectExternal((address: string) => {
			set({
				address,
				connected: true
			});

			if (onConnected) onConnected();
		});

		return link;
	};

	const disconnect = () => {
		set({
			address: '',
			connected: false
		});
	};

	return {
		subscribe,
		connectExternal,
		disconnect,
		available: wallet.available
	};
};

const makeNetworkStore = () => {
	let initialValue = {} as Network;
	const { subscribe, set } = writable(initialValue);

	return {
		subscribe,
		init(wallets: Wallets) {
			let walletStores: { [name: string]: WalletStore } = {};
			for (let [key, value] of Object.entries(wallets)) {
				walletStores[key] = makeWalletStore(value);
			}

			let connected = derived(Object.values(walletStores), (wallets) => {
				return wallets.some((wallet) => wallet.connected);
			});

			let address = derived(Object.values(walletStores), (wallets) => {
				let addresses = wallets.map((wallet) => wallet.address);
				return addresses.filter((address) => address !== '')[0];
			});

			set({ connected, address, wallets: walletStores });
		},
		disconnect() {
			let network = get(this);
			for (let wallet of Object.values(network.wallets)) {
				wallet.disconnect();
			}
		}
	};
};

export const TON = makeNetworkStore();
