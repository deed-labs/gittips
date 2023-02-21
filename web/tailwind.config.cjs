/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	daisyui: {
		themes: [
			{
				mytheme: {
					primary: '#0088CC',
					secondary: '#6aabf6',
					accent: '#232328',
					neutral: '#232328',
					'base-100': '#FFFFFF',
					'base-200': '#F7F9FB',
					info: '#406EED',
					success: '#137638',
					warning: '#E8C721',
					error: '#EB6E60',
					github: '#282c2c'
				}
			}
		]
	},
	plugins: [require('daisyui')]
};
