<script lang="ts">
	import { Tile } from 'carbon-components-svelte';
	import Picker from './picker.svelte';
</script>

<div class="m-3 w-full flex-row items-center justify-center sm:flex">
	<h1 class="text-xl">CS2 Util Roulette ▄︻┻┳═一</h1>
</div>

<div class="w-full flex-row items-center justify-center sm:flex">
	<Tile class="m-2 ">
		<p><em>This guy keeps fucking flashing me!</em> - The other guys</p>
	</Tile>
</div>
<div class="w-full flex-row items-center justify-center sm:flex">
	<Tile class="m-14 ">
		<Picker />
	</Tile>
</div>
