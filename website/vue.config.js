const {
  resolve
} = require('path')
const path = require('path')
const WebpackBar = require('webpackbar')

module.exports = {
  publicPath: '/',
  assetsDir: 'static',
  outputDir: 'dist',
  lintOnSave: false,
  devServer: {
    hot: true,
    port: 9999,
    overlay: {
      warnings: true,
      errors: true
    }
  },
  configureWebpack() {
    return {
      resolve: {
        alias: {
          '@': resolve('src'),
          'Assets': resolve('src/assets')
        }
      },
      output: {
        devtoolModuleFilenameTemplate: info =>
          (info.resourcePath.match(/^src\//) ||
            (info.resourcePath.match(/^\.\/src\//) &&
              !info.resourcePath.match(/\.vue/))) ?
          'src:///' + info.resourcePath.replace(/^src\//, '')
          .replace(/^\.\/src\//, '') :
          `webpack:///${info.resourcePath}?${info.hash}`,
        devtoolFallbackModuleFilenameTemplate: 'webpack:///[resource-path]?[hash]'
      },
      module: {
        rules: [{
          test: /\.scss$/,
          use: ['sass-loader', {
            loader: 'style-resources-loader',
            options: {
              patterns: [
                path.resolve(__dirname, 'src/styles/_variables.scss'),
                path.resolve(__dirname, 'src/styles/_mixins.scss'),
              ]
            }
          }]
        },
          {
            include: /node_modules/,
            test: /\.mjs$/,
            type: 'javascript/auto'
          }]
      },
      plugins: [
        new WebpackBar({
          name: '明珠变压器设计平台',
        })
      ]
    }
  },
  chainWebpack: config => {
    const svgRule = config.module.rule('svg')
    svgRule.uses.clear()
    svgRule
      .use('svg-sprite-loader')
      .loader('svg-sprite-loader')
      .options({
        symbolId: 'icon-[name]'
      })
      .end()
  }
}
