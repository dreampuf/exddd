var dest = './js',
  src = '.',
  mui = './node_modules/material-ui/src';

module.exports = {
  browserSync: {
    server: {
      // We're serving the src folder as well
      // for sass sourcemap linking
      baseDir: ["./"],
      directory: true,
      index: "view/test.html"
    },
    files: [
      dest + '/**'
    ]
  },
  less: {
    src: src + '/less/main.less',
    watch: [
      src + '/less/**',
      mui + '/less/**'
    ],
    dest: "./css"
  },
  markup: {
    src: src + "view/**.html",
    dest: "./view/"
  },
  browserify: {
    // Enable source maps
    debug: true,
    // A separate bundle will be generated for each
    // bundle config in the list below
    bundleConfigs: [{
      entries: './app/app.jsx',
      dest: "./js",
      outputName: 'app.js'
    }],
    watch: [
      "./app/*",
      "./app/components/*",
      "./app/components/pages/*",
    ]
  }
};
