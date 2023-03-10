import type { IWallet } from '../wallet';
import TonConnectWallet from './tonconnect';

class TonKeeper extends TonConnectWallet implements IWallet {
	readonly available: boolean;

	constructor() {
		super(0);
		this.available = true;
	}
}

class OpenMask extends TonConnectWallet implements IWallet {
	readonly available: boolean;

	constructor() {
		super(1);
		this.available = true;
	}
}

class MyTonWallet extends TonConnectWallet implements IWallet {
	readonly available: boolean;

	constructor() {
		super(2);
		this.available = true;
	}
}

class TonSafe extends TonConnectWallet implements IWallet {
	readonly available: boolean;

	constructor() {
		super(3);
		this.available = true;
	}
}

export { TonKeeper, OpenMask, MyTonWallet, TonSafe };
