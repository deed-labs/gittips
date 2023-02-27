import TonConnect, { type Wallet, type WalletInfoRemote } from '@tonconnect/sdk';
import { Address } from 'ton';
import type { IWallet } from '../wallet';

/**
 * Implements Wallet abstract class for TonKeeper.
 *
 * @class TonKeeper
 */
export default class TonKeeper implements IWallet {
	connected: boolean;
	address: string;

	readonly available: boolean;
	readonly connector: TonConnect;

	constructor() {
		this.available = true;
		this.connected = false;
		this.address = '';

		// TODO: update manifest link when repo is public
		this.connector = new TonConnect({
			manifestUrl:
				'https://raw.githubusercontent.com/bifrost-defi/bifrost/main/tonconnect-manifest.json'
		});
	}

	async connectExternal(cb: (address: string) => void): Promise<string> {
		const walletsList = await TonConnect.getWallets();

		const walletConnectionSource = {
			universalLink: (walletsList[0] as WalletInfoRemote).universalLink,
			bridgeUrl: (walletsList[0] as WalletInfoRemote).bridgeUrl
		};

		this.connector.onStatusChange((wallet: Wallet | null) => {
			if (this.connector.connected && wallet) {
				this.address = Address.parseRaw(wallet.account.address).toString();
				cb(this.address);
			}
		}, console.error);

		return this.connector.connect(walletConnectionSource);
	}
}
