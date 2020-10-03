<script>

	import { Categories, IndexRequest } from '../proto/categories_pb'
	import { CategoryServicePromiseClient } from '../proto/categories_grpc_web_pb' 
	import Cat from './services/category' 
	export let name;
	let cats = [];
	let catsJson = [];


	async function myCall(){
		const getMetadata = () => {
			return {
				['X-CSRF-Token']: localStorage.getItem('X-CSRF-Token')
			}
		}
		
		var deps = {
			tokenKey: 'X-CSRF-Token',
			getMetadata,
			proto: {
				IndexRequest,
				Categories,
				CategoryClient: CategoryServicePromiseClient
			}
		}

		let cat = new Cat(deps)
		return await cat.index()
	}

	function getCategories(){
		myCall().then(response=>{
			cats = response;
		}).catch(error=>{
			console.log('error', error);
		})
	}

	function getCategoriesJson(){
		fetch('http://localhost:9003')
		.then(response => response.json())
		.then(data => catsJson=data)
		.catch(error=>{
			console.log('error', error)
		})
	}

	function getMethods(obj)
	{
		var res = [];
		for(var m in obj) {
			if(typeof obj[m] == "function") {
				res.push(m)
			}
		}
		return res;
	}

</script>

<main>
	<h1>gRPC-web!</h1>
	<p>Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn how to build Svelte apps.</p>
	<button on:click={getCategories}>Get random cities (RPC)</button>

	<ul>
		{#each cats.slice(0,50) as cat}
			<li>
				{cat.id} - {cat.name}
			</li>
		{/each}
	</ul>

	<button on:click={getCategoriesJson}>Get random cities (JSON)</button>
		<ul>
		{#each catsJson.slice(0,50) as cat}
			<li>
				{cat.id} - {cat.name}
			</li>
		{/each}
	</ul>


</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	ul {
		display: flex;
		flex-direction: column;
		flex-wrap: wrap;
		align-content: flex-start;
		list-style: none;
		margin: 0;
		padding: 0;
		height: 30em;
	}

	li {
		/* background: gray; */
		width: 10em;
		height: 1.1em;
		margin: .1em;
		text-align: left;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>