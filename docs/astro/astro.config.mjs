import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

// https://astro.build/config
export default defineConfig({
	integrations: [
		starlight({
			title: 'Fairshare',
            description: 'Split bills and expenses shared among a group fairly.',
			social: {
				github: 'https://github.com/manojnakp/fairshare',
			},
            editLink: {
                baseUrl: 'https://github.com/manojnakp/fairshare/edit/main/docs/astro',
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
