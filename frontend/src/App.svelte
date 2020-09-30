<script>

	import { Categories, IndexRequest } from '../proto/categories_pb'
	import { CategoryServicePromiseClient } from '../proto/categories_grpc_web_pb' 
	import Cat from './services/category' 
	export let name;


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

	function doSomething(){
		myCall().then((r)=>{
			console.log('success')
			console.log(r)
		}).catch(error=>{
			console.log('error', error)
		})
	}
</script>

<main>
	<h1>Hello {name}!</h1>
	<p>Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn how to build Svelte apps.</p>
	<button on:click={doSomething}>Button</button>
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

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>