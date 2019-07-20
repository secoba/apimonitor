<template>
  <div class="container" style="text-algin:center;" >
    <div>
      <canvas id="canvas" width="633" height="548"></canvas>
    </div>
    <form class="card login-container " style="position:absolute;width:372px;background-color:rgba(255, 255, 255, 0.17);top:40%;left:50%;margin: -150px 0 0 -200px;">
      <div class="title mt-5" style="text-align:center;color:rgba(245, 247, 250, 0.88);font-size:24px;margin-top: 70px">自动化测试平台</div>
      <el-input
        placeholder="请输入用户名"
        v-model="username"
        prop="username"
        clearable style="text-align:center;width:250px;hight:32px;margin-top:80px;margin-left:16%;">
      </el-input>
      <el-input
        type="password"
        placeholder="请输入密码"
        v-model="password"
        clearable style="text-align:center;width:250px;hight:32px;margin-top:30px;margin-left:16%;">
      </el-input>
      <br>
      <el-button plain v-loading="loading" @click="login($event)" style="margin-left:80px; size:10px;width:219px;margin-top:70px;" type="primary" >登录</el-button>

      <div style="font-size:12px;color:#ffffff;text-align:center;margin-top:30px;margin-bottom:50px">
        接口自动化、APP自动化、测试报告、任务设置
      </div>
    </form>


  </div>
</template>

