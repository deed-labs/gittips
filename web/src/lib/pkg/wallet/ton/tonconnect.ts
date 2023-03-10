import TonConnect, { type Wallet, type WalletInfoRemote } from '@tonconnect/sdk';
import { Address } from 'ton';

export default class TonConnectWallet {
	index: number;
	connected: boolean;
	address: string;

	readonly connector: TonConnect;

	constructor(walletIndex: number) {
		this.index = walletIndex;
		this.connected = false;
		this.address = '';

		// TODO: update manifest link when repo is public
		this.connector = new TonConnect({
			manifestUrl:
				'https://raw.githubusercontent.com/deed-labs/gittips/master/web/tonconnect-manifest.json'
		});
	}

	async connectExternal(cb: (address: string) => void): Promise<string> {
		const walletsList = await TonConnect.getWallets();

		const walletConnectionSource = {
			universalLink: (walletsList[this.index] as WalletInfoRemote).universalLink,
			bridgeUrl: (walletsList[this.index] as WalletInfoRemote).bridgeUrl
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
