/* craco.config.js */
const CracoLessPlugin = require('craco-less');

module.exports = {
    plugins: [
        {
            plugin: CracoLessPlugin,
            options: {
                lessLoaderOptions: {
                    lessOptions: {
                        modifyVars: {
                            '@primary-color': '#418B83',
                            '@layout-header-color': 'white',
                            '@layout-header-background': '#006d77',
                        },
                        javascriptEnabled: true,
                    },
                },
            },
        },
    ],
};