<script>
import { requestLogin } from '../../api/api';
import md5 from 'js-md5';
export default {
  name: 'Login',
  data () {
    const validateUsername = (rule, value, callback) => {
        if (value.trim().length<1) {
          callback(new Error('用户名不能为空'))
        } else {
          callback()
        }
      };
      const validatePass = (rule, value, callback) => {
        if (value.trim().length < 1) {
          callback(new Error('密码不能为空'))
        } else {
          callback()
        }
      };
    return {
      msg: 'login',
      loading :false,
      code: 1,
      logining: false,
      username:"",
      password: "",
      rules: {
          name: [
            { required: true, message: '请输入活动名称', trigger: 'blur' },
            { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
          ],
        },
      rules2: {
          username: [
              { required: true, message: '请输入账号', trigger: 'blur' }
          ],
          password: [
              { required: true, message: '请输入密码', trigger: 'blur' }
          ]
      },
      checked: true
    }
  },
  mounted:function () {
      this.run();
  },
  methods: {
    vlog () {
      console.log("ok this logs")
    },
    greet: function (event) {
      alert('Hello ' + this.name + '!')
      if (event) {
        alert(event.target.tagName)
      }
    },
    say: function (message) {
      alert(message)
    },
    test: function (ms) {
      return this.ms
    },
    login(event) {
      let that = this
      if (this.username === '' || this.password === '') {
        this.$message({
          showClose: true,
          message: '账号或密码不能为空',
          type: 'error'
        })
        return false
      } else {
        var par = {"username":this.username,"password":md5(this.password)}
        requestLogin(par).then(
          _data => {
            let { msg, code, data } = _data;
            if (code == 200) {
              sessionStorage.setItem('nickname', JSON.stringify(data.user.nickname));
              if (data.user.userimg ==='') {
                sessionStorage.setItem('usrimg', JSON.stringify('static/images/icon.jpg'));
              } else {
                sessionStorage.setItem('usrimg', JSON.stringify(data.user.userimg));
              }
              
              that.$store.dispatch('setToken', data.token).then(() => {
                that.$router.push({path: '/'})
              })
            } else {
              that.$message({
                showClose: true,
                message: msg,
                type: 'error'
                })
            }
        })
      }
    },
    one:function () {
        var $ = {};

        $.Particle = function( opt ) {
          this.radius = 7;
          this.x = opt.x;
          this.y = opt.y;
          this.angle = opt.angle;
          this.speed = opt.speed;
          this.accel = opt.accel;
          this.decay = 0.01;
          this.life = 1;
        };

        $.Particle.prototype.step = function( i ) {
          this.speed += this.accel;
          this.x += Math.cos( this.angle ) * this.speed;
          this.y += Math.sin( this.angle ) * this.speed;
          this.angle += $.PI / 64;
          this.accel *= 1.01;
          this.life -= this.decay;

          if( this.life <= 0 ) {
            $.particles.splice( i, 1 );
          }
        };

      $.Particle.prototype.draw = function( i ) {
        $.ctx.fillStyle = $.ctx.strokeStyle = 'hsla(' + ( $.tick + ( this.life * 120 ) ) + ', 100%, 60%, ' + this.life + ')';
        $.ctx.beginPath();
        if( $.particles[ i - 1 ] ) {
          $.ctx.moveTo( this.x, this.y );
          $.ctx.lineTo( $.particles[ i - 1 ].x, $.particles[ i - 1 ].y );
        }
        $.ctx.stroke();

        $.ctx.beginPath();
        $.ctx.arc( this.x, this.y, Math.max( 0.001, this.life * this.radius ), 0, $.TWO_PI );
        $.ctx.fill();

        var size = Math.random() * 1.25;
        $.ctx.fillRect( ~~( this.x + ( ( Math.random() - 0.5 ) * 35 ) * this.life ), ~~( this.y + ( ( Math.random() - 0.5 ) * 35 ) * this.life ), size, size );
      }

      $.step = function() {
        $.particles.push( new $.Particle({
          x: $.width / 2 + Math.cos( $.tick / 20 ) * $.min / 2,
          y: $.height / 2 + Math.sin( $.tick / 20 ) * $.min / 2,
          angle: $.globalRotation + $.globalAngle,
          speed: 0,
          accel: 0.01
        }));

        $.particles.forEach( function( elem, index ) {
          elem.step( index );
        });

        $.globalRotation += $.PI / 6;
        $.globalAngle += $.PI / 6;
      };

      $.draw = function() {
        $.ctx.clearRect( 0, 0, $.width, $.height );

        $.particles.forEach( function( elem, index ) {
          elem.draw( index );
        });
      };

    $.init = function() {
      $.canvas = document.createElement( 'canvas' );
      $.ctx = $.canvas.getContext( '2d' );
      $.width = 300;
      $.height = 300;
      $.canvas.width = $.width * window.devicePixelRatio;
      $.canvas.height = $.height * window.devicePixelRatio;
      $.canvas.style.width = $.width + 'px';
      $.canvas.style.height = $.height + 'px';
      $.ctx.scale(window.devicePixelRatio, window.devicePixelRatio);
      $.min = $.width * 0.5;
      $.particles = [];
      $.globalAngle = 0;
      $.globalRotation = 0;
      $.tick = 0;
      $.PI = Math.PI;
      $.TWO_PI = $.PI * 2;
      $.ctx.globalCompositeOperation = 'lighter';
      document.body.appendChild( $.canvas );
      $.loop();
        };

      $.loop = function() {
        requestAnimationFrame( $.loop );
        $.step();
        $.draw();
        $.tick++;
      };
      $.init();
   },
   run: function () {
        const PHI = (1 + Math.sqrt(5)) / 2, // 1.618033988749895
        maxGeneration = (navigator.userAgent.toLowerCase().indexOf('firefox') > -1) ? 5 : 6,
        frameDuration = 1000 / 60,
        duration = 3000,
        rotationSpeed = 0.3,
        totalIterations = Math.floor(duration / frameDuration),
        maxBaseSize = 100,
        baseSizeSpeed = 0.02;

        var canvas = document.getElementById("canvas"),
        ctx = canvas.getContext("2d"),
        canvasWidth = document.documentElement.clientWidth,
        canvasHeight = document.documentElement.clientHeight,
        shapes = [],
        sizeVariation,
        iteration = 0,
        animationDirection = 1,
        sizeVariationRange = .15,
        baseRotation = 0,
        baseSize = 50,
        c1 = 43,
        c1S = 1,
        c2 = 205,
        c2S = 1,
        c3 = 255,
        c3S = 1;

        canvas.setAttribute("width", canvasWidth);
        canvas.setAttribute("height", canvasHeight);
        function Shape(gen, x, y, size, rotation) {
          this.generation = gen;
          this.size = size;
          this.rotation = -rotation;
          this.start = {
            x: x,
            y: y
          };
          this.end = {
            x_1: this.start.x + Math.cos(degToRad(this.rotation)) * this.size,
            y_1: this.start.y + Math.sin(degToRad(this.rotation)) * this.size,
            x_2: this.start.x + Math.cos(degToRad(this.rotation + 360 / 3)) * this.size,
            y_2: this.start.y + Math.sin(degToRad(this.rotation + 360 / 3)) * this.size,
            x_3:
              this.start.x +
              Math.cos(degToRad(this.rotation + 360 / 3 * 2)) * this.size,
            y_3:
              this.start.y + Math.sin(degToRad(this.rotation + 360 / 3 * 2)) * this.size
          };
          this.init();
        }
        Shape.prototype.init = function() {
            if (this.generation < maxGeneration) {
              var gen = this.generation + 1,
                newSize = this.size * sizeVariation,
                newRotation = this.rotation;

              shapes.push(
                new Shape(gen, this.end.x_1, this.end.y_1, newSize, newRotation)
              );
              shapes.push(
                new Shape(gen, this.end.x_2, this.end.y_2, newSize, newRotation)
              );
              shapes.push(
                new Shape(gen, this.end.x_3, this.end.y_3, newSize, newRotation)
              );
            }
            this.draw();
          };

          Shape.prototype.draw = function() {
            ctx.beginPath();
            ctx.moveTo(this.start.x, this.start.y);
      ctx.lineTo(this.end.x_1, this.end.y_1);
      ctx.moveTo(this.start.x, this.start.y);
      ctx.lineTo(this.end.x_2, this.end.y_2);
      ctx.moveTo(this.start.x, this.start.y);
      ctx.lineTo(this.end.x_3, this.end.y_3);
      //ctx.closePath();
      ctx.strokeStyle =
        "rgba(" + c1 + "," + c2 + "," + c3 + "," + 1 / this.generation / 5 + ")";
      ctx.stroke();
      //ctx.fill();
    };

    function animate() {
      //ctx.clearRect(0, 0, canvasWidth, canvasHeight);
      ctx.globalCompositeOperation = "source-over";
      ctx.fillStyle = "rgba(0,0,0,.1)";
      ctx.fillRect(0, 0, canvasWidth, canvasHeight);
      ctx.globalCompositeOperation = "lighter";
      shapes = [];
      shapes.push(
        new Shape(0, canvasWidth / 2, canvasHeight / 2, baseSize, baseRotation)
      );

      changeColor();
      iteration++;
      if (baseSize < maxBaseSize) baseSize += baseSizeSpeed;
      baseRotation += rotationSpeed;
      sizeVariation = easeInOutSine(
        iteration,
        1 - sizeVariationRange * animationDirection,
        sizeVariationRange * 2 * animationDirection,
        totalIterations
      );
      if (iteration >= totalIterations) {
        iteration = 0;
        animationDirection *= -1;
      }
      requestAnimationFrame(animate);
    }

    function degToRad(deg) {
      return Math.PI / 180 * deg;
    }

    function easeInOutSine(
      currentIteration,
      startValue,
      changeInValue,
      totalIterations
    ) {
      return (
        changeInValue /
          2 *
          (1 - Math.cos(Math.PI * currentIteration / totalIterations)) +
        startValue
      );
    }

    function changeColor() {
      if (c1 == 0 || c1 == 255) c1S *= -1;
      if (c2 == 0 || c2 == 255) c2S *= -1;
      if (c3 == 0 || c3 == 255) c3S *= -1;
      c1 += 1 * c1S;
      c2 += 1 * c2S;
      c3 += 1 * c3S;
    }

    ctx.globalCompositeOperation = "lighter";
    animate();

    window.addEventListener("resize", function() {
      canvasWidth = document.documentElement.clientWidth;
      canvasHeight = document.documentElement.clientHeight;

      canvas.setAttribute("width", canvasWidth);
      canvas.setAttribute("height", canvasHeight);
      ctx.strokeStyle = "rgba(66,134,240,.3)";
      ctx.globalCompositeOperation = "lighter";
    });
      },

  }

}
</script>

<style>

body {

  overflow: hidden;
}

canvas {
  bottom: 0;
  left: 0;
  margin: auto;
  position: absolute;
  right: 10%;
  top: 0;
}
</style>
