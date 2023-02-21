import { TON } from './lib/stores';
import { TonKeeper } from './lib/pkg/wallet/ton';

TON.init({
	TonKeeper: new TonKeeper()
});
