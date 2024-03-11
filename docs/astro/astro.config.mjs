import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

// https://astro.build/config
export default defineConfig({
	integrations: [
		starlight({
			title: 'Fairshare',
			social: {
				github: 'https://github.com/manojnakp/fairshare',
			},
			sidebar: [
				{
					label: 'Walkthrough',
					autogenerate: { directory: 'walkthrough' },
				},
			],
		}),
	],
    vite: {
        resolve: {
            preserveSymlinks: true,
        },
    },
});
