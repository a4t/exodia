var gulp = require('gulp');
var shell = require('gulp-shell')

gulp.task('app_build', function(){
  gulp.src('./*.go')
    .pipe(shell([
      'go build -o ./exodia ./exodia.go',
      './exodia -f ./script.yml -d'
    ]))
})

gulp.task('watch', function(){
  gulp.watch('./*.go', ['app_build']);
});

gulp.task('default', ['app_build']);