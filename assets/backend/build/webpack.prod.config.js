const webpack = require('webpack');
const ExtractTextPlugin = require('extract-text-webpack-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');
const cleanWebpackPlugin = require('clean-webpack-plugin');
const merge = require('webpack-merge');
const webpackBaseConfig = require('./webpack.base.config.js');
const fs = require('fs');
const path = require('path');
const package = require('../../../package.json');
const backendPath = path.resolve(__dirname, '../../../public/');

fs.open('./env.js', 'w', function (err, fd) {
    const buf = 'export default "production";';
    fs.write(fd, buf, 0, buf.length, 0, function (err, written, buffer) {});
});

module.exports = merge(webpackBaseConfig, {
    output: {
        //publicPath: '/public/backend/',
        publicPath: 'https://imgs.cong5.net/assets/backend/',
        filename: 'js/[name].js',
        chunkFilename: 'js/[name].chunk.js'
    },
    plugins: [
        new cleanWebpackPlugin(['backend/*'], {
            root: backendPath,
            verbose: true, //开启在控制台输出信息
            dry: false //启用删除文件
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
        new CopyWebpackPlugin([
            {
                from: 'assets/backend/src/views/main-components/theme-switch/theme'
            },
        ]),
    ]
});
