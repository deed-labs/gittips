import TonConnect, {
	isWalletInfoInjected,
	type Wallet,
	type WalletInfo,
	type WalletInfoInjected,
	type WalletInfoRemote
} from '@tonconnect/sdk';
import { Address } from 'ton';

export default class TonConnectWallet {
	index: number;
	connected: boolean;
	address: string;
	available: boolean = false;

	injected: boolean = false;
	embedded: boolean = false;

	info!: WalletInfo;

	ready: Promise<void>;

	readonly connector: TonConnect;

	constructor(walletIndex: number) {
		this.index = walletIndex;
		this.connected = false;
		this.address = '';

		this.connector = new TonConnect({
			manifestUrl:
				'https://raw.githubusercontent.com/deed-labs/gittips/master/web/tonconnect-manifest.json'
		});

		this.ready = new Promise((resolve) => {
			TonConnect.getWallets().then((list) => {
				this.info = list[walletIndex];

				this.available =
					(isWalletInfoInjected(this.info) && this.info.injected) ||
					((this.info as WalletInfoRemote).bridgeUrl !== undefined &&
						(this.info as WalletInfoRemote).bridgeUrl !== '');

				this.injected = isWalletInfoInjected(this.info) && this.info.injected;
				this.embedded = isWalletInfoInjected(this.info) && this.info.embedded;

				resolve();
			});
		});
	}

	async connectExternal(cb: (address: string) => void): Promise<string> {
		const walletConnectionSource = {
			universalLink: (this.info as WalletInfoRemote).universalLink,
			bridgeUrl: (this.info as WalletInfoRemote).bridgeUrl
		};

		this.connector.onStatusChange((wallet: Wallet | null) => {
			if (this.connector.connected && wallet) {
				this.address = Address.parseRaw(wallet.account.address).toString();
				cb(this.address);
			}
		}, console.error);

		return this.connector.connect(walletConnectionSource);
	}

	async connectInjected(cb: (address: string) => void): Promise<void> {
		if (!this.injected) {
			throw new Error('Wallet is not injected');
		}

		const walletConnectionSource = {
			jsBridgeKey: (this.info as WalletInfoInjected).jsBridgeKey
		};

		this.connector.onStatusChange((wallet: Wallet | null) => {
			if (this.connector.connected && wallet) {
				this.address = Address.parseRaw(wallet.account.address).toString();
				cb(this.address);
			}
		}, console.error);

		this.connector.connect(walletConnectionSource);
	}

	async disconnect(): Promise<void> {
		if (!this.connected) return;

		await this.connector.disconnect();
	}
}
