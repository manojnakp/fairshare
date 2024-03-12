import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';
import starlightOpenAPI, { openAPISidebarGroups } from 'starlight-openapi';

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
                ...openAPISidebarGroups,
			],
            plugins: [
                starlightOpenAPI([
                    {
                        base: 'api',
                        label: 'OpenAPI Reference',
                        schema: './public/api/openapi.json',
                    }
                ])
            ]
		}),
	],
    vite: {
        resolve: {
            preserveSymlinks: true,
        },
    },
});
