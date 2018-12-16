const webpack = require('webpack');
const ExtractTextPlugin = require('extract-text-webpack-plugin');
const cleanWebpackPlugin = require('clean-webpack-plugin');
const merge = require('webpack-merge');
const webpackBaseConfig = require('./webpack.base.config.js');
const path = require('path');
const appPath = path.resolve(__dirname, '../../../public/');

module.exports = merge(webpackBaseConfig, {
    output: {
        //publicPath: '/public/app/',
        publicPath: 'https://o75u5ooep.qnssl.com/assets/app/',
        filename: 'js/[name].js',
        chunkFilename: 'js/[name].chunk.js'
    },
    plugins: [
        new cleanWebpackPlugin(['app/*'], {
            root: appPath,
            verbose: true,
            dry: false
        }),
        new ExtractTextPlugin({
            filename: 'css/[name].css',
            allChunks: true
        }),
        new webpack.DefinePlugin({
            'process.env': {
                NODE_ENV: '"production"'
            }
        }),
        new webpack.optimize.UglifyJsPlugin({
            compress: {
                warnings: false
            }
        }),
    ]
});
