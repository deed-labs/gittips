import type { IWallet } from '../wallet';
import TonConnectWallet from './tonconnect';

class TonKeeper extends TonConnectWallet implements IWallet {
	constructor() {
		super(0);
	}
}

class OpenMask extends TonConnectWallet implements IWallet {
	constructor() {
		super(1);
	}
}

class MyTonWallet extends TonConnectWallet implements IWallet {
	constructor() {
		super(2);
	}
}

class TonSafe extends TonConnectWallet implements IWallet {
	constructor() {
		super(3);
	}
}

export { TonKeeper, OpenMask, MyTonWallet, TonSafe };
