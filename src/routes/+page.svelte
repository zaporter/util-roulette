<script lang="ts">
	import { Tile } from 'carbon-components-svelte';
	import Picker from './picker.svelte';
</script>

<div class="m-3 w-full flex-row items-center justify-center sm:flex">
	<h1 class="text-xl">CS2 Util Roulette ▄︻┻┳═一</h1>
</div>

<div class="w-full flex-row items-center justify-center sm:flex">
	<Tile class="m-3 -rotate-1">
		<p><em>This guy keeps fucking flashing me!</em> <br />- The other guys</p>
	</Tile>
</div>
<div class="w-full flex-row items-center justify-center sm:flex">
	<Tile class="m-10 ">
		<Picker />
	</Tile>
</div>

<style>
	@keyframes animatedBackground {
		from {
			background-position: 98% 0;
			background-size: 50%;
		}
		to {
			background-position: 100% 0;
			background-size: 51%;
		}
	}
	:global(body) {
		background-image: url('./molly.webp');
		background-position: 0px 0px;
		animation: animatedBackground 1s linear infinite alternate;
	}
</style>
