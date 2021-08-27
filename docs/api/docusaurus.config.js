/** @type {import('@docusaurus/types').DocusaurusConfig} */
const path = require("path");

module.exports = {
  title: 'Yotas',
  tagline: 'Oss\'s monetisation system',
  url: 'https://your-docusaurus-test-site.com',
  baseUrl: '/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.ico',
  organizationName: 'facebook', // Usually your GitHub org/user name.
  projectName: 'docusaurus', // Usually your repo name.
  themeConfig: {
    navbar: {
      title: 'Yotas',
      logo: {
        alt: 'My Site Logo',
        src: 'img/logo.svg',
      },
      items: [
        {
          type: 'doc',
          docId: 'intro',
          position: 'left',
          label: 'Docs',
        },
        {to: '/api', label: 'API', position: 'left'},
        {
          href: 'https://github.com/osscameroon/yotas',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Docs',
          items: [
            {
              label: 'API',
              to: '/api'
            },
            {
              label: 'Docs',
              to: '/docs/intro',
            },
          ],
        },
        {
          title: 'Community',
          items: [
            {
              label: 'Telegram',
              href: 'https://t.me/joinchat/UpKZh_T3W02LsGvQ',
            },
            {
              label: 'Youtube',
              href: 'https://www.youtube.com/channel/UCo89S8aPpJSXpLxkRjCbRaQ'
            },
            {
              label: 'Twitter',
              href: 'https://twitter.com/osscameroon',
            },
          ],
        },
        {
          title: 'More',
          items: [
            {
              label: 'Blog',
              to: 'https://blog.osscameroon.com',
            },
            {
              label: 'GitHub',
              href: 'https://github.com/osscameroon/yotas',
            },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} OSS Cameroon. Built with Docusaurus.`,
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          editUrl:
            'https://github.com/facebook/docusaurus/edit/master/website/',
        },
        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          editUrl:
            'https://github.com/facebook/docusaurus/edit/master/website/blog/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
  ],
  plugins: [
    path.resolve(__dirname, 'docusaurus-plugin')
  ]
};
