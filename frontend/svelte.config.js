import preprocess from 'svelte-preprocess';
import node from '@sveltejs/adapter-node';
import path from 'path';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://github.com/sveltejs/svelte-preprocess
	// for more information about preprocessors
	preprocess: preprocess(),

	kit: {
		// hydrate the <div id="svelte"> element in src/app.html
		target: '#svelte',
		// build out as simple node app
		adapter: node(),

		vite: {
			resolve: {
				alias: {
					$utils: path.resolve('./src/utils')
				}
			}
		}
	}
};

export default config;
