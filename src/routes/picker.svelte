<script lang="ts">
	import { ComboBox, Button } from 'carbon-components-svelte';
	import mirage from './mirage.json';
	import inferno from './inferno.json';
	import overpass from './overpass.json';
	import ancient from './ancient.json';
	import nuke from './nuke.json';
	import vertigo from './vertigo.json';
	import anubis from './anubis.json';
	import dust from './dust2.json';
	interface MapEntry {
		id: string;
		text: string;
		vids: MapInfo;
	}
	const maps: MapEntry[] = [
		// @ts-ignore I dont care to type this
		{ id: '0', text: 'Mirage', vids: mirage },
		{ id: '1', text: 'Inferno', vids: inferno },
		{ id: '2', text: 'Overpass', vids: overpass },
		{ id: '3', text: 'Ancient', vids: ancient },
		{ id: '4', text: 'Nuke', vids: nuke },
		{ id: '5', text: 'Vertigo', vids: vertigo },
		{ id: '6', text: 'Anubis', vids: anubis },
		{ id: '7', text: 'Dust2', vids: dust }
	];

	interface Entry {
		id: string;
		text: string;
	}
	const sides: Entry[] = [
		{ id: '0', text: 'T' },
		{ id: '1', text: 'CT' },
		{ id: '2', text: 'Both' }
	];
	let selectedMap: string | undefined = undefined;
	let selectedSide: string | undefined = undefined;
	interface GeneratedInfo {
		youTubeId: string;
		startPosition: string;
		endPosition: string;
		type: string;
	}
	interface MapInfo {
		T: GeneratedInfo[];
		CT: GeneratedInfo[];
		Both: GeneratedInfo[];
	}
	let selectedVid: GeneratedInfo | undefined = undefined;
	const generateClicked = () => {
		let selectedMapText = maps.find(({ id }) => id === selectedMap);
		let selectedSideText = sides.find(({ id }) => id === selectedSide);
		if (!selectedMapText) {
			console.log('missing selected map text. This should be impossible');
			return;
		}
		if (!selectedSideText) {
			console.log('missing selected side text. This should be impossible');
			return;
		}
		let mapVids: MapInfo = selectedMapText.vids;

		let selectedVids: GeneratedInfo[] = [];
		switch (selectedSideText.text) {
			case 'T':
				selectedVids = [...mapVids.T, ...mapVids.Both];
				break;
			case 'CT':
				selectedVids = [...mapVids.CT, ...mapVids.Both];
				break;
			case 'Both':
				selectedVids = [...mapVids.Both];
				break;
			default:
				console.log(
					'invalid selected side text. This should be impossible ' + selectedSideText.text
				);
		}
		selectedVid = selectedVids[Math.floor(Math.random() * selectedVids.length)];
		console.log('selected vid ' + JSON.stringify(selectedVid));
	};
</script>

<div class="m-3 space-y-4">
	<ComboBox bind:selectedId={selectedMap} titleText="Map" placeholder="Select map" items={maps} />

	<ComboBox
		bind:selectedId={selectedSide}
		titleText="Side"
		placeholder="Select side"
		items={sides}
	/>

	{#if selectedMap && selectedSide}
		<Button kind="danger" on:click={generateClicked}>Generate</Button>
	{:else}
		<Button kind="danger" disabled>Generate</Button>
	{/if}
	{#if selectedVid}
		<p>
			You must throw a
			<span class="text-[#ccba7c]">{selectedVid.startPosition}</span>
			→
			<span class="text-[#50d71e]">{selectedVid.endPosition}</span>
			<span class="text-[#de9b35]">{selectedVid.type}</span>
			this round:
		</p>
		<div class="w-[65vw] pb-[56.25%] relative">
			<iframe
                class="w-full h-full absolute"
				src="https://www.youtube.com/embed/{selectedVid.youTubeId}"
				title="roulette util video"
				frameborder="0"
				allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
				allowfullscreen
			></iframe>
		</div>
	{/if}
</div>
