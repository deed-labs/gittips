import type { IWallet } from '../wallet';
import TonConnectWallet from './tonconnect';

export default class TonKeeper extends TonConnectWallet implements IWallet {
	readonly available: boolean;

	constructor() {
		super(0);
		this.available = true;
	}
}
