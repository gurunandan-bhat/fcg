import * as params from '@params';

let url = window.location.href;
const mch = url.match(/page\/(\d+)\/$/);

console.log(mch);
console.log(params.moveTo);

if (mch && mch[1] !== 1) {
	let anch = document.getElementById(params.moveTo);
	console.log(anch);
	window.addEventListener('load', function (e) {
		anch.scrollIntoView();
	});
}
