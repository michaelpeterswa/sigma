<script lang="ts">
	export let name: string;
	import { onMount } from "svelte";
	import { App, applist } from './app_store.js';
	import { Count, count } from './count_store.js';

	onMount(async () => {
	fetch("http://localhost:8080/versions")
	.then(response => response.json())
	.then(data => {
		data.forEach(element => {
			$applist = [...$applist, new App(element)]
		});
		console.log($applist)
  	}).catch(error => {
   		console.log(error);
    	return [];
		});
	});

	async function getCount() {
			fetch("http://localhost:8080/count")
			.then(response => response.json())
			.then(data => {
				$count = new Count(data)
				console.log($count.count)
			}).catch(error => {
				console.log(error);
				return "0";
		});
	}
</script>

<main>
	<div class="p-10">
		<h1 class="text-center font-mono font-normal text-yellow-700 text-6xl">{name}</h1>
	</div>
	<div class="p-5">
		<h2 class="text-center font-sans font-light text-yellow-500 text-2xl">welcome to the sigma stack</h2>
	</div>
	<div class="p-5">
		<h4 class="text-center font-sans font-normal text-yellow-500 text-xl">apps:</h4>
		<ul>
			{#each $applist as app}
				<li class="text-center font-sans font-light text-yellow-500 text-xl">{app.name} - {app.version}</li>
			{/each}
		</ul>
	</div>
	<div class="p-5 flex justify-center">
		<button class="bg-yellow-500 hover:bg-yellow-700 text-white font-bold py-2 px-4 rounded-full" on:click={getCount}>
			{$count.count}
		</button>
	</div>
</main>

<style>
	/* This is a single-line comment */
	div {
		align-content: center;
	} 
</style>