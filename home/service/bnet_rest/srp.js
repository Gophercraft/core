!function(dataAndEvents, factory) {
    if ("object" == typeof exports && "object" == typeof module) {
      module.exports = factory();
    } else {
      if ("function" == typeof define && define.amd) {
        define([], factory);
      } else {
        if ("object" == typeof exports) {
          exports.srp6aRoutines = factory();
        } else {
          dataAndEvents.srp6aRoutines = factory();
        }
      }
    }
  }(this, function() {
    return function(cursor) {
      /**
       * @param {number} name
       * @return {?}
       */
      function result(name) {
        if (modules[name]) {
          return modules[name].exports;
        }
        var m = modules[name] = {
          i : name,
          l : false,
          exports : {}
        };
        return cursor[name].call(m.exports, m, m.exports, result), m.l = true, m.exports;
      }
      var modules = {};
      return result.m = cursor, result.c = modules, result.d = function(expectedNumberOfNonCommentArgs, name, putativeSpy) {
        if (!result.o(expectedNumberOfNonCommentArgs, name)) {
          Object.defineProperty(expectedNumberOfNonCommentArgs, name, {
            enumerable : true,
            /** @type {Function} */
            get : putativeSpy
          });
        }
      }, result.r = function(expectedNumberOfNonCommentArgs) {
        if ("undefined" != typeof Symbol) {
          if (Symbol.toStringTag) {
            Object.defineProperty(expectedNumberOfNonCommentArgs, Symbol.toStringTag, {
              value : "Module"
            });
          }
        }
        Object.defineProperty(expectedNumberOfNonCommentArgs, "__esModule", {
          value : true
        });
      }, result.t = function(str, args) {
        if (1 & args && (str = result(str)), 8 & args) {
          return str;
        }
        if (4 & args && ("object" == typeof str && (str && str.__esModule))) {
          return str;
        }
        /** @type {Object} */
        var expectedNumberOfNonCommentArgs = Object.create(null);
        if (result.r(expectedNumberOfNonCommentArgs), Object.defineProperty(expectedNumberOfNonCommentArgs, "default", {
          enumerable : true,
          value : str
        }), 2 & args && "string" != typeof str) {
          var path;
          for (path in str) {
            result.d(expectedNumberOfNonCommentArgs, path, function(key) {
              return str[key];
            }.bind(null, path));
          }
        }
        return expectedNumberOfNonCommentArgs;
      }, result.n = function(c) {
        /** @type {function (): ?} */
        var value = c && c.__esModule ? function() {
          return c.default;
        } : function() {
          return c;
        };
        return result.d(value, "a", value), value;
      }, result.o = function(array, keepData) {
        return Object.prototype.hasOwnProperty.call(array, keepData);
      }, result.p = "", result.p = window.serverResourceUrl || result.p, result(result.s = 61);
    }([function(module, dataAndEvents, deepDataAndEvents) {
      (function() {
        /**
         * @param {number} deepDataAndEvents
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {number} timestep
         * @return {undefined}
         */
        function self(deepDataAndEvents, expectedNumberOfNonCommentArgs, timestep) {
          if (null != deepDataAndEvents) {
            if ("number" == typeof deepDataAndEvents) {
              this.fromNumber(deepDataAndEvents, expectedNumberOfNonCommentArgs, timestep);
            } else {
              if (null == expectedNumberOfNonCommentArgs && "string" != typeof deepDataAndEvents) {
                this.fromString(deepDataAndEvents, 256);
              } else {
                this.fromString(deepDataAndEvents, expectedNumberOfNonCommentArgs);
              }
            }
          }
        }
        /**
         * @return {?}
         */
        function parseInt() {
          return new self(null);
        }
        /**
         * @param {?} val
         * @return {?}
         */
        function toString(val) {
          return "0123456789abcdefghijklmnopqrstuvwxyz".charAt(val);
        }
        /**
         * @param {(Array|Uint8Array|string)} body
         * @param {number} i
         * @return {?}
         */
        function callback(body, i) {
          var y = index[body.charCodeAt(i)];
          return null == y ? -1 : y;
        }
        /**
         * @param {number} recurring
         * @return {?}
         */
        function require(recurring) {
          var charCodeToReplace = parseInt();
          return charCodeToReplace.fromInt(recurring), charCodeToReplace;
        }
        /**
         * @param {number} type
         * @return {?}
         */
        function checkType(type) {
          var fx;
          /** @type {number} */
          var r = 1;
          return 0 != (fx = type >>> 16) && (type = fx, r += 16), 0 != (fx = type >> 8) && (type = fx, r += 8), 0 != (fx = type >> 4) && (type = fx, r += 4), 0 != (fx = type >> 2) && (type = fx, r += 2), 0 != (fx = type >> 1) && (type = fx, r += 1), r;
        }
        /**
         * @param {?} m
         * @return {undefined}
         */
        function D(m) {
          this.m = m;
        }
        /**
         * @param {Object} v
         * @return {undefined}
         */
        function Transform(v) {
          /** @type {Object} */
          this.m = v;
          this.mp = v.invDigit();
          /** @type {number} */
          this.mpl = 32767 & this.mp;
          /** @type {number} */
          this.mph = this.mp >> 15;
          /** @type {number} */
          this.um = (1 << v.DB - 15) - 1;
          /** @type {number} */
          this.mt2 = 2 * v.t;
        }
        /**
         * @param {number} dataAndEvents
         * @param {number} deepDataAndEvents
         * @return {?}
         */
        function clone(dataAndEvents, deepDataAndEvents) {
          return dataAndEvents & deepDataAndEvents;
        }
        /**
         * @param {number} dx
         * @param {number} dy
         * @return {?}
         */
        function w(dx, dy) {
          return dx | dy;
        }
        /**
         * @param {number} dataAndEvents
         * @param {number} deepDataAndEvents
         * @return {?}
         */
        function opt_obj2(dataAndEvents, deepDataAndEvents) {
          return dataAndEvents ^ deepDataAndEvents;
        }
        /**
         * @param {number} ast
         * @param {?} rev
         * @return {?}
         */
        function walk(ast, rev) {
          return ast & ~rev;
        }
        /**
         * @param {number} o
         * @return {?}
         */
        function compileNode(o) {
          if (0 == o) {
            return-1;
          }
          /** @type {number} */
          var e = 0;
          return 0 == (65535 & o) && (o >>= 16, e += 16), 0 == (255 & o) && (o >>= 8, e += 8), 0 == (15 & o) && (o >>= 4, e += 4), 0 == (3 & o) && (o >>= 2, e += 2), 0 == (1 & o) && ++e, e;
        }
        /**
         * @param {number} parent
         * @return {?}
         */
        function promote(parent) {
          /** @type {number} */
          var t = 0;
          for (;0 != parent;) {
            parent &= parent - 1;
            ++t;
          }
          return t;
        }
        /**
         * @return {undefined}
         */
        function Filter() {
        }
        /**
         * @param {?} border
         * @return {?}
         */
        function reduce(border) {
          return border;
        }
        /**
         * @param {Object} m
         * @return {undefined}
         */
        function error(m) {
          this.r2 = parseInt();
          this.q3 = parseInt();
          self.ONE.dlShiftTo(2 * m.t, this.r2);
          this.mu = this.r2.divide(m);
          /** @type {Object} */
          this.m = m;
        }
        /**
         * @return {undefined}
         */
        function onComplete() {
          var t;
          /** @type {number} */
          t = (new Date).getTime();
          data[x++] ^= 255 & t;
          data[x++] ^= t >> 8 & 255;
          data[x++] ^= t >> 16 & 255;
          data[x++] ^= t >> 24 & 255;
          if (x >= y) {
            x -= y;
          }
        }
        /**
         * @return {?}
         */
        function finish() {
          if (null == stream) {
            onComplete();
            (stream = new $).init(data);
            /** @type {number} */
            x = 0;
            for (;x < data.length;++x) {
              /** @type {number} */
              data[x] = 0;
            }
            /** @type {number} */
            x = 0;
          }
          return stream.next();
        }
        /**
         * @return {undefined}
         */
        function Type() {
        }
        /**
         * @return {undefined}
         */
        function $() {
          /** @type {number} */
          this.i = 0;
          /** @type {number} */
          this.j = 0;
          /** @type {Array} */
          this.S = new Array;
        }
        var DB;
        /** @type {boolean} */
        var Netscape = "undefined" != typeof navigator;
        if (Netscape && "Microsoft Internet Explorer" == navigator.appName) {
          /**
           * @param {number} recurring
           * @param {number} deepDataAndEvents
           * @param {Object} num
           * @param {number} noIn
           * @param {number} mayParseLabeledStatementInstead
           * @param {number} dataAndEvents
           * @return {?}
           */
          self.prototype.am = function(recurring, deepDataAndEvents, num, noIn, mayParseLabeledStatementInstead, dataAndEvents) {
            /** @type {number} */
            var a10 = 32767 & deepDataAndEvents;
            /** @type {number} */
            var x = deepDataAndEvents >> 15;
            for (;--dataAndEvents >= 0;) {
              /** @type {number} */
              var a00 = 32767 & this[recurring];
              /** @type {number} */
              var y = this[recurring++] >> 15;
              /** @type {number} */
              var f = x * a00 + y * a10;
              /** @type {number} */
              mayParseLabeledStatementInstead = ((a00 = a10 * a00 + ((32767 & f) << 15) + num[noIn] + (1073741823 & mayParseLabeledStatementInstead)) >>> 30) + (f >>> 15) + x * y + (mayParseLabeledStatementInstead >>> 30);
              /** @type {number} */
              num[noIn++] = 1073741823 & a00;
            }
            return mayParseLabeledStatementInstead;
          };
          /** @type {number} */
          DB = 30;
        } else {
          if (Netscape && "Netscape" != navigator.appName) {
            /**
             * @param {number} recurring
             * @param {number} deepDataAndEvents
             * @param {Object} num
             * @param {number} noIn
             * @param {number} mayParseLabeledStatementInstead
             * @param {number} dataAndEvents
             * @return {?}
             */
            self.prototype.am = function(recurring, deepDataAndEvents, num, noIn, mayParseLabeledStatementInstead, dataAndEvents) {
              for (;--dataAndEvents >= 0;) {
                var sectionLength = deepDataAndEvents * this[recurring++] + num[noIn] + mayParseLabeledStatementInstead;
                /** @type {number} */
                mayParseLabeledStatementInstead = Math.floor(sectionLength / 67108864);
                /** @type {number} */
                num[noIn++] = 67108863 & sectionLength;
              }
              return mayParseLabeledStatementInstead;
            };
            /** @type {number} */
            DB = 26;
          } else {
            /**
             * @param {number} recurring
             * @param {number} deepDataAndEvents
             * @param {Object} num
             * @param {number} noIn
             * @param {number} mayParseLabeledStatementInstead
             * @param {number} dataAndEvents
             * @return {?}
             */
            self.prototype.am = function(recurring, deepDataAndEvents, num, noIn, mayParseLabeledStatementInstead, dataAndEvents) {
              /** @type {number} */
              var a10 = 16383 & deepDataAndEvents;
              /** @type {number} */
              var x = deepDataAndEvents >> 14;
              for (;--dataAndEvents >= 0;) {
                /** @type {number} */
                var a00 = 16383 & this[recurring];
                /** @type {number} */
                var y = this[recurring++] >> 14;
                /** @type {number} */
                var f = x * a00 + y * a10;
                /** @type {number} */
                mayParseLabeledStatementInstead = ((a00 = a10 * a00 + ((16383 & f) << 14) + num[noIn] + mayParseLabeledStatementInstead) >> 28) + (f >> 14) + x * y;
                /** @type {number} */
                num[noIn++] = 268435455 & a00;
              }
              return mayParseLabeledStatementInstead;
            };
            /** @type {number} */
            DB = 28;
          }
        }
        self.prototype.DB = DB;
        /** @type {number} */
        self.prototype.DM = (1 << DB) - 1;
        /** @type {number} */
        self.prototype.DV = 1 << DB;
        /** @type {number} */
        self.prototype.FV = Math.pow(2, 52);
        /** @type {number} */
        self.prototype.F1 = 52 - DB;
        /** @type {number} */
        self.prototype.F2 = 2 * DB - 52;
        var o;
        var i;
        /** @type {Array} */
        var index = new Array;
        /** @type {number} */
        o = "0".charCodeAt(0);
        /** @type {number} */
        i = 0;
        for (;i <= 9;++i) {
          /** @type {number} */
          index[o++] = i;
        }
        /** @type {number} */
        o = "a".charCodeAt(0);
        /** @type {number} */
        i = 10;
        for (;i < 36;++i) {
          /** @type {number} */
          index[o++] = i;
        }
        /** @type {number} */
        o = "A".charCodeAt(0);
        /** @type {number} */
        i = 10;
        for (;i < 36;++i) {
          /** @type {number} */
          index[o++] = i;
        }
        /**
         * @param {Object} x
         * @return {?}
         */
        D.prototype.convert = function(x) {
          return x.s < 0 || x.compareTo(this.m) >= 0 ? x.mod(this.m) : x;
        };
        /**
         * @param {?} border
         * @return {?}
         */
        D.prototype.revert = function(border) {
          return border;
        };
        /**
         * @param {Object} border
         * @return {undefined}
         */
        D.prototype.reduce = function(border) {
          border.divRemTo(this.m, null, border);
        };
        /**
         * @param {?} el
         * @param {?} attribute
         * @param {Object} border
         * @return {undefined}
         */
        D.prototype.mulTo = function(el, attribute, border) {
          el.multiplyTo(attribute, border);
          this.reduce(border);
        };
        /**
         * @param {?} el
         * @param {Object} border
         * @return {undefined}
         */
        D.prototype.sqrTo = function(el, border) {
          el.squareTo(border);
          this.reduce(border);
        };
        /**
         * @param {?} value
         * @return {?}
         */
        Transform.prototype.convert = function(value) {
          var b = parseInt();
          return value.abs().dlShiftTo(this.m.t, b), b.divRemTo(this.m, null, b), value.s < 0 && (b.compareTo(self.ZERO) > 0 && this.m.subTo(b, b)), b;
        };
        /**
         * @param {?} el
         * @return {?}
         */
        Transform.prototype.revert = function(el) {
          var border = parseInt();
          return el.copyTo(border), this.reduce(border), border;
        };
        /**
         * @param {Object} border
         * @return {undefined}
         */
        Transform.prototype.reduce = function(border) {
          for (;border.t <= this.mt2;) {
            /** @type {number} */
            border[border.t++] = 0;
          }
          /** @type {number} */
          var noIn = 0;
          for (;noIn < this.m.t;++noIn) {
            /** @type {number} */
            var countrySym = 32767 & border[noIn];
            /** @type {number} */
            var deepDataAndEvents = countrySym * this.mpl + ((countrySym * this.mph + (border[noIn] >> 15) * this.mpl & this.um) << 15) & border.DM;
            border[countrySym = noIn + this.m.t] += this.m.am(0, deepDataAndEvents, border, noIn, 0, this.m.t);
            for (;border[countrySym] >= border.DV;) {
              border[countrySym] -= border.DV;
              border[++countrySym]++;
            }
          }
          border.clamp();
          border.drShiftTo(this.m.t, border);
          if (border.compareTo(this.m) >= 0) {
            border.subTo(this.m, border);
          }
        };
        /**
         * @param {?} el
         * @param {?} attribute
         * @param {Object} border
         * @return {undefined}
         */
        Transform.prototype.mulTo = function(el, attribute, border) {
          el.multiplyTo(attribute, border);
          this.reduce(border);
        };
        /**
         * @param {?} el
         * @param {Object} border
         * @return {undefined}
         */
        Transform.prototype.sqrTo = function(el, border) {
          el.squareTo(border);
          this.reduce(border);
        };
        /**
         * @param {?} value
         * @return {undefined}
         */
        self.prototype.copyTo = function(value) {
          /** @type {number} */
          var kkey = this.t - 1;
          for (;kkey >= 0;--kkey) {
            value[kkey] = this[kkey];
          }
          value.t = this.t;
          value.s = this.s;
        };
        /**
         * @param {number} recurring
         * @return {undefined}
         */
        self.prototype.fromInt = function(recurring) {
          /** @type {number} */
          this.t = 1;
          /** @type {number} */
          this.s = recurring < 0 ? -1 : 0;
          if (recurring > 0) {
            /** @type {number} */
            this[0] = recurring;
          } else {
            if (recurring < -1) {
              this[0] = recurring + this.DV;
            } else {
              /** @type {number} */
              this.t = 0;
            }
          }
        };
        /**
         * @param {Array} str
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        self.prototype.fromString = function(str, expectedNumberOfNonCommentArgs) {
          var VLQ_BASE_SHIFT;
          if (16 == expectedNumberOfNonCommentArgs) {
            /** @type {number} */
            VLQ_BASE_SHIFT = 4;
          } else {
            if (8 == expectedNumberOfNonCommentArgs) {
              /** @type {number} */
              VLQ_BASE_SHIFT = 3;
            } else {
              if (256 == expectedNumberOfNonCommentArgs) {
                /** @type {number} */
                VLQ_BASE_SHIFT = 8;
              } else {
                if (2 == expectedNumberOfNonCommentArgs) {
                  /** @type {number} */
                  VLQ_BASE_SHIFT = 1;
                } else {
                  if (32 == expectedNumberOfNonCommentArgs) {
                    /** @type {number} */
                    VLQ_BASE_SHIFT = 5;
                  } else {
                    if (4 != expectedNumberOfNonCommentArgs) {
                      return void this.fromRadix(str, expectedNumberOfNonCommentArgs);
                    }
                    /** @type {number} */
                    VLQ_BASE_SHIFT = 2;
                  }
                }
              }
            }
          }
          /** @type {number} */
          this.t = 0;
          /** @type {number} */
          this.s = 0;
          var last = str.length;
          /** @type {boolean} */
          var o = false;
          /** @type {number} */
          var shift = 0;
          for (;--last >= 0;) {
            var val = 8 == VLQ_BASE_SHIFT ? 255 & str[last] : callback(str, last);
            if (val < 0) {
              if ("-" == str.charAt(last)) {
                /** @type {boolean} */
                o = true;
              }
            } else {
              /** @type {boolean} */
              o = false;
              if (0 == shift) {
                this[this.t++] = val;
              } else {
                if (shift + VLQ_BASE_SHIFT > this.DB) {
                  this[this.t - 1] |= (val & (1 << this.DB - shift) - 1) << shift;
                  /** @type {number} */
                  this[this.t++] = val >> this.DB - shift;
                } else {
                  this[this.t - 1] |= val << shift;
                }
              }
              if ((shift += VLQ_BASE_SHIFT) >= this.DB) {
                shift -= this.DB;
              }
            }
          }
          if (8 == VLQ_BASE_SHIFT) {
            if (0 != (128 & str[0])) {
              /** @type {number} */
              this.s = -1;
              if (shift > 0) {
                this[this.t - 1] |= (1 << this.DB - shift) - 1 << shift;
              }
            }
          }
          this.clamp();
          if (o) {
            self.ZERO.subTo(this, this);
          }
        };
        /**
         * @return {undefined}
         */
        self.prototype.clamp = function() {
          /** @type {number} */
          var t = this.s & this.DM;
          for (;this.t > 0 && this[this.t - 1] == t;) {
            --this.t;
          }
        };
        /**
         * @param {number} dataAndEvents
         * @param {?} o
         * @return {undefined}
         */
        self.prototype.dlShiftTo = function(dataAndEvents, o) {
          var i;
          /** @type {number} */
          i = this.t - 1;
          for (;i >= 0;--i) {
            o[i + dataAndEvents] = this[i];
          }
          /** @type {number} */
          i = dataAndEvents - 1;
          for (;i >= 0;--i) {
            /** @type {number} */
            o[i] = 0;
          }
          o.t = this.t + dataAndEvents;
          o.s = this.s;
        };
        /**
         * @param {number} dataAndEvents
         * @param {Object} o
         * @return {undefined}
         */
        self.prototype.drShiftTo = function(dataAndEvents, o) {
          /** @type {number} */
          var i = dataAndEvents;
          for (;i < this.t;++i) {
            o[i - dataAndEvents] = this[i];
          }
          /** @type {number} */
          o.t = Math.max(this.t - dataAndEvents, 0);
          o.s = this.s;
        };
        /**
         * @param {number} pos
         * @param {Object} b
         * @return {undefined}
         */
        self.prototype.lShiftTo = function(pos, b) {
          var t;
          /** @type {number} */
          var clientTop = pos % this.DB;
          /** @type {number} */
          var top = this.DB - clientTop;
          /** @type {number} */
          var kind = (1 << top) - 1;
          /** @type {number} */
          var d = Math.floor(pos / this.DB);
          /** @type {number} */
          var num = this.s << clientTop & this.DM;
          /** @type {number} */
          t = this.t - 1;
          for (;t >= 0;--t) {
            /** @type {number} */
            b[t + d + 1] = this[t] >> top | num;
            /** @type {number} */
            num = (this[t] & kind) << clientTop;
          }
          /** @type {number} */
          t = d - 1;
          for (;t >= 0;--t) {
            /** @type {number} */
            b[t] = 0;
          }
          /** @type {number} */
          b[d] = num;
          b.t = this.t + d + 1;
          b.s = this.s;
          b.clamp();
        };
        /**
         * @param {number} dataAndEvents
         * @param {?} b
         * @return {undefined}
         */
        self.prototype.rShiftTo = function(dataAndEvents, b) {
          b.s = this.s;
          /** @type {number} */
          var minX = Math.floor(dataAndEvents / this.DB);
          if (minX >= this.t) {
            /** @type {number} */
            b.t = 0;
          } else {
            /** @type {number} */
            var clientTop = dataAndEvents % this.DB;
            /** @type {number} */
            var top = this.DB - clientTop;
            /** @type {number} */
            var s = (1 << clientTop) - 1;
            /** @type {number} */
            b[0] = this[minX] >> clientTop;
            /** @type {number} */
            var maxX = minX + 1;
            for (;maxX < this.t;++maxX) {
              b[maxX - minX - 1] |= (this[maxX] & s) << top;
              /** @type {number} */
              b[maxX - minX] = this[maxX] >> clientTop;
            }
            if (clientTop > 0) {
              b[this.t - minX - 1] |= (this.s & s) << top;
            }
            /** @type {number} */
            b.t = this.t - minX;
            b.clamp();
          }
        };
        /**
         * @param {?} b
         * @param {Object} value
         * @return {undefined}
         */
        self.prototype.subTo = function(b, value) {
          /** @type {number} */
          var i = 0;
          /** @type {number} */
          var s = 0;
          /** @type {number} */
          var padLength = Math.min(b.t, this.t);
          for (;i < padLength;) {
            s += this[i] - b[i];
            /** @type {number} */
            value[i++] = s & this.DM;
            s >>= this.DB;
          }
          if (b.t < this.t) {
            s -= b.s;
            for (;i < this.t;) {
              s += this[i];
              /** @type {number} */
              value[i++] = s & this.DM;
              s >>= this.DB;
            }
            s += this.s;
          } else {
            s += this.s;
            for (;i < b.t;) {
              s -= b[i];
              /** @type {number} */
              value[i++] = s & this.DM;
              s >>= this.DB;
            }
            s -= b.s;
          }
          /** @type {number} */
          value.s = s < 0 ? -1 : 0;
          if (s < -1) {
            value[i++] = this.DV + s;
          } else {
            if (s > 0) {
              value[i++] = s;
            }
          }
          /** @type {number} */
          value.t = i;
          value.clamp();
        };
        /**
         * @param {?} node
         * @param {Object} border
         * @return {undefined}
         */
        self.prototype.multiplyTo = function(node, border) {
          var options = this.abs();
          var row = node.abs();
          var noIn = options.t;
          border.t = noIn + row.t;
          for (;--noIn >= 0;) {
            /** @type {number} */
            border[noIn] = 0;
          }
          /** @type {number} */
          noIn = 0;
          for (;noIn < row.t;++noIn) {
            border[noIn + options.t] = options.am(0, row[noIn], border, noIn, 0, options.t);
          }
          /** @type {number} */
          border.s = 0;
          border.clamp();
          if (this.s != node.s) {
            self.ZERO.subTo(border, border);
          }
        };
        /**
         * @param {Object} num
         * @return {undefined}
         */
        self.prototype.squareTo = function(num) {
          var options = this.abs();
          /** @type {number} */
          var recurring = num.t = 2 * options.t;
          for (;--recurring >= 0;) {
            /** @type {number} */
            num[recurring] = 0;
          }
          /** @type {number} */
          recurring = 0;
          for (;recurring < options.t - 1;++recurring) {
            var mayParseLabeledStatementInstead = options.am(recurring, options[recurring], num, 2 * recurring, 0, 1);
            if ((num[recurring + options.t] += options.am(recurring + 1, 2 * options[recurring], num, 2 * recurring + 1, mayParseLabeledStatementInstead, options.t - recurring - 1)) >= options.DV) {
              num[recurring + options.t] -= options.DV;
              /** @type {number} */
              num[recurring + options.t + 1] = 1;
            }
          }
          if (num.t > 0) {
            num[num.t - 1] += options.am(recurring, options[recurring], num, 2 * recurring, 0, 1);
          }
          /** @type {number} */
          num.s = 0;
          num.clamp();
        };
        /**
         * @param {?} value
         * @param {Object} crossScope
         * @param {Object} d
         * @return {?}
         */
        self.prototype.divRemTo = function(value, crossScope, d) {
          var b = value.abs();
          if (!(b.t <= 0)) {
            var a = this.abs();
            if (a.t < b.t) {
              return null != crossScope && crossScope.fromInt(0), void(null != d && this.copyTo(d));
            }
            if (null == d) {
              d = parseInt();
            }
            var m = parseInt();
            var tmp_str = this.s;
            var s = value.s;
            /** @type {number} */
            var node = this.DB - checkType(b[b.t - 1]);
            if (node > 0) {
              b.lShiftTo(node, m);
              a.lShiftTo(node, d);
            } else {
              b.copyTo(m);
              a.copyTo(d);
            }
            var dataAndEvents = m.t;
            var formula = m[dataAndEvents - 1];
            if (0 != formula) {
              /** @type {number} */
              var amount = formula * (1 << this.F1) + (dataAndEvents > 1 ? m[dataAndEvents - 2] >> this.F2 : 0);
              /** @type {number} */
              var pct = this.FV / amount;
              /** @type {number} */
              var t = (1 << this.F1) / amount;
              /** @type {number} */
              var cy = 1 << this.F2;
              var j = d.t;
              /** @type {number} */
              var noIn = j - dataAndEvents;
              var v = null == crossScope ? parseInt() : crossScope;
              m.dlShiftTo(noIn, v);
              if (d.compareTo(v) >= 0) {
                /** @type {number} */
                d[d.t++] = 1;
                d.subTo(v, d);
              }
              self.ONE.dlShiftTo(dataAndEvents, v);
              v.subTo(m, m);
              for (;m.t < dataAndEvents;) {
                /** @type {number} */
                m[m.t++] = 0;
              }
              for (;--noIn >= 0;) {
                var deepDataAndEvents = d[--j] == formula ? this.DM : Math.floor(d[j] * pct + (d[j - 1] + cy) * t);
                if ((d[j] += m.am(0, deepDataAndEvents, d, noIn, 0, dataAndEvents)) < deepDataAndEvents) {
                  m.dlShiftTo(noIn, v);
                  d.subTo(v, d);
                  for (;d[j] < --deepDataAndEvents;) {
                    d.subTo(v, d);
                  }
                }
              }
              if (null != crossScope) {
                d.drShiftTo(dataAndEvents, crossScope);
                if (tmp_str != s) {
                  self.ZERO.subTo(crossScope, crossScope);
                }
              }
              d.t = dataAndEvents;
              d.clamp();
              if (node > 0) {
                d.rShiftTo(node, d);
              }
              if (tmp_str < 0) {
                self.ZERO.subTo(d, d);
              }
            }
          }
        };
        /**
         * @return {?}
         */
        self.prototype.invDigit = function() {
          if (this.t < 1) {
            return 0;
          }
          var a1 = this[0];
          if (0 == (1 & a1)) {
            return 0;
          }
          /** @type {number} */
          var b4 = 3 & a1;
          return(b4 = (b4 = (b4 = (b4 = b4 * (2 - (15 & a1) * b4) & 15) * (2 - (255 & a1) * b4) & 255) * (2 - ((65535 & a1) * b4 & 65535)) & 65535) * (2 - a1 * b4 % this.DV) % this.DV) > 0 ? this.DV - b4 : -b4;
        };
        /**
         * @return {?}
         */
        self.prototype.isEven = function() {
          return 0 == (this.t > 0 ? 1 & this[0] : this.s);
        };
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Object} v
         * @return {?}
         */
        self.prototype.exp = function(expectedNumberOfNonCommentArgs, v) {
          if (expectedNumberOfNonCommentArgs > 4294967295 || expectedNumberOfNonCommentArgs < 1) {
            return self.ONE;
          }
          var border = parseInt();
          var failuresLink = parseInt();
          var prop = v.convert(this);
          /** @type {number} */
          var a = checkType(expectedNumberOfNonCommentArgs) - 1;
          prop.copyTo(border);
          for (;--a >= 0;) {
            if (v.sqrTo(border, failuresLink), (expectedNumberOfNonCommentArgs & 1 << a) > 0) {
              v.mulTo(failuresLink, prop, border);
            } else {
              var name = border;
              border = failuresLink;
              failuresLink = name;
            }
          }
          return v.revert(border);
        };
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        self.prototype.toString = function(expectedNumberOfNonCommentArgs) {
          if (this.s < 0) {
            return "-" + this.negate().toString(expectedNumberOfNonCommentArgs);
          }
          var right;
          if (16 == expectedNumberOfNonCommentArgs) {
            /** @type {number} */
            right = 4;
          } else {
            if (8 == expectedNumberOfNonCommentArgs) {
              /** @type {number} */
              right = 3;
            } else {
              if (2 == expectedNumberOfNonCommentArgs) {
                /** @type {number} */
                right = 1;
              } else {
                if (32 == expectedNumberOfNonCommentArgs) {
                  /** @type {number} */
                  right = 5;
                } else {
                  if (4 != expectedNumberOfNonCommentArgs) {
                    return this.toRadix(expectedNumberOfNonCommentArgs);
                  }
                  /** @type {number} */
                  right = 2;
                }
              }
            }
          }
          var value;
          /** @type {number} */
          var UM = (1 << right) - 1;
          /** @type {boolean} */
          var escape = false;
          /** @type {string} */
          var str = "";
          var t = this.t;
          /** @type {number} */
          var left = this.DB - t * this.DB % right;
          if (t-- > 0) {
            if (left < this.DB) {
              if ((value = this[t] >> left) > 0) {
                /** @type {boolean} */
                escape = true;
                str = toString(value);
              }
            }
            for (;t >= 0;) {
              if (left < right) {
                /** @type {number} */
                value = (this[t] & (1 << left) - 1) << right - left;
                value |= this[--t] >> (left += this.DB - right);
              } else {
                /** @type {number} */
                value = this[t] >> (left -= right) & UM;
                if (left <= 0) {
                  left += this.DB;
                  --t;
                }
              }
              if (value > 0) {
                /** @type {boolean} */
                escape = true;
              }
              if (escape) {
                str += toString(value);
              }
            }
          }
          return escape ? str : "0";
        };
        /**
         * @return {?}
         */
        self.prototype.negate = function() {
          var pdataOld = parseInt();
          return self.ZERO.subTo(this, pdataOld), pdataOld;
        };
        /**
         * @return {?}
         */
        self.prototype.abs = function() {
          return this.s < 0 ? this.negate() : this;
        };
        /**
         * @param {?} o
         * @return {?}
         */
        self.prototype.compareTo = function(o) {
          /** @type {number} */
          var s2 = this.s - o.s;
          if (0 != s2) {
            return s2;
          }
          var t = this.t;
          if (0 != (s2 = t - o.t)) {
            return this.s < 0 ? -s2 : s2;
          }
          for (;--t >= 0;) {
            if (0 != (s2 = this[t] - o[t])) {
              return s2;
            }
          }
          return 0;
        };
        /**
         * @return {?}
         */
        self.prototype.bitLength = function() {
          return this.t <= 0 ? 0 : this.DB * (this.t - 1) + checkType(this[this.t - 1] ^ this.s & this.DM);
        };
        /**
         * @param {?} base
         * @return {?}
         */
        self.prototype.mod = function(base) {
          var b = parseInt();
          return this.abs().divRemTo(base, null, b), this.s < 0 && (b.compareTo(self.ZERO) > 0 && base.subTo(b, b)), b;
        };
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {?} value
         * @return {?}
         */
        self.prototype.modPowInt = function(expectedNumberOfNonCommentArgs, value) {
          var attempted;
          return attempted = expectedNumberOfNonCommentArgs < 256 || value.isEven() ? new D(value) : new Transform(value), this.exp(expectedNumberOfNonCommentArgs, attempted);
        };
        self.ZERO = require(0);
        self.ONE = require(1);
        /** @type {function (?): ?} */
        Filter.prototype.convert = reduce;
        /** @type {function (?): ?} */
        Filter.prototype.revert = reduce;
        /**
         * @param {?} el
         * @param {?} attribute
         * @param {Object} border
         * @return {undefined}
         */
        Filter.prototype.mulTo = function(el, attribute, border) {
          el.multiplyTo(attribute, border);
        };
        /**
         * @param {?} child
         * @param {?} el
         * @return {undefined}
         */
        Filter.prototype.sqrTo = function(child, el) {
          child.squareTo(el);
        };
        /**
         * @param {Object} value
         * @return {?}
         */
        error.prototype.convert = function(value) {
          if (value.s < 0 || value.t > 2 * this.m.t) {
            return value.mod(this.m);
          }
          if (value.compareTo(this.m) < 0) {
            return value;
          }
          var border = parseInt();
          return value.copyTo(border), this.reduce(border), border;
        };
        /**
         * @param {?} border
         * @return {?}
         */
        error.prototype.revert = function(border) {
          return border;
        };
        /**
         * @param {Object} border
         * @return {undefined}
         */
        error.prototype.reduce = function(border) {
          border.drShiftTo(this.m.t - 1, this.r2);
          if (border.t > this.m.t + 1) {
            border.t = this.m.t + 1;
            border.clamp();
          }
          this.mu.multiplyUpperTo(this.r2, this.m.t + 1, this.q3);
          this.m.multiplyLowerTo(this.q3, this.m.t + 1, this.r2);
          for (;border.compareTo(this.r2) < 0;) {
            border.dAddOffset(1, this.m.t + 1);
          }
          border.subTo(this.r2, border);
          for (;border.compareTo(this.m) >= 0;) {
            border.subTo(this.m, border);
          }
        };
        /**
         * @param {?} el
         * @param {?} attribute
         * @param {Object} border
         * @return {undefined}
         */
        error.prototype.mulTo = function(el, attribute, border) {
          el.multiplyTo(attribute, border);
          this.reduce(border);
        };
        /**
         * @param {?} el
         * @param {Object} border
         * @return {undefined}
         */
        error.prototype.sqrTo = function(el, border) {
          el.squareTo(border);
          this.reduce(border);
        };
        var stream;
        var data;
        var x;
        /** @type {Array} */
        var nodes = [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541, 547, 557, 563, 
        569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911, 919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997];
        /** @type {number} */
        var total_size = (1 << 26) / nodes[nodes.length - 1];
        if (self.prototype.chunkSize = function(expectedNumberOfNonCommentArgs) {
          return Math.floor(Math.LN2 * this.DB / Math.log(expectedNumberOfNonCommentArgs));
        }, self.prototype.toRadix = function(expectedNumberOfNonCommentArgs) {
          if (null == expectedNumberOfNonCommentArgs && (expectedNumberOfNonCommentArgs = 10), 0 == this.signum() || (expectedNumberOfNonCommentArgs < 2 || expectedNumberOfNonCommentArgs > 36)) {
            return "0";
          }
          var exponent = this.chunkSize(expectedNumberOfNonCommentArgs);
          /** @type {number} */
          var recurring = Math.pow(expectedNumberOfNonCommentArgs, exponent);
          var ctor = require(recurring);
          var b = parseInt();
          var end = parseInt();
          /** @type {string} */
          var buttons = "";
          this.divRemTo(ctor, b, end);
          for (;b.signum() > 0;) {
            buttons = (recurring + end.intValue()).toString(expectedNumberOfNonCommentArgs).substr(1) + buttons;
            b.divRemTo(ctor, b, end);
          }
          return end.intValue().toString(expectedNumberOfNonCommentArgs) + buttons;
        }, self.prototype.fromRadix = function(body, expectedNumberOfNonCommentArgs) {
          this.fromInt(0);
          if (null == expectedNumberOfNonCommentArgs) {
            /** @type {number} */
            expectedNumberOfNonCommentArgs = 10;
          }
          var exponent = this.chunkSize(expectedNumberOfNonCommentArgs);
          /** @type {number} */
          var dataAndEvents = Math.pow(expectedNumberOfNonCommentArgs, exponent);
          /** @type {boolean} */
          var o = false;
          /** @type {number} */
          var bits = 0;
          /** @type {number} */
          var dns = 0;
          /** @type {number} */
          var j = 0;
          for (;j < body.length;++j) {
            var host = callback(body, j);
            if (host < 0) {
              if ("-" == body.charAt(j)) {
                if (0 == this.signum()) {
                  /** @type {boolean} */
                  o = true;
                }
              }
            } else {
              dns = expectedNumberOfNonCommentArgs * dns + host;
              if (++bits >= exponent) {
                this.dMultiply(dataAndEvents);
                this.dAddOffset(dns, 0);
                /** @type {number} */
                bits = 0;
                /** @type {number} */
                dns = 0;
              }
            }
          }
          if (bits > 0) {
            this.dMultiply(Math.pow(expectedNumberOfNonCommentArgs, bits));
            this.dAddOffset(dns, 0);
          }
          if (o) {
            self.ZERO.subTo(this, this);
          }
        }, self.prototype.fromNumber = function(deepDataAndEvents, expectedNumberOfNonCommentArgs, n) {
          if ("number" == typeof expectedNumberOfNonCommentArgs) {
            if (deepDataAndEvents < 2) {
              this.fromInt(1);
            } else {
              this.fromNumber(deepDataAndEvents, n);
              if (!this.testBit(deepDataAndEvents - 1)) {
                this.bitwiseTo(self.ONE.shiftLeft(deepDataAndEvents - 1), w, this);
              }
              if (this.isEven()) {
                this.dAddOffset(1, 0);
              }
              for (;!this.isProbablePrime(expectedNumberOfNonCommentArgs);) {
                this.dAddOffset(2, 0);
                if (this.bitLength() > deepDataAndEvents) {
                  this.subTo(self.ONE.shiftLeft(deepDataAndEvents - 1), this);
                }
              }
            }
          } else {
            /** @type {Array} */
            var simple = new Array;
            /** @type {number} */
            var o = 7 & deepDataAndEvents;
            /** @type {number} */
            simple.length = 1 + (deepDataAndEvents >> 3);
            expectedNumberOfNonCommentArgs.nextBytes(simple);
            if (o > 0) {
              simple[0] &= (1 << o) - 1;
            } else {
              /** @type {number} */
              simple[0] = 0;
            }
            this.fromString(simple, 256);
          }
        }, self.prototype.bitwiseTo = function(t, f, r) {
          var i;
          var x;
          /** @type {number} */
          var lastLine = Math.min(t.t, this.t);
          /** @type {number} */
          i = 0;
          for (;i < lastLine;++i) {
            r[i] = f(this[i], t[i]);
          }
          if (t.t < this.t) {
            /** @type {number} */
            x = t.s & this.DM;
            /** @type {number} */
            i = lastLine;
            for (;i < this.t;++i) {
              r[i] = f(this[i], x);
            }
            r.t = this.t;
          } else {
            /** @type {number} */
            x = this.s & this.DM;
            /** @type {number} */
            i = lastLine;
            for (;i < t.t;++i) {
              r[i] = f(x, t[i]);
            }
            r.t = t.t;
          }
          r.s = f(this.s, t.s);
          r.clamp();
        }, self.prototype.changeBit = function(pixels, f) {
          var n = self.ONE.shiftLeft(pixels);
          return this.bitwiseTo(n, f, n), n;
        }, self.prototype.addTo = function(dataAndEvents, value) {
          /** @type {number} */
          var top = 0;
          /** @type {number} */
          var y = 0;
          /** @type {number} */
          var wTop = Math.min(dataAndEvents.t, this.t);
          for (;top < wTop;) {
            y += this[top] + dataAndEvents[top];
            /** @type {number} */
            value[top++] = y & this.DM;
            y >>= this.DB;
          }
          if (dataAndEvents.t < this.t) {
            y += dataAndEvents.s;
            for (;top < this.t;) {
              y += this[top];
              /** @type {number} */
              value[top++] = y & this.DM;
              y >>= this.DB;
            }
            y += this.s;
          } else {
            y += this.s;
            for (;top < dataAndEvents.t;) {
              y += dataAndEvents[top];
              /** @type {number} */
              value[top++] = y & this.DM;
              y >>= this.DB;
            }
            y += dataAndEvents.s;
          }
          /** @type {number} */
          value.s = y < 0 ? -1 : 0;
          if (y > 0) {
            value[top++] = y;
          } else {
            if (y < -1) {
              value[top++] = this.DV + y;
            }
          }
          /** @type {number} */
          value.t = top;
          value.clamp();
        }, self.prototype.dMultiply = function(dataAndEvents) {
          this[this.t] = this.am(0, dataAndEvents - 1, this, 0, 0, this.t);
          ++this.t;
          this.clamp();
        }, self.prototype.dAddOffset = function(expectedNumberOfNonCommentArgs, mayParseLabeledStatementInstead) {
          if (0 != expectedNumberOfNonCommentArgs) {
            for (;this.t <= mayParseLabeledStatementInstead;) {
              /** @type {number} */
              this[this.t++] = 0;
            }
            this[mayParseLabeledStatementInstead] += expectedNumberOfNonCommentArgs;
            for (;this[mayParseLabeledStatementInstead] >= this.DV;) {
              this[mayParseLabeledStatementInstead] -= this.DV;
              if (++mayParseLabeledStatementInstead >= this.t) {
                /** @type {number} */
                this[this.t++] = 0;
              }
              ++this[mayParseLabeledStatementInstead];
            }
          }
        }, self.prototype.multiplyLowerTo = function(o, value, cur) {
          var ret;
          /** @type {number} */
          var noIn = Math.min(this.t + o.t, value);
          /** @type {number} */
          cur.s = 0;
          /** @type {number} */
          cur.t = noIn;
          for (;noIn > 0;) {
            /** @type {number} */
            cur[--noIn] = 0;
          }
          /** @type {number} */
          ret = cur.t - this.t;
          for (;noIn < ret;++noIn) {
            cur[noIn + this.t] = this.am(0, o[noIn], cur, noIn, 0, this.t);
          }
          /** @type {number} */
          ret = Math.min(o.t, value);
          for (;noIn < ret;++noIn) {
            this.am(0, o[noIn], cur, noIn, 0, value - noIn);
          }
          cur.clamp();
        }, self.prototype.multiplyUpperTo = function(tail, len, cur) {
          --len;
          /** @type {number} */
          var i = cur.t = this.t + tail.t - len;
          /** @type {number} */
          cur.s = 0;
          for (;--i >= 0;) {
            /** @type {number} */
            cur[i] = 0;
          }
          /** @type {number} */
          i = Math.max(len - this.t, 0);
          for (;i < tail.t;++i) {
            cur[this.t + i - len] = this.am(len - i, tail[i], cur, 0, 0, this.t + i - len);
          }
          cur.clamp();
          cur.drShiftTo(1, cur);
        }, self.prototype.modInt = function(dataAndEvents) {
          if (dataAndEvents <= 0) {
            return 0;
          }
          /** @type {number} */
          var n = this.DV % dataAndEvents;
          /** @type {number} */
          var t = this.s < 0 ? dataAndEvents - 1 : 0;
          if (this.t > 0) {
            if (0 == n) {
              /** @type {number} */
              t = this[0] % dataAndEvents;
            } else {
              /** @type {number} */
              var i = this.t - 1;
              for (;i >= 0;--i) {
                /** @type {number} */
                t = (n * t + this[i]) % dataAndEvents;
              }
            }
          }
          return t;
        }, self.prototype.millerRabin = function(n) {
          var rvar = this.subtract(self.ONE);
          var dataAndEvents = rvar.getLowestSetBit();
          if (dataAndEvents <= 0) {
            return false;
          }
          var rreturn = rvar.shiftRight(dataAndEvents);
          if ((n = n + 1 >> 1) > nodes.length) {
            /** @type {number} */
            n = nodes.length;
          }
          var ret = parseInt();
          /** @type {number} */
          var i = 0;
          for (;i < n;++i) {
            ret.fromInt(nodes[Math.floor(Math.random() * nodes.length)]);
            var body = ret.modPow(rreturn, this);
            if (0 != body.compareTo(self.ONE) && 0 != body.compareTo(rvar)) {
              /** @type {number} */
              var c = 1;
              for (;c++ < dataAndEvents && 0 != body.compareTo(rvar);) {
                if (0 == (body = body.modPowInt(2, this)).compareTo(self.ONE)) {
                  return false;
                }
              }
              if (0 != body.compareTo(rvar)) {
                return false;
              }
            }
          }
          return true;
        }, self.prototype.clone = function() {
          var pdataOld = parseInt();
          return this.copyTo(pdataOld), pdataOld;
        }, self.prototype.intValue = function() {
          if (this.s < 0) {
            if (1 == this.t) {
              return this[0] - this.DV;
            }
            if (0 == this.t) {
              return-1;
            }
          } else {
            if (1 == this.t) {
              return this[0];
            }
            if (0 == this.t) {
              return 0;
            }
          }
          return(this[1] & (1 << 32 - this.DB) - 1) << this.DB | this[0];
        }, self.prototype.byteValue = function() {
          return 0 == this.t ? this.s : this[0] << 24 >> 24;
        }, self.prototype.shortValue = function() {
          return 0 == this.t ? this.s : this[0] << 16 >> 16;
        }, self.prototype.signum = function() {
          return this.s < 0 ? -1 : this.t <= 0 || 1 == this.t && this[0] <= 0 ? 0 : 1;
        }, self.prototype.toByteArray = function() {
          var t = this.t;
          /** @type {Array} */
          var res = new Array;
          res[0] = this.s;
          var key;
          /** @type {number} */
          var DB = this.DB - t * this.DB % 8;
          /** @type {number} */
          var resLength = 0;
          if (t-- > 0) {
            if (DB < this.DB) {
              if ((key = this[t] >> DB) != (this.s & this.DM) >> DB) {
                /** @type {number} */
                res[resLength++] = key | this.s << this.DB - DB;
              }
            }
            for (;t >= 0;) {
              if (DB < 8) {
                /** @type {number} */
                key = (this[t] & (1 << DB) - 1) << 8 - DB;
                key |= this[--t] >> (DB += this.DB - 8);
              } else {
                /** @type {number} */
                key = this[t] >> (DB -= 8) & 255;
                if (DB <= 0) {
                  DB += this.DB;
                  --t;
                }
              }
              if (0 != (128 & key)) {
                key |= -256;
              }
              if (0 == resLength) {
                if ((128 & this.s) != (128 & key)) {
                  ++resLength;
                }
              }
              if (resLength > 0 || key != this.s) {
                /** @type {(number|undefined)} */
                res[resLength++] = key;
              }
            }
          }
          return res;
        }, self.prototype.equals = function(recurring) {
          return 0 == this.compareTo(recurring);
        }, self.prototype.min = function(obj) {
          return this.compareTo(obj) < 0 ? this : obj;
        }, self.prototype.max = function(val) {
          return this.compareTo(val) > 0 ? this : val;
        }, self.prototype.and = function(sqlt) {
          var red = parseInt();
          return this.bitwiseTo(sqlt, clone, red), red;
        }, self.prototype.or = function(o) {
          var red = parseInt();
          return this.bitwiseTo(o, w, red), red;
        }, self.prototype.xor = function(sqlt) {
          var red = parseInt();
          return this.bitwiseTo(sqlt, opt_obj2, red), red;
        }, self.prototype.andNot = function(sqlt) {
          var red = parseInt();
          return this.bitwiseTo(sqlt, walk, red), red;
        }, self.prototype.not = function() {
          var t = parseInt();
          /** @type {number} */
          var func = 0;
          for (;func < this.t;++func) {
            /** @type {number} */
            t[func] = this.DM & ~this[func];
          }
          return t.t = this.t, t.s = ~this.s, t;
        }, self.prototype.shiftLeft = function(pixels) {
          var oldconfig = parseInt();
          return pixels < 0 ? this.rShiftTo(-pixels, oldconfig) : this.lShiftTo(pixels, oldconfig), oldconfig;
        }, self.prototype.shiftRight = function(dataAndEvents) {
          var oldconfig = parseInt();
          return dataAndEvents < 0 ? this.lShiftTo(-dataAndEvents, oldconfig) : this.rShiftTo(dataAndEvents, oldconfig), oldconfig;
        }, self.prototype.getLowestSetBit = function() {
          /** @type {number} */
          var key = 0;
          for (;key < this.t;++key) {
            if (0 != this[key]) {
              return key * this.DB + compileNode(this[key]);
            }
          }
          return this.s < 0 ? this.t * this.DB : -1;
        }, self.prototype.bitCount = function() {
          /** @type {number} */
          var bitCount = 0;
          /** @type {number} */
          var number = this.s & this.DM;
          /** @type {number} */
          var regTo = 0;
          for (;regTo < this.t;++regTo) {
            bitCount += promote(this[regTo] ^ number);
          }
          return bitCount;
        }, self.prototype.testBit = function(m1) {
          /** @type {number} */
          var f = Math.floor(m1 / this.DB);
          return f >= this.t ? 0 != this.s : 0 != (this[f] & 1 << m1 % this.DB);
        }, self.prototype.setBit = function(pixels) {
          return this.changeBit(pixels, w);
        }, self.prototype.clearBit = function(pixels) {
          return this.changeBit(pixels, walk);
        }, self.prototype.flipBit = function(pixels) {
          return this.changeBit(pixels, opt_obj2);
        }, self.prototype.add = function(dataAndEvents) {
          var pdataOld = parseInt();
          return this.addTo(dataAndEvents, pdataOld), pdataOld;
        }, self.prototype.subtract = function(b) {
          var pdataOld = parseInt();
          return this.subTo(b, pdataOld), pdataOld;
        }, self.prototype.multiply = function(n) {
          var border = parseInt();
          return this.multiplyTo(n, border), border;
        }, self.prototype.divide = function(m) {
          var crossScope = parseInt();
          return this.divRemTo(m, crossScope, null), crossScope;
        }, self.prototype.remainder = function(isXML) {
          var month = parseInt();
          return this.divRemTo(isXML, null, month), month;
        }, self.prototype.divideAndRemainder = function(isXML) {
          var crossScope = parseInt();
          var month = parseInt();
          return this.divRemTo(isXML, crossScope, month), new Array(crossScope, month);
        }, self.prototype.modPow = function(regex, value) {
          var i;
          var self;
          var stop = regex.bitLength();
          var border = require(1);
          if (stop <= 0) {
            return border;
          }
          /** @type {number} */
          i = stop < 18 ? 1 : stop < 48 ? 3 : stop < 144 ? 4 : stop < 768 ? 5 : 6;
          self = stop < 8 ? new D(value) : value.isEven() ? new error(value) : new Transform(value);
          /** @type {Array} */
          var data = new Array;
          /** @type {number} */
          var index = 3;
          /** @type {number} */
          var start = i - 1;
          /** @type {number} */
          var firingIndex = (1 << i) - 1;
          if (data[1] = self.convert(this), i > 1) {
            var failuresLink = parseInt();
            self.sqrTo(data[1], failuresLink);
            for (;index <= firingIndex;) {
              data[index] = parseInt();
              self.mulTo(failuresLink, data[index - 2], data[index]);
              index += 2;
            }
          }
          var associationKey;
          var pre;
          /** @type {number} */
          var j = regex.t - 1;
          /** @type {boolean} */
          var b = true;
          var h2 = parseInt();
          /** @type {number} */
          stop = checkType(regex[j]) - 1;
          for (;j >= 0;) {
            if (stop >= start) {
              /** @type {number} */
              associationKey = regex[j] >> stop - start & firingIndex;
            } else {
              /** @type {number} */
              associationKey = (regex[j] & (1 << stop + 1) - 1) << start - stop;
              if (j > 0) {
                associationKey |= regex[j - 1] >> this.DB + stop - start;
              }
            }
            /** @type {number} */
            index = i;
            for (;0 == (1 & associationKey);) {
              associationKey >>= 1;
              --index;
            }
            if ((stop -= index) < 0 && (stop += this.DB, --j), b) {
              data[associationKey].copyTo(border);
              /** @type {boolean} */
              b = false;
            } else {
              for (;index > 1;) {
                self.sqrTo(border, h2);
                self.sqrTo(h2, border);
                index -= 2;
              }
              if (index > 0) {
                self.sqrTo(border, h2);
              } else {
                pre = border;
                border = h2;
                h2 = pre;
              }
              self.mulTo(h2, data[associationKey], border);
            }
            for (;j >= 0 && 0 == (regex[j] & 1 << stop);) {
              self.sqrTo(border, h2);
              pre = border;
              border = h2;
              h2 = pre;
              if (--stop < 0) {
                /** @type {number} */
                stop = this.DB - 1;
                --j;
              }
            }
          }
          return self.revert(border);
        }, self.prototype.modInverse = function(node) {
          var e = node.isEven();
          if (this.isEven() && e || 0 == node.signum()) {
            return self.ZERO;
          }
          var config = node.clone();
          var b = this.clone();
          var needle = require(1);
          var d = require(0);
          var header = require(0);
          var m = require(1);
          for (;0 != config.signum();) {
            for (;config.isEven();) {
              config.rShiftTo(1, config);
              if (e) {
                if (!(needle.isEven() && d.isEven())) {
                  needle.addTo(this, needle);
                  d.subTo(node, d);
                }
                needle.rShiftTo(1, needle);
              } else {
                if (!d.isEven()) {
                  d.subTo(node, d);
                }
              }
              d.rShiftTo(1, d);
            }
            for (;b.isEven();) {
              b.rShiftTo(1, b);
              if (e) {
                if (!(header.isEven() && m.isEven())) {
                  header.addTo(this, header);
                  m.subTo(node, m);
                }
                header.rShiftTo(1, header);
              } else {
                if (!m.isEven()) {
                  m.subTo(node, m);
                }
              }
              m.rShiftTo(1, m);
            }
            if (config.compareTo(b) >= 0) {
              config.subTo(b, config);
              if (e) {
                needle.subTo(header, needle);
              }
              d.subTo(m, d);
            } else {
              b.subTo(config, b);
              if (e) {
                header.subTo(needle, header);
              }
              m.subTo(d, m);
            }
          }
          return 0 != b.compareTo(self.ONE) ? self.ZERO : m.compareTo(node) >= 0 ? m.subtract(node) : m.signum() < 0 ? (m.addTo(node, m), m.signum() < 0 ? m.add(node) : m) : m;
        }, self.prototype.pow = function(expectedNumberOfNonCommentArgs) {
          return this.exp(expectedNumberOfNonCommentArgs, new Filter);
        }, self.prototype.gcd = function(m) {
          var a = this.s < 0 ? this.negate() : this.clone();
          var b = m.s < 0 ? m.negate() : m.clone();
          if (a.compareTo(b) < 0) {
            var temp = a;
            a = b;
            b = temp;
          }
          var dataAndEvents = a.getLowestSetBit();
          var node = b.getLowestSetBit();
          if (node < 0) {
            return a;
          }
          if (dataAndEvents < node) {
            node = dataAndEvents;
          }
          if (node > 0) {
            a.rShiftTo(node, a);
            b.rShiftTo(node, b);
          }
          for (;a.signum() > 0;) {
            if ((dataAndEvents = a.getLowestSetBit()) > 0) {
              a.rShiftTo(dataAndEvents, a);
            }
            if ((dataAndEvents = b.getLowestSetBit()) > 0) {
              b.rShiftTo(dataAndEvents, b);
            }
            if (a.compareTo(b) >= 0) {
              a.subTo(b, a);
              a.rShiftTo(1, a);
            } else {
              b.subTo(a, b);
              b.rShiftTo(1, b);
            }
          }
          return node > 0 && b.lShiftTo(node, b), b;
        }, self.prototype.isProbablePrime = function(expectedNumberOfNonCommentArgs) {
          var i;
          var a = this.abs();
          if (1 == a.t && a[0] <= nodes[nodes.length - 1]) {
            /** @type {number} */
            i = 0;
            for (;i < nodes.length;++i) {
              if (a[0] == nodes[i]) {
                return true;
              }
            }
            return false;
          }
          if (a.isEven()) {
            return false;
          }
          /** @type {number} */
          i = 1;
          for (;i < nodes.length;) {
            var node = nodes[i];
            /** @type {number} */
            var j = i + 1;
            for (;j < nodes.length && node < total_size;) {
              node *= nodes[j++];
            }
            node = a.modInt(node);
            for (;i < j;) {
              if (node % nodes[i++] == 0) {
                return false;
              }
            }
          }
          return a.millerRabin(expectedNumberOfNonCommentArgs);
        }, self.prototype.square = function() {
          var cDigit = parseInt();
          return this.squareTo(cDigit), cDigit;
        }, self.prototype.Barrett = error, null == data) {
          var n;
          if (data = new Array, x = 0, "undefined" != typeof window && window.crypto) {
            if (window.crypto.getRandomValues) {
              /** @type {Uint8Array} */
              var buffer = new Uint8Array(32);
              window.crypto.getRandomValues(buffer);
              /** @type {number} */
              n = 0;
              for (;n < 32;++n) {
                data[x++] = buffer[n];
              }
            } else {
              if ("Netscape" == navigator.appName && navigator.appVersion < "5") {
                var string = window.crypto.random(32);
                /** @type {number} */
                n = 0;
                for (;n < string.length;++n) {
                  /** @type {number} */
                  data[x++] = 255 & string.charCodeAt(n);
                }
              }
            }
          }
          for (;x < y;) {
            /** @type {number} */
            n = Math.floor(65536 * Math.random());
            /** @type {number} */
            data[x++] = n >>> 8;
            /** @type {number} */
            data[x++] = 255 & n;
          }
          /** @type {number} */
          x = 0;
          onComplete();
        }
        /**
         * @param {Array} str
         * @return {undefined}
         */
        Type.prototype.nextBytes = function(str) {
          var strCounter;
          /** @type {number} */
          strCounter = 0;
          for (;strCounter < str.length;++strCounter) {
            str[strCounter] = finish();
          }
        };
        /**
         * @param {string} key
         * @return {undefined}
         */
        $.prototype.init = function(key) {
          var i;
          var j;
          var tempi;
          /** @type {number} */
          i = 0;
          for (;i < 256;++i) {
            /** @type {number} */
            this.S[i] = i;
          }
          /** @type {number} */
          j = 0;
          /** @type {number} */
          i = 0;
          for (;i < 256;++i) {
            /** @type {number} */
            j = j + this.S[i] + key[i % key.length] & 255;
            tempi = this.S[i];
            this.S[i] = this.S[j];
            this.S[j] = tempi;
          }
          /** @type {number} */
          this.i = 0;
          /** @type {number} */
          this.j = 0;
        };
        /**
         * @return {?}
         */
        $.prototype.next = function() {
          var opcode;
          return this.i = this.i + 1 & 255, this.j = this.j + this.S[this.i] & 255, opcode = this.S[this.i], this.S[this.i] = this.S[this.j], this.S[this.j] = opcode, this.S[opcode + this.S[this.i] & 255];
        };
        /** @type {number} */
        var y = 256;
        module.exports = {
          /** @type {function (number, number, number): undefined} */
          default : self,
          /** @type {function (number, number, number): undefined} */
          BigInteger : self,
          /** @type {function (): undefined} */
          SecureRandom : Type
        };
      }).call(this);
    }, function(module, dataAndEvents, callback) {
      (function(arg) {
        /**
         * @param {Object} context
         * @return {?}
         */
        var runInContext = function(context) {
          return context && (context.Math == Math && context);
        };
        module.exports = runInContext("object" == typeof globalThis && globalThis) || (runInContext("object" == typeof window && window) || (runInContext("object" == typeof self && self) || (runInContext("object" == typeof arg && arg) || Function("return this")())));
      }).call(this, callback(65));
    }, function(module, dataAndEvents, deepDataAndEvents) {
      module.exports = function(cursor) {
        /**
         * @param {number} name
         * @return {?}
         */
        function result(name) {
          if (modules[name]) {
            return modules[name].exports;
          }
          var m = modules[name] = {
            i : name,
            l : false,
            exports : {}
          };
          return cursor[name].call(m.exports, m, m.exports, result), m.l = true, m.exports;
        }
        var modules = {};
        return result.m = cursor, result.c = modules, result.d = function(expectedNumberOfNonCommentArgs, name, putativeSpy) {
          if (!result.o(expectedNumberOfNonCommentArgs, name)) {
            Object.defineProperty(expectedNumberOfNonCommentArgs, name, {
              enumerable : true,
              /** @type {Function} */
              get : putativeSpy
            });
          }
        }, result.r = function(expectedNumberOfNonCommentArgs) {
          if ("undefined" != typeof Symbol) {
            if (Symbol.toStringTag) {
              Object.defineProperty(expectedNumberOfNonCommentArgs, Symbol.toStringTag, {
                value : "Module"
              });
            }
          }
          Object.defineProperty(expectedNumberOfNonCommentArgs, "__esModule", {
            value : true
          });
        }, result.t = function(str, args) {
          if (1 & args && (str = result(str)), 8 & args) {
            return str;
          }
          if (4 & args && ("object" == typeof str && (str && str.__esModule))) {
            return str;
          }
          /** @type {Object} */
          var expectedNumberOfNonCommentArgs = Object.create(null);
          if (result.r(expectedNumberOfNonCommentArgs), Object.defineProperty(expectedNumberOfNonCommentArgs, "default", {
            enumerable : true,
            value : str
          }), 2 & args && "string" != typeof str) {
            var path;
            for (path in str) {
              result.d(expectedNumberOfNonCommentArgs, path, function(key) {
                return str[key];
              }.bind(null, path));
            }
          }
          return expectedNumberOfNonCommentArgs;
        }, result.n = function(c) {
          /** @type {function (): ?} */
          var value = c && c.__esModule ? function() {
            return c.default;
          } : function() {
            return c;
          };
          return result.d(value, "a", value), value;
        }, result.o = function(array, keepData) {
          return Object.prototype.hasOwnProperty.call(array, keepData);
        }, result.p = "", result(result.s = 126);
      }([function(dataAndEvents, expectedNumberOfNonCommentArgs, f) {
        (function(points) {
          /**
           * @param {number} expectedNumberOfNonCommentArgs
           * @param {Array} obj
           * @return {undefined}
           */
          function defineProperty(expectedNumberOfNonCommentArgs, obj) {
            /** @type {number} */
            var i = 0;
            for (;i < obj.length;i++) {
              var desc = obj[i];
              desc.enumerable = desc.enumerable || false;
              /** @type {boolean} */
              desc.configurable = true;
              if ("value" in desc) {
                /** @type {boolean} */
                desc.writable = true;
              }
              Object.defineProperty(expectedNumberOfNonCommentArgs, desc.key, desc);
            }
          }
          f.d(expectedNumberOfNonCommentArgs, "a", function() {
            return u;
          });
          var r = f(29);
          var word = f.n(r);
          var tree = f(1);
          var status = f(69);
          var u = function() {
            /**
             * @return {undefined}
             */
            function core() {
              !function(dataAndEvents, core) {
                if (!(dataAndEvents instanceof core)) {
                  throw new TypeError("Cannot call a class as a function");
                }
              }(this, core);
            }
            var memoized;
            var suiteView;
            return memoized = core, (suiteView = [{
              key : "stringToArrayBuffer",
              /**
               * @param {number} expectedNumberOfNonCommentArgs
               * @return {?}
               */
              value : function(expectedNumberOfNonCommentArgs) {
                return(new status.TextEncoder("utf-8")).encode(expectedNumberOfNonCommentArgs);
              }
            }, {
              key : "hexStringToArrayBufferBE",
              /**
               * @param {number} expectedNumberOfNonCommentArgs
               * @return {?}
               */
              value : function(expectedNumberOfNonCommentArgs) {
                var curr = expectedNumberOfNonCommentArgs.length;
                return points.from(expectedNumberOfNonCommentArgs.padStart(curr, "0").slice(0, curr), "hex");
              }
            }, {
              key : "hexStringToArrayBufferLE",
              /**
               * @param {number} expectedNumberOfNonCommentArgs
               * @return {?}
               */
              value : function(expectedNumberOfNonCommentArgs) {
                var curr = expectedNumberOfNonCommentArgs.length;
                var clr = points.from(expectedNumberOfNonCommentArgs.padStart(curr, "0").slice(0, curr), "hex");
                return word()(clr), clr;
              }
            }, {
              key : "signedBigIntegerToUnsigned",
              /**
               * @param {number} expectedNumberOfNonCommentArgs
               * @param {Function} object
               * @return {?}
               */
              value : function(expectedNumberOfNonCommentArgs, object) {
                return expectedNumberOfNonCommentArgs.andNot((new tree.BigInteger("-1")).shiftLeft(8 * object));
              }
            }, {
              key : "arrayBufferToHexString",
              /**
               * @param {number} expectedNumberOfNonCommentArgs
               * @return {?}
               */
              value : function(expectedNumberOfNonCommentArgs) {
                var code;
                /** @type {number} */
                var resultItems = expectedNumberOfNonCommentArgs;
                /** @type {string} */
                var output = "";
                /** @type {number} */
                var i = 0;
                for (;i < resultItems.byteLength;i++) {
                  if ((code = resultItems[i].toString(16)).length < 2) {
                    /** @type {string} */
                    code = "0" + code;
                  }
                  output += code;
                }
                return output;
              }
            }, {
              key : "byteArrayToInteger",
              /**
               * @param {number} expectedNumberOfNonCommentArgs
               * @return {?}
               */
              value : function(expectedNumberOfNonCommentArgs) {
                /** @type {number} */
                var value = 0;
                /** @type {number} */
                var unlock = 0;
                for (;unlock < 4;unlock++) {
                  /** @type {number} */
                  var shift = 8 * (3 - unlock);
                  value += (255 & expectedNumberOfNonCommentArgs[unlock]) << shift;
                }
                return value;
              }
            }, {
              key : "bufferToBigNumberLE",
              /**
               * @param {number} expectedNumberOfNonCommentArgs
               * @return {?}
               */
              value : function(expectedNumberOfNonCommentArgs) {
                var received = points.from(expectedNumberOfNonCommentArgs);
                word()(received);
                var rgb = received.toString("hex");
                return 0 === rgb.length ? new tree.BigInteger(0) : new tree.BigInteger(rgb, 16);
              }
            }, {
              key : "bufferToBigNumberBE",
              /**
               * @param {number} expectedNumberOfNonCommentArgs
               * @return {?}
               */
              value : function(expectedNumberOfNonCommentArgs) {
                var rgb = points.from(expectedNumberOfNonCommentArgs).toString("hex");
                return 0 === rgb.length ? new tree.BigInteger(0) : new tree.BigInteger(rgb, 16);
              }
            }, {
              key : "bigNumbertoBufferBE",
              /**
               * @param {number} expectedNumberOfNonCommentArgs
               * @param {number} object
               * @return {?}
               */
              value : function(expectedNumberOfNonCommentArgs, object) {
                var padStart = expectedNumberOfNonCommentArgs.toString(16);
                return points.from(padStart.padStart(2 * object, "0").slice(0, 2 * object), "hex");
              }
            }, {
              key : "bigNumbertoBufferLE",
              /**
               * @param {number} expectedNumberOfNonCommentArgs
               * @param {number} object
               * @return {?}
               */
              value : function(expectedNumberOfNonCommentArgs, object) {
                var padStart = expectedNumberOfNonCommentArgs.toString(16);
                var clr = points.from(padStart.padStart(2 * object, "0").slice(0, 2 * object), "hex");
                return word()(clr), clr;
              }
            }]) && defineProperty(memoized, suiteView), core;
          }();
        }).call(this, f(120).Buffer);
      }, function(module, dataAndEvents, deepDataAndEvents) {
        (function() {
          /**
           * @param {number} deepDataAndEvents
           * @param {number} expectedNumberOfNonCommentArgs
           * @param {number} timestep
           * @return {undefined}
           */
          function self(deepDataAndEvents, expectedNumberOfNonCommentArgs, timestep) {
            if (null != deepDataAndEvents) {
              if ("number" == typeof deepDataAndEvents) {
                this.fromNumber(deepDataAndEvents, expectedNumberOfNonCommentArgs, timestep);
              } else {
                if (null == expectedNumberOfNonCommentArgs && "string" != typeof deepDataAndEvents) {
                  this.fromString(deepDataAndEvents, 256);
                } else {
                  this.fromString(deepDataAndEvents, expectedNumberOfNonCommentArgs);
                }
              }
            }
          }
          /**
           * @return {?}
           */
          function parseInt() {
            return new self(null);
          }
          /**
           * @param {?} val
           * @return {?}
           */
          function toString(val) {
            return "0123456789abcdefghijklmnopqrstuvwxyz".charAt(val);
          }
          /**
           * @param {(Array|Uint8Array|string)} body
           * @param {number} i
           * @return {?}
           */
          function callback(body, i) {
            var y = index[body.charCodeAt(i)];
            return null == y ? -1 : y;
          }
          /**
           * @param {number} recurring
           * @return {?}
           */
          function round(recurring) {
            var charCodeToReplace = parseInt();
            return charCodeToReplace.fromInt(recurring), charCodeToReplace;
          }
          /**
           * @param {number} node
           * @return {?}
           */
          function traverseNode(node) {
            var fragment;
            /** @type {number} */
            var r = 1;
            return 0 != (fragment = node >>> 16) && (node = fragment, r += 16), 0 != (fragment = node >> 8) && (node = fragment, r += 8), 0 != (fragment = node >> 4) && (node = fragment, r += 4), 0 != (fragment = node >> 2) && (node = fragment, r += 2), 0 != (fragment = node >> 1) && (node = fragment, r += 1), r;
          }
          /**
           * @param {Array} m
           * @return {undefined}
           */
          function D(m) {
            /** @type {Array} */
            this.m = m;
          }
          /**
           * @param {Object} v
           * @return {undefined}
           */
          function Transform(v) {
            /** @type {Object} */
            this.m = v;
            this.mp = v.invDigit();
            /** @type {number} */
            this.mpl = 32767 & this.mp;
            /** @type {number} */
            this.mph = this.mp >> 15;
            /** @type {number} */
            this.um = (1 << v.DB - 15) - 1;
            /** @type {number} */
            this.mt2 = 2 * v.t;
          }
          /**
           * @param {number} dataAndEvents
           * @param {number} deepDataAndEvents
           * @return {?}
           */
          function clone(dataAndEvents, deepDataAndEvents) {
            return dataAndEvents & deepDataAndEvents;
          }
          /**
           * @param {number} dataAndEvents
           * @param {number} deepDataAndEvents
           * @return {?}
           */
          function opt_obj2(dataAndEvents, deepDataAndEvents) {
            return dataAndEvents | deepDataAndEvents;
          }
          /**
           * @param {number} a
           * @param {number} b
           * @return {?}
           */
          function walk(a, b) {
            return a ^ b;
          }
          /**
           * @param {number} c
           * @param {?} d
           * @return {?}
           */
          function w(c, d) {
            return c & ~d;
          }
          /**
           * @param {number} o
           * @return {?}
           */
          function compileNode(o) {
            if (0 == o) {
              return-1;
            }
            /** @type {number} */
            var e = 0;
            return 0 == (65535 & o) && (o >>= 16, e += 16), 0 == (255 & o) && (o >>= 8, e += 8), 0 == (15 & o) && (o >>= 4, e += 4), 0 == (3 & o) && (o >>= 2, e += 2), 0 == (1 & o) && ++e, e;
          }
          /**
           * @param {number} parent
           * @return {?}
           */
          function promote(parent) {
            /** @type {number} */
            var t = 0;
            for (;0 != parent;) {
              parent &= parent - 1;
              ++t;
            }
            return t;
          }
          /**
           * @return {undefined}
           */
          function Filter() {
          }
          /**
           * @param {?} border
           * @return {?}
           */
          function reduce(border) {
            return border;
          }
          /**
           * @param {Object} m
           * @return {undefined}
           */
          function error(m) {
            this.r2 = parseInt();
            this.q3 = parseInt();
            self.ONE.dlShiftTo(2 * m.t, this.r2);
            this.mu = this.r2.divide(m);
            /** @type {Object} */
            this.m = m;
          }
          /**
           * @return {undefined}
           */
          function onComplete() {
            var t;
            /** @type {number} */
            t = (new Date).getTime();
            data[x++] ^= 255 & t;
            data[x++] ^= t >> 8 & 255;
            data[x++] ^= t >> 16 & 255;
            data[x++] ^= t >> 24 & 255;
            if (x >= y) {
              x -= y;
            }
          }
          /**
           * @return {?}
           */
          function finish() {
            if (null == stream) {
              onComplete();
              (stream = new $).init(data);
              /** @type {number} */
              x = 0;
              for (;x < data.length;++x) {
                /** @type {number} */
                data[x] = 0;
              }
              /** @type {number} */
              x = 0;
            }
            return stream.next();
          }
          /**
           * @return {undefined}
           */
          function Type() {
          }
          /**
           * @return {undefined}
           */
          function $() {
            /** @type {number} */
            this.i = 0;
            /** @type {number} */
            this.j = 0;
            /** @type {Array} */
            this.S = new Array;
          }
          var DB;
          /** @type {boolean} */
          var Netscape = "undefined" != typeof navigator;
          if (Netscape && "Microsoft Internet Explorer" == navigator.appName) {
            /**
             * @param {number} recurring
             * @param {number} deepDataAndEvents
             * @param {Object} num
             * @param {number} noIn
             * @param {number} mayParseLabeledStatementInstead
             * @param {number} dataAndEvents
             * @return {?}
             */
            self.prototype.am = function(recurring, deepDataAndEvents, num, noIn, mayParseLabeledStatementInstead, dataAndEvents) {
              /** @type {number} */
              var a10 = 32767 & deepDataAndEvents;
              /** @type {number} */
              var x = deepDataAndEvents >> 15;
              for (;--dataAndEvents >= 0;) {
                /** @type {number} */
                var a00 = 32767 & this[recurring];
                /** @type {number} */
                var y = this[recurring++] >> 15;
                /** @type {number} */
                var f = x * a00 + y * a10;
                /** @type {number} */
                mayParseLabeledStatementInstead = ((a00 = a10 * a00 + ((32767 & f) << 15) + num[noIn] + (1073741823 & mayParseLabeledStatementInstead)) >>> 30) + (f >>> 15) + x * y + (mayParseLabeledStatementInstead >>> 30);
                /** @type {number} */
                num[noIn++] = 1073741823 & a00;
              }
              return mayParseLabeledStatementInstead;
            };
            /** @type {number} */
            DB = 30;
          } else {
            if (Netscape && "Netscape" != navigator.appName) {
              /**
               * @param {number} recurring
               * @param {number} deepDataAndEvents
               * @param {Object} num
               * @param {number} noIn
               * @param {number} mayParseLabeledStatementInstead
               * @param {number} dataAndEvents
               * @return {?}
               */
              self.prototype.am = function(recurring, deepDataAndEvents, num, noIn, mayParseLabeledStatementInstead, dataAndEvents) {
                for (;--dataAndEvents >= 0;) {
                  var sectionLength = deepDataAndEvents * this[recurring++] + num[noIn] + mayParseLabeledStatementInstead;
                  /** @type {number} */
                  mayParseLabeledStatementInstead = Math.floor(sectionLength / 67108864);
                  /** @type {number} */
                  num[noIn++] = 67108863 & sectionLength;
                }
                return mayParseLabeledStatementInstead;
              };
              /** @type {number} */
              DB = 26;
            } else {
              /**
               * @param {number} recurring
               * @param {number} deepDataAndEvents
               * @param {Object} num
               * @param {number} noIn
               * @param {number} mayParseLabeledStatementInstead
               * @param {number} dataAndEvents
               * @return {?}
               */
              self.prototype.am = function(recurring, deepDataAndEvents, num, noIn, mayParseLabeledStatementInstead, dataAndEvents) {
                /** @type {number} */
                var a10 = 16383 & deepDataAndEvents;
                /** @type {number} */
                var x = deepDataAndEvents >> 14;
                for (;--dataAndEvents >= 0;) {
                  /** @type {number} */
                  var a00 = 16383 & this[recurring];
                  /** @type {number} */
                  var y = this[recurring++] >> 14;
                  /** @type {number} */
                  var f = x * a00 + y * a10;
                  /** @type {number} */
                  mayParseLabeledStatementInstead = ((a00 = a10 * a00 + ((16383 & f) << 14) + num[noIn] + mayParseLabeledStatementInstead) >> 28) + (f >> 14) + x * y;
                  /** @type {number} */
                  num[noIn++] = 268435455 & a00;
                }
                return mayParseLabeledStatementInstead;
              };
              /** @type {number} */
              DB = 28;
            }
          }
          self.prototype.DB = DB;
          /** @type {number} */
          self.prototype.DM = (1 << DB) - 1;
          /** @type {number} */
          self.prototype.DV = 1 << DB;
          /** @type {number} */
          self.prototype.FV = Math.pow(2, 52);
          /** @type {number} */
          self.prototype.F1 = 52 - DB;
          /** @type {number} */
          self.prototype.F2 = 2 * DB - 52;
          var o;
          var i;
          /** @type {Array} */
          var index = new Array;
          /** @type {number} */
          o = "0".charCodeAt(0);
          /** @type {number} */
          i = 0;
          for (;i <= 9;++i) {
            /** @type {number} */
            index[o++] = i;
          }
          /** @type {number} */
          o = "a".charCodeAt(0);
          /** @type {number} */
          i = 10;
          for (;i < 36;++i) {
            /** @type {number} */
            index[o++] = i;
          }
          /** @type {number} */
          o = "A".charCodeAt(0);
          /** @type {number} */
          i = 10;
          for (;i < 36;++i) {
            /** @type {number} */
            index[o++] = i;
          }
          /**
           * @param {Object} x
           * @return {?}
           */
          D.prototype.convert = function(x) {
            return x.s < 0 || x.compareTo(this.m) >= 0 ? x.mod(this.m) : x;
          };
          /**
           * @param {?} border
           * @return {?}
           */
          D.prototype.revert = function(border) {
            return border;
          };
          /**
           * @param {Object} border
           * @return {undefined}
           */
          D.prototype.reduce = function(border) {
            border.divRemTo(this.m, null, border);
          };
          /**
           * @param {?} el
           * @param {?} attribute
           * @param {Object} border
           * @return {undefined}
           */
          D.prototype.mulTo = function(el, attribute, border) {
            el.multiplyTo(attribute, border);
            this.reduce(border);
          };
          /**
           * @param {?} el
           * @param {Object} border
           * @return {undefined}
           */
          D.prototype.sqrTo = function(el, border) {
            el.squareTo(border);
            this.reduce(border);
          };
          /**
           * @param {?} value
           * @return {?}
           */
          Transform.prototype.convert = function(value) {
            var b = parseInt();
            return value.abs().dlShiftTo(this.m.t, b), b.divRemTo(this.m, null, b), value.s < 0 && (b.compareTo(self.ZERO) > 0 && this.m.subTo(b, b)), b;
          };
          /**
           * @param {?} el
           * @return {?}
           */
          Transform.prototype.revert = function(el) {
            var border = parseInt();
            return el.copyTo(border), this.reduce(border), border;
          };
          /**
           * @param {Object} border
           * @return {undefined}
           */
          Transform.prototype.reduce = function(border) {
            for (;border.t <= this.mt2;) {
              /** @type {number} */
              border[border.t++] = 0;
            }
            /** @type {number} */
            var noIn = 0;
            for (;noIn < this.m.t;++noIn) {
              /** @type {number} */
              var countrySym = 32767 & border[noIn];
              /** @type {number} */
              var deepDataAndEvents = countrySym * this.mpl + ((countrySym * this.mph + (border[noIn] >> 15) * this.mpl & this.um) << 15) & border.DM;
              border[countrySym = noIn + this.m.t] += this.m.am(0, deepDataAndEvents, border, noIn, 0, this.m.t);
              for (;border[countrySym] >= border.DV;) {
                border[countrySym] -= border.DV;
                border[++countrySym]++;
              }
            }
            border.clamp();
            border.drShiftTo(this.m.t, border);
            if (border.compareTo(this.m) >= 0) {
              border.subTo(this.m, border);
            }
          };
          /**
           * @param {?} el
           * @param {?} attribute
           * @param {Object} border
           * @return {undefined}
           */
          Transform.prototype.mulTo = function(el, attribute, border) {
            el.multiplyTo(attribute, border);
            this.reduce(border);
          };
          /**
           * @param {?} el
           * @param {Object} border
           * @return {undefined}
           */
          Transform.prototype.sqrTo = function(el, border) {
            el.squareTo(border);
            this.reduce(border);
          };
          /**
           * @param {?} value
           * @return {undefined}
           */
          self.prototype.copyTo = function(value) {
            /** @type {number} */
            var kkey = this.t - 1;
            for (;kkey >= 0;--kkey) {
              value[kkey] = this[kkey];
            }
            value.t = this.t;
            value.s = this.s;
          };
          /**
           * @param {number} recurring
           * @return {undefined}
           */
          self.prototype.fromInt = function(recurring) {
            /** @type {number} */
            this.t = 1;
            /** @type {number} */
            this.s = recurring < 0 ? -1 : 0;
            if (recurring > 0) {
              /** @type {number} */
              this[0] = recurring;
            } else {
              if (recurring < -1) {
                this[0] = recurring + this.DV;
              } else {
                /** @type {number} */
                this.t = 0;
              }
            }
          };
          /**
           * @param {Array} str
           * @param {number} expectedNumberOfNonCommentArgs
           * @return {?}
           */
          self.prototype.fromString = function(str, expectedNumberOfNonCommentArgs) {
            var value;
            if (16 == expectedNumberOfNonCommentArgs) {
              /** @type {number} */
              value = 4;
            } else {
              if (8 == expectedNumberOfNonCommentArgs) {
                /** @type {number} */
                value = 3;
              } else {
                if (256 == expectedNumberOfNonCommentArgs) {
                  /** @type {number} */
                  value = 8;
                } else {
                  if (2 == expectedNumberOfNonCommentArgs) {
                    /** @type {number} */
                    value = 1;
                  } else {
                    if (32 == expectedNumberOfNonCommentArgs) {
                      /** @type {number} */
                      value = 5;
                    } else {
                      if (4 != expectedNumberOfNonCommentArgs) {
                        return void this.fromRadix(str, expectedNumberOfNonCommentArgs);
                      }
                      /** @type {number} */
                      value = 2;
                    }
                  }
                }
              }
            }
            /** @type {number} */
            this.t = 0;
            /** @type {number} */
            this.s = 0;
            var last = str.length;
            /** @type {boolean} */
            var o = false;
            /** @type {number} */
            var sum = 0;
            for (;--last >= 0;) {
              var val = 8 == value ? 255 & str[last] : callback(str, last);
              if (val < 0) {
                if ("-" == str.charAt(last)) {
                  /** @type {boolean} */
                  o = true;
                }
              } else {
                /** @type {boolean} */
                o = false;
                if (0 == sum) {
                  this[this.t++] = val;
                } else {
                  if (sum + value > this.DB) {
                    this[this.t - 1] |= (val & (1 << this.DB - sum) - 1) << sum;
                    /** @type {number} */
                    this[this.t++] = val >> this.DB - sum;
                  } else {
                    this[this.t - 1] |= val << sum;
                  }
                }
                if ((sum += value) >= this.DB) {
                  sum -= this.DB;
                }
              }
            }
            if (8 == value) {
              if (0 != (128 & str[0])) {
                /** @type {number} */
                this.s = -1;
                if (sum > 0) {
                  this[this.t - 1] |= (1 << this.DB - sum) - 1 << sum;
                }
              }
            }
            this.clamp();
            if (o) {
              self.ZERO.subTo(this, this);
            }
          };
          /**
           * @return {undefined}
           */
          self.prototype.clamp = function() {
            /** @type {number} */
            var t = this.s & this.DM;
            for (;this.t > 0 && this[this.t - 1] == t;) {
              --this.t;
            }
          };
          /**
           * @param {number} dataAndEvents
           * @param {?} o
           * @return {undefined}
           */
          self.prototype.dlShiftTo = function(dataAndEvents, o) {
            var i;
            /** @type {number} */
            i = this.t - 1;
            for (;i >= 0;--i) {
              o[i + dataAndEvents] = this[i];
            }
            /** @type {number} */
            i = dataAndEvents - 1;
            for (;i >= 0;--i) {
              /** @type {number} */
              o[i] = 0;
            }
            o.t = this.t + dataAndEvents;
            o.s = this.s;
          };
          /**
           * @param {number} dataAndEvents
           * @param {Object} o
           * @return {undefined}
           */
          self.prototype.drShiftTo = function(dataAndEvents, o) {
            /** @type {number} */
            var i = dataAndEvents;
            for (;i < this.t;++i) {
              o[i - dataAndEvents] = this[i];
            }
            /** @type {number} */
            o.t = Math.max(this.t - dataAndEvents, 0);
            o.s = this.s;
          };
          /**
           * @param {number} pos
           * @param {Object} b
           * @return {undefined}
           */
          self.prototype.lShiftTo = function(pos, b) {
            var t;
            /** @type {number} */
            var clientTop = pos % this.DB;
            /** @type {number} */
            var top = this.DB - clientTop;
            /** @type {number} */
            var kind = (1 << top) - 1;
            /** @type {number} */
            var d = Math.floor(pos / this.DB);
            /** @type {number} */
            var num = this.s << clientTop & this.DM;
            /** @type {number} */
            t = this.t - 1;
            for (;t >= 0;--t) {
              /** @type {number} */
              b[t + d + 1] = this[t] >> top | num;
              /** @type {number} */
              num = (this[t] & kind) << clientTop;
            }
            /** @type {number} */
            t = d - 1;
            for (;t >= 0;--t) {
              /** @type {number} */
              b[t] = 0;
            }
            /** @type {number} */
            b[d] = num;
            b.t = this.t + d + 1;
            b.s = this.s;
            b.clamp();
          };
          /**
           * @param {number} dataAndEvents
           * @param {?} b
           * @return {undefined}
           */
          self.prototype.rShiftTo = function(dataAndEvents, b) {
            b.s = this.s;
            /** @type {number} */
            var minX = Math.floor(dataAndEvents / this.DB);
            if (minX >= this.t) {
              /** @type {number} */
              b.t = 0;
            } else {
              /** @type {number} */
              var clientTop = dataAndEvents % this.DB;
              /** @type {number} */
              var top = this.DB - clientTop;
              /** @type {number} */
              var s = (1 << clientTop) - 1;
              /** @type {number} */
              b[0] = this[minX] >> clientTop;
              /** @type {number} */
              var maxX = minX + 1;
              for (;maxX < this.t;++maxX) {
                b[maxX - minX - 1] |= (this[maxX] & s) << top;
                /** @type {number} */
                b[maxX - minX] = this[maxX] >> clientTop;
              }
              if (clientTop > 0) {
                b[this.t - minX - 1] |= (this.s & s) << top;
              }
              /** @type {number} */
              b.t = this.t - minX;
              b.clamp();
            }
          };
          /**
           * @param {?} b
           * @param {Object} value
           * @return {undefined}
           */
          self.prototype.subTo = function(b, value) {
            /** @type {number} */
            var i = 0;
            /** @type {number} */
            var s = 0;
            /** @type {number} */
            var padLength = Math.min(b.t, this.t);
            for (;i < padLength;) {
              s += this[i] - b[i];
              /** @type {number} */
              value[i++] = s & this.DM;
              s >>= this.DB;
            }
            if (b.t < this.t) {
              s -= b.s;
              for (;i < this.t;) {
                s += this[i];
                /** @type {number} */
                value[i++] = s & this.DM;
                s >>= this.DB;
              }
              s += this.s;
            } else {
              s += this.s;
              for (;i < b.t;) {
                s -= b[i];
                /** @type {number} */
                value[i++] = s & this.DM;
                s >>= this.DB;
              }
              s -= b.s;
            }
            /** @type {number} */
            value.s = s < 0 ? -1 : 0;
            if (s < -1) {
              value[i++] = this.DV + s;
            } else {
              if (s > 0) {
                value[i++] = s;
              }
            }
            /** @type {number} */
            value.t = i;
            value.clamp();
          };
          /**
           * @param {?} node
           * @param {Object} border
           * @return {undefined}
           */
          self.prototype.multiplyTo = function(node, border) {
            var options = this.abs();
            var row = node.abs();
            var noIn = options.t;
            border.t = noIn + row.t;
            for (;--noIn >= 0;) {
              /** @type {number} */
              border[noIn] = 0;
            }
            /** @type {number} */
            noIn = 0;
            for (;noIn < row.t;++noIn) {
              border[noIn + options.t] = options.am(0, row[noIn], border, noIn, 0, options.t);
            }
            /** @type {number} */
            border.s = 0;
            border.clamp();
            if (this.s != node.s) {
              self.ZERO.subTo(border, border);
            }
          };
          /**
           * @param {Object} num
           * @return {undefined}
           */
          self.prototype.squareTo = function(num) {
            var options = this.abs();
            /** @type {number} */
            var recurring = num.t = 2 * options.t;
            for (;--recurring >= 0;) {
              /** @type {number} */
              num[recurring] = 0;
            }
            /** @type {number} */
            recurring = 0;
            for (;recurring < options.t - 1;++recurring) {
              var mayParseLabeledStatementInstead = options.am(recurring, options[recurring], num, 2 * recurring, 0, 1);
              if ((num[recurring + options.t] += options.am(recurring + 1, 2 * options[recurring], num, 2 * recurring + 1, mayParseLabeledStatementInstead, options.t - recurring - 1)) >= options.DV) {
                num[recurring + options.t] -= options.DV;
                /** @type {number} */
                num[recurring + options.t + 1] = 1;
              }
            }
            if (num.t > 0) {
              num[num.t - 1] += options.am(recurring, options[recurring], num, 2 * recurring, 0, 1);
            }
            /** @type {number} */
            num.s = 0;
            num.clamp();
          };
          /**
           * @param {?} value
           * @param {Object} crossScope
           * @param {Object} d
           * @return {?}
           */
          self.prototype.divRemTo = function(value, crossScope, d) {
            var b = value.abs();
            if (!(b.t <= 0)) {
              var a = this.abs();
              if (a.t < b.t) {
                return null != crossScope && crossScope.fromInt(0), void(null != d && this.copyTo(d));
              }
              if (null == d) {
                d = parseInt();
              }
              var m = parseInt();
              var tmp_str = this.s;
              var s = value.s;
              /** @type {number} */
              var node = this.DB - traverseNode(b[b.t - 1]);
              if (node > 0) {
                b.lShiftTo(node, m);
                a.lShiftTo(node, d);
              } else {
                b.copyTo(m);
                a.copyTo(d);
              }
              var dataAndEvents = m.t;
              var formula = m[dataAndEvents - 1];
              if (0 != formula) {
                /** @type {number} */
                var amount = formula * (1 << this.F1) + (dataAndEvents > 1 ? m[dataAndEvents - 2] >> this.F2 : 0);
                /** @type {number} */
                var pct = this.FV / amount;
                /** @type {number} */
                var t = (1 << this.F1) / amount;
                /** @type {number} */
                var cy = 1 << this.F2;
                var j = d.t;
                /** @type {number} */
                var noIn = j - dataAndEvents;
                var v = null == crossScope ? parseInt() : crossScope;
                m.dlShiftTo(noIn, v);
                if (d.compareTo(v) >= 0) {
                  /** @type {number} */
                  d[d.t++] = 1;
                  d.subTo(v, d);
                }
                self.ONE.dlShiftTo(dataAndEvents, v);
                v.subTo(m, m);
                for (;m.t < dataAndEvents;) {
                  /** @type {number} */
                  m[m.t++] = 0;
                }
                for (;--noIn >= 0;) {
                  var deepDataAndEvents = d[--j] == formula ? this.DM : Math.floor(d[j] * pct + (d[j - 1] + cy) * t);
                  if ((d[j] += m.am(0, deepDataAndEvents, d, noIn, 0, dataAndEvents)) < deepDataAndEvents) {
                    m.dlShiftTo(noIn, v);
                    d.subTo(v, d);
                    for (;d[j] < --deepDataAndEvents;) {
                      d.subTo(v, d);
                    }
                  }
                }
                if (null != crossScope) {
                  d.drShiftTo(dataAndEvents, crossScope);
                  if (tmp_str != s) {
                    self.ZERO.subTo(crossScope, crossScope);
                  }
                }
                d.t = dataAndEvents;
                d.clamp();
                if (node > 0) {
                  d.rShiftTo(node, d);
                }
                if (tmp_str < 0) {
                  self.ZERO.subTo(d, d);
                }
              }
            }
          };
          /**
           * @return {?}
           */
          self.prototype.invDigit = function() {
            if (this.t < 1) {
              return 0;
            }
            var mask = this[0];
            if (0 == (1 & mask)) {
              return 0;
            }
            /** @type {number} */
            var result = 3 & mask;
            return(result = (result = (result = (result = result * (2 - (15 & mask) * result) & 15) * (2 - (255 & mask) * result) & 255) * (2 - ((65535 & mask) * result & 65535)) & 65535) * (2 - mask * result % this.DV) % this.DV) > 0 ? this.DV - result : -result;
          };
          /**
           * @return {?}
           */
          self.prototype.isEven = function() {
            return 0 == (this.t > 0 ? 1 & this[0] : this.s);
          };
          /**
           * @param {number} expectedNumberOfNonCommentArgs
           * @param {Object} v
           * @return {?}
           */
          self.prototype.exp = function(expectedNumberOfNonCommentArgs, v) {
            if (expectedNumberOfNonCommentArgs > 4294967295 || expectedNumberOfNonCommentArgs < 1) {
              return self.ONE;
            }
            var border = parseInt();
            var failuresLink = parseInt();
            var prop = v.convert(this);
            /** @type {number} */
            var a = traverseNode(expectedNumberOfNonCommentArgs) - 1;
            prop.copyTo(border);
            for (;--a >= 0;) {
              if (v.sqrTo(border, failuresLink), (expectedNumberOfNonCommentArgs & 1 << a) > 0) {
                v.mulTo(failuresLink, prop, border);
              } else {
                var name = border;
                border = failuresLink;
                failuresLink = name;
              }
            }
            return v.revert(border);
          };
          /**
           * @param {number} expectedNumberOfNonCommentArgs
           * @return {?}
           */
          self.prototype.toString = function(expectedNumberOfNonCommentArgs) {
            if (this.s < 0) {
              return "-" + this.negate().toString(expectedNumberOfNonCommentArgs);
            }
            var right;
            if (16 == expectedNumberOfNonCommentArgs) {
              /** @type {number} */
              right = 4;
            } else {
              if (8 == expectedNumberOfNonCommentArgs) {
                /** @type {number} */
                right = 3;
              } else {
                if (2 == expectedNumberOfNonCommentArgs) {
                  /** @type {number} */
                  right = 1;
                } else {
                  if (32 == expectedNumberOfNonCommentArgs) {
                    /** @type {number} */
                    right = 5;
                  } else {
                    if (4 != expectedNumberOfNonCommentArgs) {
                      return this.toRadix(expectedNumberOfNonCommentArgs);
                    }
                    /** @type {number} */
                    right = 2;
                  }
                }
              }
            }
            var value;
            /** @type {number} */
            var UM = (1 << right) - 1;
            /** @type {boolean} */
            var escape = false;
            /** @type {string} */
            var str = "";
            var t = this.t;
            /** @type {number} */
            var left = this.DB - t * this.DB % right;
            if (t-- > 0) {
              if (left < this.DB) {
                if ((value = this[t] >> left) > 0) {
                  /** @type {boolean} */
                  escape = true;
                  str = toString(value);
                }
              }
              for (;t >= 0;) {
                if (left < right) {
                  /** @type {number} */
                  value = (this[t] & (1 << left) - 1) << right - left;
                  value |= this[--t] >> (left += this.DB - right);
                } else {
                  /** @type {number} */
                  value = this[t] >> (left -= right) & UM;
                  if (left <= 0) {
                    left += this.DB;
                    --t;
                  }
                }
                if (value > 0) {
                  /** @type {boolean} */
                  escape = true;
                }
                if (escape) {
                  str += toString(value);
                }
              }
            }
            return escape ? str : "0";
          };
          /**
           * @return {?}
           */
          self.prototype.negate = function() {
            var pdataOld = parseInt();
            return self.ZERO.subTo(this, pdataOld), pdataOld;
          };
          /**
           * @return {?}
           */
          self.prototype.abs = function() {
            return this.s < 0 ? this.negate() : this;
          };
          /**
           * @param {?} o
           * @return {?}
           */
          self.prototype.compareTo = function(o) {
            /** @type {number} */
            var s2 = this.s - o.s;
            if (0 != s2) {
              return s2;
            }
            var t = this.t;
            if (0 != (s2 = t - o.t)) {
              return this.s < 0 ? -s2 : s2;
            }
            for (;--t >= 0;) {
              if (0 != (s2 = this[t] - o[t])) {
                return s2;
              }
            }
            return 0;
          };
          /**
           * @return {?}
           */
          self.prototype.bitLength = function() {
            return this.t <= 0 ? 0 : this.DB * (this.t - 1) + traverseNode(this[this.t - 1] ^ this.s & this.DM);
          };
          /**
           * @param {?} base
           * @return {?}
           */
          self.prototype.mod = function(base) {
            var b = parseInt();
            return this.abs().divRemTo(base, null, b), this.s < 0 && (b.compareTo(self.ZERO) > 0 && base.subTo(b, b)), b;
          };
          /**
           * @param {number} expectedNumberOfNonCommentArgs
           * @param {?} value
           * @return {?}
           */
          self.prototype.modPowInt = function(expectedNumberOfNonCommentArgs, value) {
            var attempted;
            return attempted = expectedNumberOfNonCommentArgs < 256 || value.isEven() ? new D(value) : new Transform(value), this.exp(expectedNumberOfNonCommentArgs, attempted);
          };
          self.ZERO = round(0);
          self.ONE = round(1);
          /** @type {function (?): ?} */
          Filter.prototype.convert = reduce;
          /** @type {function (?): ?} */
          Filter.prototype.revert = reduce;
          /**
           * @param {?} el
           * @param {?} attribute
           * @param {Object} border
           * @return {undefined}
           */
          Filter.prototype.mulTo = function(el, attribute, border) {
            el.multiplyTo(attribute, border);
          };
          /**
           * @param {?} child
           * @param {?} el
           * @return {undefined}
           */
          Filter.prototype.sqrTo = function(child, el) {
            child.squareTo(el);
          };
          /**
           * @param {Object} value
           * @return {?}
           */
          error.prototype.convert = function(value) {
            if (value.s < 0 || value.t > 2 * this.m.t) {
              return value.mod(this.m);
            }
            if (value.compareTo(this.m) < 0) {
              return value;
            }
            var border = parseInt();
            return value.copyTo(border), this.reduce(border), border;
          };
          /**
           * @param {?} border
           * @return {?}
           */
          error.prototype.revert = function(border) {
            return border;
          };
          /**
           * @param {Object} border
           * @return {undefined}
           */
          error.prototype.reduce = function(border) {
            border.drShiftTo(this.m.t - 1, this.r2);
            if (border.t > this.m.t + 1) {
              border.t = this.m.t + 1;
              border.clamp();
            }
            this.mu.multiplyUpperTo(this.r2, this.m.t + 1, this.q3);
            this.m.multiplyLowerTo(this.q3, this.m.t + 1, this.r2);
            for (;border.compareTo(this.r2) < 0;) {
              border.dAddOffset(1, this.m.t + 1);
            }
            border.subTo(this.r2, border);
            for (;border.compareTo(this.m) >= 0;) {
              border.subTo(this.m, border);
            }
          };
          /**
           * @param {?} el
           * @param {?} attribute
           * @param {Object} border
           * @return {undefined}
           */
          error.prototype.mulTo = function(el, attribute, border) {
            el.multiplyTo(attribute, border);
            this.reduce(border);
          };
          /**
           * @param {?} el
           * @param {Object} border
           * @return {undefined}
           */
          error.prototype.sqrTo = function(el, border) {
            el.squareTo(border);
            this.reduce(border);
          };
          var stream;
          var data;
          var x;
          /** @type {Array} */
          var nodes = [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541, 547, 557, 
          563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911, 919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997];
          /** @type {number} */
          var total_size = (1 << 26) / nodes[nodes.length - 1];
          if (self.prototype.chunkSize = function(expectedNumberOfNonCommentArgs) {
            return Math.floor(Math.LN2 * this.DB / Math.log(expectedNumberOfNonCommentArgs));
          }, self.prototype.toRadix = function(expectedNumberOfNonCommentArgs) {
            if (null == expectedNumberOfNonCommentArgs && (expectedNumberOfNonCommentArgs = 10), 0 == this.signum() || (expectedNumberOfNonCommentArgs < 2 || expectedNumberOfNonCommentArgs > 36)) {
              return "0";
            }
            var exponent = this.chunkSize(expectedNumberOfNonCommentArgs);
            /** @type {number} */
            var recurring = Math.pow(expectedNumberOfNonCommentArgs, exponent);
            var pdataOld = round(recurring);
            var crossScope = parseInt();
            var month = parseInt();
            /** @type {string} */
            var buttons = "";
            this.divRemTo(pdataOld, crossScope, month);
            for (;crossScope.signum() > 0;) {
              buttons = (recurring + month.intValue()).toString(expectedNumberOfNonCommentArgs).substr(1) + buttons;
              crossScope.divRemTo(pdataOld, crossScope, month);
            }
            return month.intValue().toString(expectedNumberOfNonCommentArgs) + buttons;
          }, self.prototype.fromRadix = function(body, expectedNumberOfNonCommentArgs) {
            this.fromInt(0);
            if (null == expectedNumberOfNonCommentArgs) {
              /** @type {number} */
              expectedNumberOfNonCommentArgs = 10;
            }
            var exponent = this.chunkSize(expectedNumberOfNonCommentArgs);
            /** @type {number} */
            var dataAndEvents = Math.pow(expectedNumberOfNonCommentArgs, exponent);
            /** @type {boolean} */
            var o = false;
            /** @type {number} */
            var bits = 0;
            /** @type {number} */
            var dns = 0;
            /** @type {number} */
            var j = 0;
            for (;j < body.length;++j) {
              var host = callback(body, j);
              if (host < 0) {
                if ("-" == body.charAt(j)) {
                  if (0 == this.signum()) {
                    /** @type {boolean} */
                    o = true;
                  }
                }
              } else {
                dns = expectedNumberOfNonCommentArgs * dns + host;
                if (++bits >= exponent) {
                  this.dMultiply(dataAndEvents);
                  this.dAddOffset(dns, 0);
                  /** @type {number} */
                  bits = 0;
                  /** @type {number} */
                  dns = 0;
                }
              }
            }
            if (bits > 0) {
              this.dMultiply(Math.pow(expectedNumberOfNonCommentArgs, bits));
              this.dAddOffset(dns, 0);
            }
            if (o) {
              self.ZERO.subTo(this, this);
            }
          }, self.prototype.fromNumber = function(deepDataAndEvents, expectedNumberOfNonCommentArgs, n) {
            if ("number" == typeof expectedNumberOfNonCommentArgs) {
              if (deepDataAndEvents < 2) {
                this.fromInt(1);
              } else {
                this.fromNumber(deepDataAndEvents, n);
                if (!this.testBit(deepDataAndEvents - 1)) {
                  this.bitwiseTo(self.ONE.shiftLeft(deepDataAndEvents - 1), opt_obj2, this);
                }
                if (this.isEven()) {
                  this.dAddOffset(1, 0);
                }
                for (;!this.isProbablePrime(expectedNumberOfNonCommentArgs);) {
                  this.dAddOffset(2, 0);
                  if (this.bitLength() > deepDataAndEvents) {
                    this.subTo(self.ONE.shiftLeft(deepDataAndEvents - 1), this);
                  }
                }
              }
            } else {
              /** @type {Array} */
              var simple = new Array;
              /** @type {number} */
              var o = 7 & deepDataAndEvents;
              /** @type {number} */
              simple.length = 1 + (deepDataAndEvents >> 3);
              expectedNumberOfNonCommentArgs.nextBytes(simple);
              if (o > 0) {
                simple[0] &= (1 << o) - 1;
              } else {
                /** @type {number} */
                simple[0] = 0;
              }
              this.fromString(simple, 256);
            }
          }, self.prototype.bitwiseTo = function(t, f, r) {
            var i;
            var x;
            /** @type {number} */
            var lastLine = Math.min(t.t, this.t);
            /** @type {number} */
            i = 0;
            for (;i < lastLine;++i) {
              r[i] = f(this[i], t[i]);
            }
            if (t.t < this.t) {
              /** @type {number} */
              x = t.s & this.DM;
              /** @type {number} */
              i = lastLine;
              for (;i < this.t;++i) {
                r[i] = f(this[i], x);
              }
              r.t = this.t;
            } else {
              /** @type {number} */
              x = this.s & this.DM;
              /** @type {number} */
              i = lastLine;
              for (;i < t.t;++i) {
                r[i] = f(x, t[i]);
              }
              r.t = t.t;
            }
            r.s = f(this.s, t.s);
            r.clamp();
          }, self.prototype.changeBit = function(pixels, f) {
            var n = self.ONE.shiftLeft(pixels);
            return this.bitwiseTo(n, f, n), n;
          }, self.prototype.addTo = function(dataAndEvents, value) {
            /** @type {number} */
            var top = 0;
            /** @type {number} */
            var y = 0;
            /** @type {number} */
            var wTop = Math.min(dataAndEvents.t, this.t);
            for (;top < wTop;) {
              y += this[top] + dataAndEvents[top];
              /** @type {number} */
              value[top++] = y & this.DM;
              y >>= this.DB;
            }
            if (dataAndEvents.t < this.t) {
              y += dataAndEvents.s;
              for (;top < this.t;) {
                y += this[top];
                /** @type {number} */
                value[top++] = y & this.DM;
                y >>= this.DB;
              }
              y += this.s;
            } else {
              y += this.s;
              for (;top < dataAndEvents.t;) {
                y += dataAndEvents[top];
                /** @type {number} */
                value[top++] = y & this.DM;
                y >>= this.DB;
              }
              y += dataAndEvents.s;
            }
            /** @type {number} */
            value.s = y < 0 ? -1 : 0;
            if (y > 0) {
              value[top++] = y;
            } else {
              if (y < -1) {
                value[top++] = this.DV + y;
              }
            }
            /** @type {number} */
            value.t = top;
            value.clamp();
          }, self.prototype.dMultiply = function(dataAndEvents) {
            this[this.t] = this.am(0, dataAndEvents - 1, this, 0, 0, this.t);
            ++this.t;
            this.clamp();
          }, self.prototype.dAddOffset = function(expectedNumberOfNonCommentArgs, mayParseLabeledStatementInstead) {
            if (0 != expectedNumberOfNonCommentArgs) {
              for (;this.t <= mayParseLabeledStatementInstead;) {
                /** @type {number} */
                this[this.t++] = 0;
              }
              this[mayParseLabeledStatementInstead] += expectedNumberOfNonCommentArgs;
              for (;this[mayParseLabeledStatementInstead] >= this.DV;) {
                this[mayParseLabeledStatementInstead] -= this.DV;
                if (++mayParseLabeledStatementInstead >= this.t) {
                  /** @type {number} */
                  this[this.t++] = 0;
                }
                ++this[mayParseLabeledStatementInstead];
              }
            }
          }, self.prototype.multiplyLowerTo = function(o, value, cur) {
            var ret;
            /** @type {number} */
            var noIn = Math.min(this.t + o.t, value);
            /** @type {number} */
            cur.s = 0;
            /** @type {number} */
            cur.t = noIn;
            for (;noIn > 0;) {
              /** @type {number} */
              cur[--noIn] = 0;
            }
            /** @type {number} */
            ret = cur.t - this.t;
            for (;noIn < ret;++noIn) {
              cur[noIn + this.t] = this.am(0, o[noIn], cur, noIn, 0, this.t);
            }
            /** @type {number} */
            ret = Math.min(o.t, value);
            for (;noIn < ret;++noIn) {
              this.am(0, o[noIn], cur, noIn, 0, value - noIn);
            }
            cur.clamp();
          }, self.prototype.multiplyUpperTo = function(b, t, cur) {
            --t;
            /** @type {number} */
            var j = cur.t = this.t + b.t - t;
            /** @type {number} */
            cur.s = 0;
            for (;--j >= 0;) {
              /** @type {number} */
              cur[j] = 0;
            }
            /** @type {number} */
            j = Math.max(t - this.t, 0);
            for (;j < b.t;++j) {
              cur[this.t + j - t] = this.am(t - j, b[j], cur, 0, 0, this.t + j - t);
            }
            cur.clamp();
            cur.drShiftTo(1, cur);
          }, self.prototype.modInt = function(dataAndEvents) {
            if (dataAndEvents <= 0) {
              return 0;
            }
            /** @type {number} */
            var n = this.DV % dataAndEvents;
            /** @type {number} */
            var t = this.s < 0 ? dataAndEvents - 1 : 0;
            if (this.t > 0) {
              if (0 == n) {
                /** @type {number} */
                t = this[0] % dataAndEvents;
              } else {
                /** @type {number} */
                var i = this.t - 1;
                for (;i >= 0;--i) {
                  /** @type {number} */
                  t = (n * t + this[i]) % dataAndEvents;
                }
              }
            }
            return t;
          }, self.prototype.millerRabin = function(n) {
            var rvar = this.subtract(self.ONE);
            var dataAndEvents = rvar.getLowestSetBit();
            if (dataAndEvents <= 0) {
              return false;
            }
            var rreturn = rvar.shiftRight(dataAndEvents);
            if ((n = n + 1 >> 1) > nodes.length) {
              /** @type {number} */
              n = nodes.length;
            }
            var ret = parseInt();
            /** @type {number} */
            var i = 0;
            for (;i < n;++i) {
              ret.fromInt(nodes[Math.floor(Math.random() * nodes.length)]);
              var body = ret.modPow(rreturn, this);
              if (0 != body.compareTo(self.ONE) && 0 != body.compareTo(rvar)) {
                /** @type {number} */
                var c = 1;
                for (;c++ < dataAndEvents && 0 != body.compareTo(rvar);) {
                  if (0 == (body = body.modPowInt(2, this)).compareTo(self.ONE)) {
                    return false;
                  }
                }
                if (0 != body.compareTo(rvar)) {
                  return false;
                }
              }
            }
            return true;
          }, self.prototype.clone = function() {
            var pdataOld = parseInt();
            return this.copyTo(pdataOld), pdataOld;
          }, self.prototype.intValue = function() {
            if (this.s < 0) {
              if (1 == this.t) {
                return this[0] - this.DV;
              }
              if (0 == this.t) {
                return-1;
              }
            } else {
              if (1 == this.t) {
                return this[0];
              }
              if (0 == this.t) {
                return 0;
              }
            }
            return(this[1] & (1 << 32 - this.DB) - 1) << this.DB | this[0];
          }, self.prototype.byteValue = function() {
            return 0 == this.t ? this.s : this[0] << 24 >> 24;
          }, self.prototype.shortValue = function() {
            return 0 == this.t ? this.s : this[0] << 16 >> 16;
          }, self.prototype.signum = function() {
            return this.s < 0 ? -1 : this.t <= 0 || 1 == this.t && this[0] <= 0 ? 0 : 1;
          }, self.prototype.toByteArray = function() {
            var t = this.t;
            /** @type {Array} */
            var res = new Array;
            res[0] = this.s;
            var key;
            /** @type {number} */
            var DB = this.DB - t * this.DB % 8;
            /** @type {number} */
            var resLength = 0;
            if (t-- > 0) {
              if (DB < this.DB) {
                if ((key = this[t] >> DB) != (this.s & this.DM) >> DB) {
                  /** @type {number} */
                  res[resLength++] = key | this.s << this.DB - DB;
                }
              }
              for (;t >= 0;) {
                if (DB < 8) {
                  /** @type {number} */
                  key = (this[t] & (1 << DB) - 1) << 8 - DB;
                  key |= this[--t] >> (DB += this.DB - 8);
                } else {
                  /** @type {number} */
                  key = this[t] >> (DB -= 8) & 255;
                  if (DB <= 0) {
                    DB += this.DB;
                    --t;
                  }
                }
                if (0 != (128 & key)) {
                  key |= -256;
                }
                if (0 == resLength) {
                  if ((128 & this.s) != (128 & key)) {
                    ++resLength;
                  }
                }
                if (resLength > 0 || key != this.s) {
                  /** @type {(number|undefined)} */
                  res[resLength++] = key;
                }
              }
            }
            return res;
          }, self.prototype.equals = function(recurring) {
            return 0 == this.compareTo(recurring);
          }, self.prototype.min = function(obj) {
            return this.compareTo(obj) < 0 ? this : obj;
          }, self.prototype.max = function(val) {
            return this.compareTo(val) > 0 ? this : val;
          }, self.prototype.and = function(sqlt) {
            var red = parseInt();
            return this.bitwiseTo(sqlt, clone, red), red;
          }, self.prototype.or = function(o) {
            var red = parseInt();
            return this.bitwiseTo(o, opt_obj2, red), red;
          }, self.prototype.xor = function(sqlt) {
            var red = parseInt();
            return this.bitwiseTo(sqlt, walk, red), red;
          }, self.prototype.andNot = function(sqlt) {
            var red = parseInt();
            return this.bitwiseTo(sqlt, w, red), red;
          }, self.prototype.not = function() {
            var t = parseInt();
            /** @type {number} */
            var func = 0;
            for (;func < this.t;++func) {
              /** @type {number} */
              t[func] = this.DM & ~this[func];
            }
            return t.t = this.t, t.s = ~this.s, t;
          }, self.prototype.shiftLeft = function(pixels) {
            var oldconfig = parseInt();
            return pixels < 0 ? this.rShiftTo(-pixels, oldconfig) : this.lShiftTo(pixels, oldconfig), oldconfig;
          }, self.prototype.shiftRight = function(dataAndEvents) {
            var oldconfig = parseInt();
            return dataAndEvents < 0 ? this.lShiftTo(-dataAndEvents, oldconfig) : this.rShiftTo(dataAndEvents, oldconfig), oldconfig;
          }, self.prototype.getLowestSetBit = function() {
            /** @type {number} */
            var key = 0;
            for (;key < this.t;++key) {
              if (0 != this[key]) {
                return key * this.DB + compileNode(this[key]);
              }
            }
            return this.s < 0 ? this.t * this.DB : -1;
          }, self.prototype.bitCount = function() {
            /** @type {number} */
            var bitCount = 0;
            /** @type {number} */
            var number = this.s & this.DM;
            /** @type {number} */
            var regTo = 0;
            for (;regTo < this.t;++regTo) {
              bitCount += promote(this[regTo] ^ number);
            }
            return bitCount;
          }, self.prototype.testBit = function(m1) {
            /** @type {number} */
            var f = Math.floor(m1 / this.DB);
            return f >= this.t ? 0 != this.s : 0 != (this[f] & 1 << m1 % this.DB);
          }, self.prototype.setBit = function(pixels) {
            return this.changeBit(pixels, opt_obj2);
          }, self.prototype.clearBit = function(pixels) {
            return this.changeBit(pixels, w);
          }, self.prototype.flipBit = function(pixels) {
            return this.changeBit(pixels, walk);
          }, self.prototype.add = function(dataAndEvents) {
            var pdataOld = parseInt();
            return this.addTo(dataAndEvents, pdataOld), pdataOld;
          }, self.prototype.subtract = function(b) {
            var pdataOld = parseInt();
            return this.subTo(b, pdataOld), pdataOld;
          }, self.prototype.multiply = function(n) {
            var border = parseInt();
            return this.multiplyTo(n, border), border;
          }, self.prototype.divide = function(m) {
            var crossScope = parseInt();
            return this.divRemTo(m, crossScope, null), crossScope;
          }, self.prototype.remainder = function(isXML) {
            var month = parseInt();
            return this.divRemTo(isXML, null, month), month;
          }, self.prototype.divideAndRemainder = function(isXML) {
            var crossScope = parseInt();
            var month = parseInt();
            return this.divRemTo(isXML, crossScope, month), new Array(crossScope, month);
          }, self.prototype.modPow = function(regex, value) {
            var i;
            var self;
            var stop = regex.bitLength();
            var border = round(1);
            if (stop <= 0) {
              return border;
            }
            /** @type {number} */
            i = stop < 18 ? 1 : stop < 48 ? 3 : stop < 144 ? 4 : stop < 768 ? 5 : 6;
            self = stop < 8 ? new D(value) : value.isEven() ? new error(value) : new Transform(value);
            /** @type {Array} */
            var data = new Array;
            /** @type {number} */
            var index = 3;
            /** @type {number} */
            var start = i - 1;
            /** @type {number} */
            var firingIndex = (1 << i) - 1;
            if (data[1] = self.convert(this), i > 1) {
              var failuresLink = parseInt();
              self.sqrTo(data[1], failuresLink);
              for (;index <= firingIndex;) {
                data[index] = parseInt();
                self.mulTo(failuresLink, data[index - 2], data[index]);
                index += 2;
              }
            }
            var associationKey;
            var pre;
            /** @type {number} */
            var key = regex.t - 1;
            /** @type {boolean} */
            var b = true;
            var h2 = parseInt();
            /** @type {number} */
            stop = traverseNode(regex[key]) - 1;
            for (;key >= 0;) {
              if (stop >= start) {
                /** @type {number} */
                associationKey = regex[key] >> stop - start & firingIndex;
              } else {
                /** @type {number} */
                associationKey = (regex[key] & (1 << stop + 1) - 1) << start - stop;
                if (key > 0) {
                  associationKey |= regex[key - 1] >> this.DB + stop - start;
                }
              }
              /** @type {number} */
              index = i;
              for (;0 == (1 & associationKey);) {
                associationKey >>= 1;
                --index;
              }
              if ((stop -= index) < 0 && (stop += this.DB, --key), b) {
                data[associationKey].copyTo(border);
                /** @type {boolean} */
                b = false;
              } else {
                for (;index > 1;) {
                  self.sqrTo(border, h2);
                  self.sqrTo(h2, border);
                  index -= 2;
                }
                if (index > 0) {
                  self.sqrTo(border, h2);
                } else {
                  pre = border;
                  border = h2;
                  h2 = pre;
                }
                self.mulTo(h2, data[associationKey], border);
              }
              for (;key >= 0 && 0 == (regex[key] & 1 << stop);) {
                self.sqrTo(border, h2);
                pre = border;
                border = h2;
                h2 = pre;
                if (--stop < 0) {
                  /** @type {number} */
                  stop = this.DB - 1;
                  --key;
                }
              }
            }
            return self.revert(border);
          }, self.prototype.modInverse = function(node) {
            var e = node.isEven();
            if (this.isEven() && e || 0 == node.signum()) {
              return self.ZERO;
            }
            var config = node.clone();
            var b = this.clone();
            var needle = round(1);
            var elem = round(0);
            var header = round(0);
            var target = round(1);
            for (;0 != config.signum();) {
              for (;config.isEven();) {
                config.rShiftTo(1, config);
                if (e) {
                  if (!(needle.isEven() && elem.isEven())) {
                    needle.addTo(this, needle);
                    elem.subTo(node, elem);
                  }
                  needle.rShiftTo(1, needle);
                } else {
                  if (!elem.isEven()) {
                    elem.subTo(node, elem);
                  }
                }
                elem.rShiftTo(1, elem);
              }
              for (;b.isEven();) {
                b.rShiftTo(1, b);
                if (e) {
                  if (!(header.isEven() && target.isEven())) {
                    header.addTo(this, header);
                    target.subTo(node, target);
                  }
                  header.rShiftTo(1, header);
                } else {
                  if (!target.isEven()) {
                    target.subTo(node, target);
                  }
                }
                target.rShiftTo(1, target);
              }
              if (config.compareTo(b) >= 0) {
                config.subTo(b, config);
                if (e) {
                  needle.subTo(header, needle);
                }
                elem.subTo(target, elem);
              } else {
                b.subTo(config, b);
                if (e) {
                  header.subTo(needle, header);
                }
                target.subTo(elem, target);
              }
            }
            return 0 != b.compareTo(self.ONE) ? self.ZERO : target.compareTo(node) >= 0 ? target.subtract(node) : target.signum() < 0 ? (target.addTo(node, target), target.signum() < 0 ? target.add(node) : target) : target;
          }, self.prototype.pow = function(expectedNumberOfNonCommentArgs) {
            return this.exp(expectedNumberOfNonCommentArgs, new Filter);
          }, self.prototype.gcd = function(m) {
            var a = this.s < 0 ? this.negate() : this.clone();
            var b = m.s < 0 ? m.negate() : m.clone();
            if (a.compareTo(b) < 0) {
              var temp = a;
              a = b;
              b = temp;
            }
            var dataAndEvents = a.getLowestSetBit();
            var node = b.getLowestSetBit();
            if (node < 0) {
              return a;
            }
            if (dataAndEvents < node) {
              node = dataAndEvents;
            }
            if (node > 0) {
              a.rShiftTo(node, a);
              b.rShiftTo(node, b);
            }
            for (;a.signum() > 0;) {
              if ((dataAndEvents = a.getLowestSetBit()) > 0) {
                a.rShiftTo(dataAndEvents, a);
              }
              if ((dataAndEvents = b.getLowestSetBit()) > 0) {
                b.rShiftTo(dataAndEvents, b);
              }
              if (a.compareTo(b) >= 0) {
                a.subTo(b, a);
                a.rShiftTo(1, a);
              } else {
                b.subTo(a, b);
                b.rShiftTo(1, b);
              }
            }
            return node > 0 && b.lShiftTo(node, b), b;
          }, self.prototype.isProbablePrime = function(expectedNumberOfNonCommentArgs) {
            var i;
            var a = this.abs();
            if (1 == a.t && a[0] <= nodes[nodes.length - 1]) {
              /** @type {number} */
              i = 0;
              for (;i < nodes.length;++i) {
                if (a[0] == nodes[i]) {
                  return true;
                }
              }
              return false;
            }
            if (a.isEven()) {
              return false;
            }
            /** @type {number} */
            i = 1;
            for (;i < nodes.length;) {
              var node = nodes[i];
              /** @type {number} */
              var j = i + 1;
              for (;j < nodes.length && node < total_size;) {
                node *= nodes[j++];
              }
              node = a.modInt(node);
              for (;i < j;) {
                if (node % nodes[i++] == 0) {
                  return false;
                }
              }
            }
            return a.millerRabin(expectedNumberOfNonCommentArgs);
          }, self.prototype.square = function() {
            var cDigit = parseInt();
            return this.squareTo(cDigit), cDigit;
          }, self.prototype.Barrett = error, null == data) {
            var n;
            if (data = new Array, x = 0, "undefined" != typeof window && window.crypto) {
              if (window.crypto.getRandomValues) {
                /** @type {Uint8Array} */
                var buffer = new Uint8Array(32);
                window.crypto.getRandomValues(buffer);
                /** @type {number} */
                n = 0;
                for (;n < 32;++n) {
                  data[x++] = buffer[n];
                }
              } else {
                if ("Netscape" == navigator.appName && navigator.appVersion < "5") {
                  var string = window.crypto.random(32);
                  /** @type {number} */
                  n = 0;
                  for (;n < string.length;++n) {
                    /** @type {number} */
                    data[x++] = 255 & string.charCodeAt(n);
                  }
                }
              }
            }
            for (;x < y;) {
              /** @type {number} */
              n = Math.floor(65536 * Math.random());
              /** @type {number} */
              data[x++] = n >>> 8;
              /** @type {number} */
              data[x++] = 255 & n;
            }
            /** @type {number} */
            x = 0;
            onComplete();
          }
          /**
           * @param {Array} str
           * @return {undefined}
           */
          Type.prototype.nextBytes = function(str) {
            var strCounter;
            /** @type {number} */
            strCounter = 0;
            for (;strCounter < str.length;++strCounter) {
              str[strCounter] = finish();
            }
          };
          /**
           * @param {string} key
           * @return {undefined}
           */
          $.prototype.init = function(key) {
            var i;
            var j;
            var tempi;
            /** @type {number} */
            i = 0;
            for (;i < 256;++i) {
              /** @type {number} */
              this.S[i] = i;
            }
            /** @type {number} */
            j = 0;
            /** @type {number} */
            i = 0;
            for (;i < 256;++i) {
              /** @type {number} */
              j = j + this.S[i] + key[i % key.length] & 255;
              tempi = this.S[i];
              this.S[i] = this.S[j];
              this.S[j] = tempi;
            }
            /** @type {number} */
            this.i = 0;
            /** @type {number} */
            this.j = 0;
          };
          /**
           * @return {?}
           */
          $.prototype.next = function() {
            var opcode;
            return this.i = this.i + 1 & 255, this.j = this.j + this.S[this.i] & 255, opcode = this.S[this.i], this.S[this.i] = this.S[this.j], this.S[this.j] = opcode, this.S[opcode + this.S[this.i] & 255];
          };
          /** @type {number} */
          var y = 256;
          module.exports = {
            /** @type {function (number, number, number): undefined} */
            default : self,
            /** @type {function (number, number, number): undefined} */
            BigInteger : self,
            /** @type {function (): undefined} */
            SecureRandom : Type
          };
        }).call(this);
      }, function(module, global, $sanitize) {
        /**
         * @param {string} __
         * @return {?}
         */
        function callback(__) {
          throw __;
        }
        /**
         * @param {(Array|Element)} root
         * @param {number} item
         * @param {boolean} i
         * @return {?}
         */
        function has(root, item, i) {
          if (4 !== item.length) {
            callback(new options.exception.invalid("invalid aes block size"));
          }
          var stack = root.d[i];
          /** @type {number} */
          var result = item[0] ^ stack[0];
          /** @type {number} */
          var data = item[i ? 3 : 1] ^ stack[1];
          /** @type {number} */
          var selector = item[2] ^ stack[2];
          /** @type {number} */
          item = item[i ? 1 : 3] ^ stack[3];
          var obj;
          var tmp;
          var until;
          var left;
          /** @type {number} */
          var right = stack.length / 4 - 2;
          /** @type {number} */
          var sp = 4;
          /** @type {Array} */
          var pos = [0, 0, 0, 0];
          root = (obj = root.A[i])[0];
          var orig = obj[1];
          var existingFunction = obj[2];
          var valueDate = obj[3];
          var temp = obj[4];
          /** @type {number} */
          left = 0;
          for (;left < right;left++) {
            /** @type {number} */
            obj = root[result >>> 24] ^ orig[data >> 16 & 255] ^ existingFunction[selector >> 8 & 255] ^ valueDate[255 & item] ^ stack[sp];
            /** @type {number} */
            tmp = root[data >>> 24] ^ orig[selector >> 16 & 255] ^ existingFunction[item >> 8 & 255] ^ valueDate[255 & result] ^ stack[sp + 1];
            /** @type {number} */
            until = root[selector >>> 24] ^ orig[item >> 16 & 255] ^ existingFunction[result >> 8 & 255] ^ valueDate[255 & data] ^ stack[sp + 2];
            /** @type {number} */
            item = root[item >>> 24] ^ orig[result >> 16 & 255] ^ existingFunction[data >> 8 & 255] ^ valueDate[255 & selector] ^ stack[sp + 3];
            sp += 4;
            /** @type {number} */
            result = obj;
            /** @type {number} */
            data = tmp;
            /** @type {number} */
            selector = until;
          }
          /** @type {number} */
          left = 0;
          for (;4 > left;left++) {
            /** @type {number} */
            pos[i ? 3 & -left : left] = temp[result >>> 24] << 24 ^ temp[data >> 16 & 255] << 16 ^ temp[selector >> 8 & 255] << 8 ^ temp[255 & item] ^ stack[sp++];
            /** @type {number} */
            obj = result;
            /** @type {number} */
            result = data;
            /** @type {number} */
            data = selector;
            /** @type {number} */
            selector = item;
            /** @type {number} */
            item = obj;
          }
          return pos;
        }
        /**
         * @param {string} name
         * @param {number} value
         * @return {undefined}
         */
        function declare(name, value) {
          var i;
          var original = options.random.P[name];
          /** @type {Array} */
          var functions = [];
          for (i in original) {
            if (original.hasOwnProperty(i)) {
              functions.push(original[i]);
            }
          }
          /** @type {number} */
          i = 0;
          for (;i < functions.length;i++) {
            functions[i](value);
          }
        }
        /**
         * @param {number} recurring
         * @return {undefined}
         */
        function run(recurring) {
          if ("undefined" != typeof window && (window.performance && "function" == typeof window.performance.now)) {
            options.random.addEntropy(window.performance.now(), recurring, "loadtime");
          } else {
            options.random.addEntropy((new Date).valueOf(), recurring, "loadtime");
          }
        }
        /**
         * @param {Object} args
         * @return {undefined}
         */
        function go(args) {
          args.d = slice(args).concat(slice(args));
          args.Q = new options.cipher.aes(args.d);
        }
        /**
         * @param {Object} key
         * @return {?}
         */
        function slice(key) {
          /** @type {number} */
          var i = 0;
          for (;4 > i && (key.q[i] = key.q[i] + 1 | 0, !key.q[i]);i++) {
          }
          return key.Q.encrypt(key.q);
        }
        /**
         * @param {?} first
         * @param {?} matcherFunction
         * @return {?}
         */
        function agree(first, matcherFunction) {
          return function() {
            matcherFunction.apply(first, arguments);
          };
        }
        var ret;
        var target = void 0;
        /** @type {boolean} */
        var pdataCur = true;
        /** @type {boolean} */
        var FALSE = false;
        var options = {
          cipher : {},
          hash : {},
          keyexchange : {},
          mode : {},
          misc : {},
          codec : {},
          exception : {
            /**
             * @param {string} message
             * @return {undefined}
             */
            corrupt : function(message) {
              /**
               * @return {?}
               */
              this.toString = function() {
                return "CORRUPT: " + this.message;
              };
              /** @type {string} */
              this.message = message;
            },
            /**
             * @param {string} message
             * @return {undefined}
             */
            invalid : function(message) {
              /**
               * @return {?}
               */
              this.toString = function() {
                return "INVALID: " + this.message;
              };
              /** @type {string} */
              this.message = message;
            },
            /**
             * @param {string} message
             * @return {undefined}
             */
            bug : function(message) {
              /**
               * @return {?}
               */
              this.toString = function() {
                return "BUG: " + this.message;
              };
              /** @type {string} */
              this.message = message;
            },
            /**
             * @param {string} message
             * @return {undefined}
             */
            notReady : function(message) {
              /**
               * @return {?}
               */
              this.toString = function() {
                return "NOT READY: " + this.message;
              };
              /** @type {string} */
              this.message = message;
            }
          }
        };
        if (module.exports) {
          module.exports = options;
        }
        if (!(void 0 === (ret = function() {
          return options;
        }.apply(global, [])))) {
          module.exports = ret;
        }
        /**
         * @param {number} start
         * @return {undefined}
         */
        options.cipher.aes = function(start) {
          if (!this.A[0][0][0]) {
            this.J();
          }
          var i;
          var value;
          var o;
          var obj;
          var path = this.A[0][4];
          var seen = this.A[1];
          /** @type {number} */
          var c = 1;
          if (4 !== (i = start.length)) {
            if (6 !== i) {
              if (8 !== i) {
                callback(new options.exception.invalid("invalid aes key size"));
              }
            }
          }
          /** @type {Array} */
          this.d = [o = start.slice(0), obj = []];
          start = i;
          for (;start < 4 * i + 28;start++) {
            value = o[start - 1];
            if (0 == start % i || 8 === i && 4 == start % i) {
              /** @type {number} */
              value = path[value >>> 24] << 24 ^ path[value >> 16 & 255] << 16 ^ path[value >> 8 & 255] << 8 ^ path[255 & value];
              if (0 == start % i) {
                /** @type {number} */
                value = value << 8 ^ value >>> 24 ^ c << 24;
                /** @type {number} */
                c = c << 1 ^ 283 * (c >> 7);
              }
            }
            /** @type {number} */
            o[start] = o[start - i] ^ value;
          }
          /** @type {number} */
          i = 0;
          for (;start;i++, start--) {
            value = o[3 & i ? start : start - 4];
            obj[i] = 4 >= start || 4 > i ? value : seen[0][path[value >>> 24]] ^ seen[1][path[value >> 16 & 255]] ^ seen[2][path[value >> 8 & 255]] ^ seen[3][path[255 & value]];
          }
        };
        options.cipher.aes.prototype = {
          /**
           * @param {?} opt_attributes
           * @return {?}
           */
          encrypt : function(opt_attributes) {
            return has(this, opt_attributes, 0);
          },
          /**
           * @param {number} val
           * @return {?}
           */
          decrypt : function(val) {
            return has(this, val, 1);
          },
          A : [[[], [], [], [], []], [[], [], [], [], []]],
          /**
           * @return {undefined}
           */
          J : function() {
            var type;
            var key;
            var xi;
            var val;
            var record;
            var i;
            var doc;
            var data = this.A[0];
            var node = this.A[1];
            var index = data[4];
            var keys = node[4];
            /** @type {Array} */
            var _ref = [];
            /** @type {Array} */
            var key_types = [];
            /** @type {number} */
            type = 0;
            for (;256 > type;type++) {
              /** @type {number} */
              key_types[(_ref[type] = type << 1 ^ 283 * (type >> 7)) ^ type] = type;
            }
            /** @type {number} */
            key = xi = 0;
            for (;!index[key];key ^= val || 1, xi = key_types[xi] || 1) {
              /** @type {number} */
              i = (i = xi ^ xi << 1 ^ xi << 2 ^ xi << 3 ^ xi << 4) >> 8 ^ 255 & i ^ 99;
              /** @type {number} */
              index[key] = i;
              /** @type {number} */
              keys[i] = key;
              /** @type {number} */
              doc = 16843009 * (record = _ref[type = _ref[val = _ref[key]]]) ^ 65537 * type ^ 257 * val ^ 16843008 * key;
              /** @type {number} */
              record = 257 * _ref[i] ^ 16843008 * i;
              /** @type {number} */
              type = 0;
              for (;4 > type;type++) {
                /** @type {number} */
                data[type][key] = record = record << 24 ^ record >>> 8;
                /** @type {number} */
                node[type][i] = doc = doc << 24 ^ doc >>> 8;
              }
            }
            /** @type {number} */
            type = 0;
            for (;5 > type;type++) {
              data[type] = data[type].slice(0);
              node[type] = node[type].slice(0);
            }
          }
        };
        options.bitArray = {
          /**
           * @param {Object} data
           * @param {number} mayParseLabeledStatementInstead
           * @param {number} value
           * @return {?}
           */
          bitSlice : function(data, mayParseLabeledStatementInstead, value) {
            return data = options.bitArray.ea(data.slice(mayParseLabeledStatementInstead / 32), 32 - (31 & mayParseLabeledStatementInstead)).slice(1), value === target ? data : options.bitArray.clamp(data, value - mayParseLabeledStatementInstead);
          },
          /**
           * @param {(Function|string)} arr
           * @param {number} start
           * @param {number} index
           * @return {?}
           */
          extract : function(arr, start, index) {
            /** @type {number} */
            var bits = Math.floor(-start - index & 31);
            return(-32 & (start + index - 1 ^ start) ? arr[start / 32 | 0] << 32 - bits ^ arr[start / 32 + 1 | 0] >>> bits : arr[start / 32 | 0] >>> bits) & (1 << index) - 1;
          },
          /**
           * @param {?} obj
           * @param {?} data
           * @return {?}
           */
          concat : function(obj, data) {
            if (0 === obj.length || 0 === data.length) {
              return obj.concat(data);
            }
            var udataCur = obj[obj.length - 1];
            var string = options.bitArray.getPartial(udataCur);
            return 32 === string ? obj.concat(data) : options.bitArray.ea(data, string, 0 | udataCur, obj.slice(0, obj.length - 1));
          },
          /**
           * @param {?} data
           * @return {?}
           */
          bitLength : function(data) {
            var linesLen = data.length;
            return 0 === linesLen ? 0 : 32 * (linesLen - 1) + options.bitArray.getPartial(data[linesLen - 1]);
          },
          /**
           * @param {Object} data
           * @param {number} num
           * @return {?}
           */
          clamp : function(data, num) {
            if (32 * data.length < num) {
              return data;
            }
            var linesLen = (data = data.slice(0, Math.ceil(num / 32))).length;
            return num &= 31, 0 < linesLen && (num && (data[linesLen - 1] = options.bitArray.partial(num, data[linesLen - 1] & 2147483648 >> num - 1, 1))), data;
          },
          /**
           * @param {number} expectedNumberOfNonCommentArgs
           * @param {number} dataAndEvents
           * @param {number} deepDataAndEvents
           * @return {?}
           */
          partial : function(expectedNumberOfNonCommentArgs, dataAndEvents, deepDataAndEvents) {
            return 32 === expectedNumberOfNonCommentArgs ? dataAndEvents : (deepDataAndEvents ? 0 | dataAndEvents : dataAndEvents << 32 - expectedNumberOfNonCommentArgs) + 1099511627776 * expectedNumberOfNonCommentArgs;
          },
          /**
           * @param {number} value
           * @return {?}
           */
          getPartial : function(value) {
            return Math.round(value / 1099511627776) || 32;
          },
          /**
           * @param {Array} a
           * @param {Object} b
           * @return {?}
           */
          equal : function(a, b) {
            if (options.bitArray.bitLength(a) !== options.bitArray.bitLength(b)) {
              return FALSE;
            }
            var i;
            /** @type {number} */
            var n = 0;
            /** @type {number} */
            i = 0;
            for (;i < a.length;i++) {
              n |= a[i] ^ b[i];
            }
            return 0 === n;
          },
          /**
           * @param {string} data
           * @param {number} d
           * @param {number} a
           * @param {Array} stack
           * @return {?}
           */
          ea : function(data, d, a, stack) {
            var i;
            /** @type {number} */
            i = 0;
            if (stack === target) {
              /** @type {Array} */
              stack = [];
            }
            for (;32 <= d;d -= 32) {
              stack.push(a);
              /** @type {number} */
              a = 0;
            }
            if (0 === d) {
              return stack.concat(data);
            }
            /** @type {number} */
            i = 0;
            for (;i < data.length;i++) {
              stack.push(a | data[i] >>> d);
              /** @type {number} */
              a = data[i] << 32 - d;
            }
            return i = data.length ? data[data.length - 1] : 0, data = options.bitArray.getPartial(i), stack.push(options.bitArray.partial(d + data & 31, 32 < d + data ? a : stack.pop(), 1)), stack;
          },
          /**
           * @param {?} array
           * @param {Array} keepData
           * @return {?}
           */
          o : function(array, keepData) {
            return[array[0] ^ keepData[0], array[1] ^ keepData[1], array[2] ^ keepData[2], array[3] ^ keepData[3]];
          },
          /**
           * @param {Array} codeSegments
           * @return {?}
           */
          byteswapM : function(codeSegments) {
            var i;
            var part;
            /** @type {number} */
            i = 0;
            for (;i < codeSegments.length;++i) {
              part = codeSegments[i];
              /** @type {number} */
              codeSegments[i] = part >>> 24 | part >>> 8 & 65280 | (65280 & part) << 8 | part << 24;
            }
            return codeSegments;
          }
        };
        options.codec.utf8String = {
          /**
           * @param {Object} data
           * @return {?}
           */
          fromBits : function(data) {
            var t;
            var tmp;
            /** @type {string} */
            var s = "";
            var d = options.bitArray.bitLength(data);
            /** @type {number} */
            t = 0;
            for (;t < d / 8;t++) {
              if (0 == (3 & t)) {
                tmp = data[t / 4];
              }
              s += String.fromCharCode(tmp >>> 24);
              tmp <<= 8;
            }
            return decodeURIComponent(escape(s));
          },
          /**
           * @param {string} str
           * @return {?}
           */
          toBits : function(str) {
            /** @type {string} */
            str = unescape(encodeURIComponent(str));
            var i;
            /** @type {Array} */
            var buffer = [];
            /** @type {number} */
            var node = 0;
            /** @type {number} */
            i = 0;
            for (;i < str.length;i++) {
              /** @type {number} */
              node = node << 8 | str.charCodeAt(i);
              if (3 == (3 & i)) {
                buffer.push(node);
                /** @type {number} */
                node = 0;
              }
            }
            return 3 & i && buffer.push(options.bitArray.partial(8 * (3 & i), node)), buffer;
          }
        };
        options.codec.hex = {
          /**
           * @param {Object} data
           * @return {?}
           */
          fromBits : function(data) {
            var byteIndex;
            /** @type {string} */
            var headBuffer = "";
            /** @type {number} */
            byteIndex = 0;
            for (;byteIndex < data.length;byteIndex++) {
              headBuffer += (0xf00000000000 + (0 | data[byteIndex])).toString(16).substr(4);
            }
            return headBuffer.substr(0, options.bitArray.bitLength(data) / 4);
          },
          /**
           * @param {string} str
           * @return {?}
           */
          toBits : function(str) {
            var s;
            var valsLength;
            /** @type {Array} */
            var memory = [];
            valsLength = (str = str.replace(/\s|0x/g, "")).length;
            str += "00000000";
            /** @type {number} */
            s = 0;
            for (;s < str.length;s += 8) {
              memory.push(0 ^ parseInt(str.substr(s, 8), 16));
            }
            return options.bitArray.clamp(memory, 4 * valsLength);
          }
        };
        options.codec.base32 = {
          D : "0123456789abcdefghjkmnpqrstvwxyz",
          BITS : 32,
          BASE : 5,
          REMAINING : 27,
          /**
           * @param {Object} data
           * @return {?}
           */
          fromBits : function(data) {
            var rootProperty;
            /** @type {number} */
            var sz = options.codec.base32.BASE;
            /** @type {number} */
            var REMAINING = options.codec.base32.REMAINING;
            /** @type {string} */
            var out = "";
            /** @type {number} */
            var at = 0;
            /** @type {string} */
            var tail = options.codec.base32.D;
            /** @type {number} */
            var a = 0;
            var newState = options.bitArray.bitLength(data);
            /** @type {number} */
            rootProperty = 0;
            for (;out.length * sz <= newState;) {
              out += tail.charAt((a ^ data[rootProperty] >>> at) >>> REMAINING);
              if (at < sz) {
                /** @type {number} */
                a = data[rootProperty] << sz - at;
                at += REMAINING;
                rootProperty++;
              } else {
                a <<= sz;
                at -= sz;
              }
            }
            return out;
          },
          /**
           * @param {?} str
           * @return {?}
           */
          toBits : function(str) {
            var i;
            var x;
            /** @type {number} */
            var sz = options.codec.base32.BITS;
            /** @type {number} */
            var el = options.codec.base32.BASE;
            /** @type {number} */
            var mod = options.codec.base32.REMAINING;
            /** @type {Array} */
            var out = [];
            /** @type {number} */
            var bits = 0;
            /** @type {string} */
            var whitespace = options.codec.base32.D;
            /** @type {number} */
            var ta = 0;
            /** @type {number} */
            i = 0;
            for (;i < str.length;i++) {
              if (0 > (x = whitespace.indexOf(str.charAt(i)))) {
                callback(new options.exception.invalid("this isn't base32!"));
              }
              if (bits > mod) {
                bits -= mod;
                out.push(ta ^ x >>> bits);
                /** @type {number} */
                ta = x << sz - bits;
              } else {
                ta ^= x << sz - (bits += el);
              }
            }
            return 56 & bits && out.push(options.bitArray.partial(56 & bits, ta, 1)), out;
          }
        };
        options.codec.base64 = {
          D : "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/",
          /**
           * @param {Object} data
           * @param {number} recurring
           * @param {number} i
           * @return {?}
           */
          fromBits : function(data, recurring, i) {
            /** @type {string} */
            var out = "";
            /** @type {number} */
            var bits = 0;
            /** @type {string} */
            var token = options.codec.base64.D;
            /** @type {number} */
            var ta = 0;
            var newState = options.bitArray.bitLength(data);
            if (i) {
              /** @type {string} */
              token = token.substr(0, 62) + "-_";
            }
            /** @type {number} */
            i = 0;
            for (;6 * out.length < newState;) {
              out += token.charAt((ta ^ data[i] >>> bits) >>> 26);
              if (6 > bits) {
                /** @type {number} */
                ta = data[i] << 6 - bits;
                bits += 26;
                i++;
              } else {
                ta <<= 6;
                bits -= 6;
              }
            }
            for (;3 & out.length && !recurring;) {
              out += "=";
            }
            return out;
          },
          /**
           * @param {string} str
           * @param {number} dataAndEvents
           * @return {?}
           */
          toBits : function(str, dataAndEvents) {
            str = str.replace(/\s|=/g, "");
            var i;
            var x;
            /** @type {Array} */
            var out = [];
            /** @type {number} */
            var bits = 0;
            /** @type {string} */
            var data = options.codec.base64.D;
            /** @type {number} */
            var ta = 0;
            if (dataAndEvents) {
              /** @type {string} */
              data = data.substr(0, 62) + "-_";
            }
            /** @type {number} */
            i = 0;
            for (;i < str.length;i++) {
              if (0 > (x = data.indexOf(str.charAt(i)))) {
                callback(new options.exception.invalid("this isn't base64!"));
              }
              if (26 < bits) {
                bits -= 26;
                out.push(ta ^ x >>> bits);
                /** @type {number} */
                ta = x << 32 - bits;
              } else {
                ta ^= x << 32 - (bits += 6);
              }
            }
            return 56 & bits && out.push(options.bitArray.partial(56 & bits, ta, 1)), out;
          }
        };
        options.codec.base64url = {
          /**
           * @param {Object} data
           * @return {?}
           */
          fromBits : function(data) {
            return options.codec.base64.fromBits(data, 1, 1);
          },
          /**
           * @param {string} str
           * @return {?}
           */
          toBits : function(str) {
            return options.codec.base64.toBits(str, 1);
          }
        };
        options.codec.bytes = {
          /**
           * @param {Object} data
           * @return {?}
           */
          fromBits : function(data) {
            var t;
            var tmp;
            /** @type {Array} */
            var out = [];
            var d = options.bitArray.bitLength(data);
            /** @type {number} */
            t = 0;
            for (;t < d / 8;t++) {
              if (0 == (3 & t)) {
                tmp = data[t / 4];
              }
              out.push(tmp >>> 24);
              tmp <<= 8;
            }
            return out;
          },
          /**
           * @param {?} str
           * @return {?}
           */
          toBits : function(str) {
            var strCounter;
            /** @type {Array} */
            var buffer = [];
            /** @type {number} */
            var node = 0;
            /** @type {number} */
            strCounter = 0;
            for (;strCounter < str.length;strCounter++) {
              /** @type {number} */
              node = node << 8 | str[strCounter];
              if (3 == (3 & strCounter)) {
                buffer.push(node);
                /** @type {number} */
                node = 0;
              }
            }
            return 3 & strCounter && buffer.push(options.bitArray.partial(8 * (3 & strCounter), node)), buffer;
          }
        };
        /**
         * @param {Object} data
         * @return {undefined}
         */
        options.hash.sha256 = function(data) {
          if (!this.d[0]) {
            this.J();
          }
          if (data) {
            this.h = data.h.slice(0);
            this.e = data.e.slice(0);
            this.c = data.c;
          } else {
            this.reset();
          }
        };
        /**
         * @param {?} data
         * @return {?}
         */
        options.hash.sha256.hash = function(data) {
          return(new options.hash.sha256).update(data).finalize();
        };
        options.hash.sha256.prototype = {
          blockSize : 512,
          /**
           * @return {?}
           */
          reset : function() {
            return this.h = this.m.slice(0), this.e = [], this.c = 0, this;
          },
          /**
           * @param {?} data
           * @return {?}
           */
          update : function(data) {
            if ("string" == typeof data) {
              data = options.codec.utf8String.toBits(data);
            }
            var c;
            var e = this.e = options.bitArray.concat(this.e, data);
            c = this.c;
            data = this.c = c + options.bitArray.bitLength(data);
            /** @type {number} */
            c = 512 + c & -512;
            for (;c <= data;c += 512) {
              this.k(e.splice(0, 16));
            }
            return this;
          },
          /**
           * @return {?}
           */
          finalize : function() {
            var t;
            var queue = this.e;
            var h = this.h;
            t = (queue = options.bitArray.concat(queue, [options.bitArray.partial(1, 1)])).length + 2;
            for (;15 & t;t++) {
              queue.push(0);
            }
            queue.push(Math.floor(this.c / 4294967296));
            queue.push(0 | this.c);
            for (;queue.length;) {
              this.k(queue.splice(0, 16));
            }
            return this.reset(), h;
          },
          m : [],
          d : [],
          /**
           * @return {undefined}
           */
          J : function() {
            /**
             * @param {number} num
             * @return {?}
             */
            function isInt(num) {
              return 4294967296 * (num - Math.floor(num)) | 0;
            }
            var factor;
            /** @type {number} */
            var opt = 0;
            /** @type {number} */
            var prime = 2;
            t: for (;64 > opt;prime++) {
              /** @type {number} */
              factor = 2;
              for (;factor * factor <= prime;factor++) {
                if (0 == prime % factor) {
                  continue t;
                }
              }
              if (8 > opt) {
                this.m[opt] = isInt(Math.pow(prime, 0.5));
              }
              this.d[opt] = isInt(Math.pow(prime, 1 / 3));
              opt++;
            }
          },
          /**
           * @param {number} c
           * @return {undefined}
           */
          k : function(c) {
            var i;
            var val;
            var map = c.slice(0);
            var p = this.h;
            var d = this.d;
            var value = p[0];
            var x = p[1];
            var y = p[2];
            var Y = p[3];
            var fragment = p[4];
            var node = p[5];
            var t = p[6];
            var s = p[7];
            /** @type {number} */
            c = 0;
            for (;64 > c;c++) {
              if (16 > c) {
                i = map[c];
              } else {
                i = map[c + 1 & 15];
                val = map[c + 14 & 15];
                /** @type {number} */
                i = map[15 & c] = (i >>> 7 ^ i >>> 18 ^ i >>> 3 ^ i << 25 ^ i << 14) + (val >>> 17 ^ val >>> 19 ^ val >>> 10 ^ val << 15 ^ val << 13) + map[15 & c] + map[c + 9 & 15] | 0;
              }
              i = i + s + (fragment >>> 6 ^ fragment >>> 11 ^ fragment >>> 25 ^ fragment << 26 ^ fragment << 21 ^ fragment << 7) + (t ^ fragment & (node ^ t)) + d[c];
              s = t;
              t = node;
              node = fragment;
              /** @type {number} */
              fragment = Y + i | 0;
              Y = y;
              y = x;
              /** @type {number} */
              value = i + ((x = value) & y ^ Y & (x ^ y)) + (x >>> 2 ^ x >>> 13 ^ x >>> 22 ^ x << 30 ^ x << 19 ^ x << 10) | 0;
            }
            /** @type {number} */
            p[0] = p[0] + value | 0;
            /** @type {number} */
            p[1] = p[1] + x | 0;
            /** @type {number} */
            p[2] = p[2] + y | 0;
            /** @type {number} */
            p[3] = p[3] + Y | 0;
            /** @type {number} */
            p[4] = p[4] + fragment | 0;
            /** @type {number} */
            p[5] = p[5] + node | 0;
            /** @type {number} */
            p[6] = p[6] + t | 0;
            /** @type {number} */
            p[7] = p[7] + s | 0;
          }
        };
        /**
         * @param {Object} a
         * @return {undefined}
         */
        options.hash.sha512 = function(a) {
          if (!this.d[0]) {
            this.J();
          }
          if (a) {
            this.h = a.h.slice(0);
            this.e = a.e.slice(0);
            this.c = a.c;
          } else {
            this.reset();
          }
        };
        /**
         * @param {?} data
         * @return {?}
         */
        options.hash.sha512.hash = function(data) {
          return(new options.hash.sha512).update(data).finalize();
        };
        options.hash.sha512.prototype = {
          blockSize : 1024,
          /**
           * @return {?}
           */
          reset : function() {
            return this.h = this.m.slice(0), this.e = [], this.c = 0, this;
          },
          /**
           * @param {?} data
           * @return {?}
           */
          update : function(data) {
            if ("string" == typeof data) {
              data = options.codec.utf8String.toBits(data);
            }
            var c;
            var e = this.e = options.bitArray.concat(this.e, data);
            c = this.c;
            data = this.c = c + options.bitArray.bitLength(data);
            /** @type {number} */
            c = 1024 + c & -1024;
            for (;c <= data;c += 1024) {
              this.k(e.splice(0, 32));
            }
            return this;
          },
          /**
           * @return {?}
           */
          finalize : function() {
            var t;
            var queue = this.e;
            var h = this.h;
            t = (queue = options.bitArray.concat(queue, [options.bitArray.partial(1, 1)])).length + 4;
            for (;31 & t;t++) {
              queue.push(0);
            }
            queue.push(0);
            queue.push(0);
            queue.push(Math.floor(this.c / 4294967296));
            queue.push(0 | this.c);
            for (;queue.length;) {
              this.k(queue.splice(0, 32));
            }
            return this.reset(), h;
          },
          m : [],
          ra : [12372232, 13281083, 9762859, 1914609, 15106769, 4090911, 4308331, 8266105],
          d : [],
          ta : [2666018, 15689165, 5061423, 9034684, 4764984, 380953, 1658779, 7176472, 197186, 7368638, 14987916, 16757986, 8096111, 1480369, 13046325, 6891156, 15813330, 5187043, 9229749, 11312229, 2818677, 10937475, 4324308, 1135541, 6741931, 11809296, 16458047, 15666916, 11046850, 698149, 229999, 945776, 13774844, 2541862, 12856045, 9810911, 11494366, 7844520, 15576806, 8533307, 15795044, 4337665, 16291729, 5553712, 15684120, 6662416, 7413802, 12308920, 13816008, 4303699, 9366425, 10176680, 13195875, 
          4295371, 6546291, 11712675, 15708924, 1519456, 15772530, 6568428, 6495784, 8568297, 13007125, 7492395, 2515356, 12632583, 14740254, 7262584, 1535930, 13146278, 16321966, 1853211, 294276, 13051027, 13221564, 1051980, 4080310, 6651434, 14088940, 4675607],
          /**
           * @return {undefined}
           */
          J : function() {
            /**
             * @param {number} num
             * @return {?}
             */
            function isInt(num) {
              return 4294967296 * (num - Math.floor(num)) | 0;
            }
            /**
             * @param {number} num
             * @return {?}
             */
            function frac(num) {
              return 1099511627776 * (num - Math.floor(num)) & 255;
            }
            var factor;
            /** @type {number} */
            var unlock = 0;
            /** @type {number} */
            var prime = 2;
            t: for (;80 > unlock;prime++) {
              /** @type {number} */
              factor = 2;
              for (;factor * factor <= prime;factor++) {
                if (0 == prime % factor) {
                  continue t;
                }
              }
              if (8 > unlock) {
                this.m[2 * unlock] = isInt(Math.pow(prime, 0.5));
                /** @type {number} */
                this.m[2 * unlock + 1] = frac(Math.pow(prime, 0.5)) << 24 | this.ra[unlock];
              }
              this.d[2 * unlock] = isInt(Math.pow(prime, 1 / 3));
              /** @type {number} */
              this.d[2 * unlock + 1] = frac(Math.pow(prime, 1 / 3)) << 24 | this.ta[unlock];
              unlock++;
            }
          },
          /**
           * @param {number} l
           * @return {undefined}
           */
          k : function(l) {
            var filename;
            var str;
            var args = l.slice(0);
            var h = this.h;
            var d = this.d;
            var h0 = h[0];
            var h0h = h[1];
            var r = h[2];
            var h1h = h[3];
            var value = h[4];
            var h2h = h[5];
            var startIndex = h[6];
            var h0l = h[7];
            var indent = h[8];
            var h1l = h[9];
            var remove = h[10];
            var h2l = h[11];
            var t = h[12];
            var h3l = h[13];
            var max = h[14];
            var h5l = h[15];
            var a = h0;
            var ah = h0h;
            var b = r;
            var bh = h1h;
            var c = value;
            var ch = h2h;
            var i = startIndex;
            var al = h0l;
            var line = indent;
            var bl = h1l;
            var type = remove;
            var cl = h2l;
            var data = t;
            var dl = h3l;
            var min = max;
            var fl = h5l;
            /** @type {number} */
            l = 0;
            for (;80 > l;l++) {
              if (16 > l) {
                filename = args[2 * l];
                str = args[2 * l + 1];
              } else {
                str = args[2 * (l - 15)];
                /** @type {number} */
                filename = ((html = args[2 * (l - 15) + 1]) << 31 | str >>> 1) ^ (html << 24 | str >>> 8) ^ str >>> 7;
                /** @type {number} */
                var prefix = (str << 31 | html >>> 1) ^ (str << 24 | html >>> 8) ^ (str << 25 | html >>> 7);
                str = args[2 * (l - 2)];
                /** @type {number} */
                var html = ((s = args[2 * (l - 2) + 1]) << 13 | str >>> 19) ^ (str << 3 | s >>> 29) ^ str >>> 6;
                /** @type {number} */
                var s = (str << 13 | s >>> 19) ^ (s << 3 | str >>> 29) ^ (str << 26 | s >>> 6);
                var pageY = args[2 * (l - 7)];
                var name = args[2 * (l - 16)];
                var x = args[2 * (l - 16) + 1];
                filename = filename + pageY + ((str = prefix + args[2 * (l - 7) + 1]) >>> 0 < prefix >>> 0 ? 1 : 0);
                filename += html + ((str += s) >>> 0 < s >>> 0 ? 1 : 0);
                filename += name + ((str += x) >>> 0 < x >>> 0 ? 1 : 0);
              }
              /** @type {number} */
              args[2 * l] = filename |= 0;
              /** @type {number} */
              args[2 * l + 1] = str |= 0;
              /** @type {number} */
              pageY = line & type ^ ~line & data;
              /** @type {number} */
              var group = bl & cl ^ ~bl & dl;
              /** @type {number} */
              var last = (s = a & b ^ a & c ^ b & c, ah & bh ^ ah & ch ^ bh & ch);
              var G = (name = (ah << 4 | a >>> 28) ^ (a << 30 | ah >>> 2) ^ (a << 25 | ah >>> 7), x = (a << 4 | ah >>> 28) ^ (ah << 30 | a >>> 2) ^ (ah << 25 | a >>> 7), d[2 * l]);
              var icon = d[2 * l + 1];
              prefix = min + ((bl << 18 | line >>> 14) ^ (bl << 14 | line >>> 18) ^ (line << 23 | bl >>> 9)) + ((html = fl + ((line << 18 | bl >>> 14) ^ (line << 14 | bl >>> 18) ^ (bl << 23 | line >>> 9))) >>> 0 < fl >>> 0 ? 1 : 0);
              min = data;
              fl = dl;
              data = type;
              dl = cl;
              type = line;
              cl = bl;
              /** @type {number} */
              line = i + (prefix = (prefix = (prefix += pageY + ((html += group) >>> 0 < group >>> 0 ? 1 : 0)) + (G + ((html += icon) >>> 0 < icon >>> 0 ? 1 : 0))) + (filename + ((html = html + str | 0) >>> 0 < str >>> 0 ? 1 : 0))) + ((bl = al + html | 0) >>> 0 < al >>> 0 ? 1 : 0) | 0;
              i = c;
              al = ch;
              c = b;
              ch = bh;
              b = a;
              bh = ah;
              /** @type {number} */
              a = prefix + (filename = name + s + ((str = x + last) >>> 0 < x >>> 0 ? 1 : 0)) + ((ah = html + str | 0) >>> 0 < html >>> 0 ? 1 : 0) | 0;
            }
            /** @type {number} */
            h0h = h[1] = h0h + ah | 0;
            /** @type {number} */
            h[0] = h0 + a + (h0h >>> 0 < ah >>> 0 ? 1 : 0) | 0;
            /** @type {number} */
            h1h = h[3] = h1h + bh | 0;
            /** @type {number} */
            h[2] = r + b + (h1h >>> 0 < bh >>> 0 ? 1 : 0) | 0;
            /** @type {number} */
            h2h = h[5] = h2h + ch | 0;
            /** @type {number} */
            h[4] = value + c + (h2h >>> 0 < ch >>> 0 ? 1 : 0) | 0;
            /** @type {number} */
            h0l = h[7] = h0l + al | 0;
            /** @type {number} */
            h[6] = startIndex + i + (h0l >>> 0 < al >>> 0 ? 1 : 0) | 0;
            /** @type {number} */
            h1l = h[9] = h1l + bl | 0;
            /** @type {number} */
            h[8] = indent + line + (h1l >>> 0 < bl >>> 0 ? 1 : 0) | 0;
            /** @type {number} */
            h2l = h[11] = h2l + cl | 0;
            /** @type {number} */
            h[10] = remove + type + (h2l >>> 0 < cl >>> 0 ? 1 : 0) | 0;
            /** @type {number} */
            h3l = h[13] = h3l + dl | 0;
            /** @type {number} */
            h[12] = t + data + (h3l >>> 0 < dl >>> 0 ? 1 : 0) | 0;
            /** @type {number} */
            h5l = h[15] = h5l + fl | 0;
            /** @type {number} */
            h[14] = max + min + (h5l >>> 0 < fl >>> 0 ? 1 : 0) | 0;
          }
        };
        /**
         * @param {Object} data
         * @return {undefined}
         */
        options.hash.sha1 = function(data) {
          if (data) {
            this.h = data.h.slice(0);
            this.e = data.e.slice(0);
            this.c = data.c;
          } else {
            this.reset();
          }
        };
        /**
         * @param {?} data
         * @return {?}
         */
        options.hash.sha1.hash = function(data) {
          return(new options.hash.sha1).update(data).finalize();
        };
        options.hash.sha1.prototype = {
          blockSize : 512,
          /**
           * @return {?}
           */
          reset : function() {
            return this.h = this.m.slice(0), this.e = [], this.c = 0, this;
          },
          /**
           * @param {?} data
           * @return {?}
           */
          update : function(data) {
            if ("string" == typeof data) {
              data = options.codec.utf8String.toBits(data);
            }
            var c;
            var e = this.e = options.bitArray.concat(this.e, data);
            c = this.c;
            data = this.c = c + options.bitArray.bitLength(data);
            /** @type {number} */
            c = this.blockSize + c & -this.blockSize;
            for (;c <= data;c += this.blockSize) {
              this.k(e.splice(0, 16));
            }
            return this;
          },
          /**
           * @return {?}
           */
          finalize : function() {
            var t;
            var queue = this.e;
            var h = this.h;
            t = (queue = options.bitArray.concat(queue, [options.bitArray.partial(1, 1)])).length + 2;
            for (;15 & t;t++) {
              queue.push(0);
            }
            queue.push(Math.floor(this.c / 4294967296));
            queue.push(0 | this.c);
            for (;queue.length;) {
              this.k(queue.splice(0, 16));
            }
            return this.reset(), h;
          },
          m : [1732584193, 4023233417, 2562383102, 271733878, 3285377520],
          d : [1518500249, 1859775393, 2400959708, 3395469782],
          /**
           * @param {number} n
           * @return {undefined}
           */
          k : function(n) {
            var TEMP;
            var A;
            var B;
            var C;
            var D;
            var E;
            var a = n.slice(0);
            var H = this.h;
            A = H[0];
            B = H[1];
            C = H[2];
            D = H[3];
            E = H[4];
            /** @type {number} */
            n = 0;
            for (;79 >= n;n++) {
              if (16 <= n) {
                /** @type {number} */
                a[n] = (a[n - 3] ^ a[n - 8] ^ a[n - 14] ^ a[n - 16]) << 1 | (a[n - 3] ^ a[n - 8] ^ a[n - 14] ^ a[n - 16]) >>> 31;
              }
              /** @type {number} */
              TEMP = (A << 5 | A >>> 27) + (TEMP = 19 >= n ? B & C | ~B & D : 39 >= n ? B ^ C ^ D : 59 >= n ? B & C | B & D | C & D : 79 >= n ? B ^ C ^ D : target) + E + a[n] + this.d[Math.floor(n / 20)] | 0;
              E = D;
              D = C;
              /** @type {number} */
              C = B << 30 | B >>> 2;
              B = A;
              /** @type {number} */
              A = TEMP;
            }
            /** @type {number} */
            H[0] = H[0] + A | 0;
            /** @type {number} */
            H[1] = H[1] + B | 0;
            /** @type {number} */
            H[2] = H[2] + C | 0;
            /** @type {number} */
            H[3] = H[3] + D | 0;
            /** @type {number} */
            H[4] = H[4] + E | 0;
          }
        };
        options.mode.ccm = {
          name : "ccm",
          /**
           * @param {?} opt_attributes
           * @param {string} str
           * @param {?} data
           * @param {Array} prop
           * @param {number} tlen
           * @return {?}
           */
          encrypt : function(opt_attributes, str, data, prop, tlen) {
            var y;
            var out = str.slice(0);
            var w = options.bitArray;
            /** @type {number} */
            var winPageY = w.bitLength(data) / 8;
            /** @type {number} */
            var h = w.bitLength(out) / 8;
            tlen = tlen || 64;
            prop = prop || [];
            if (7 > winPageY) {
              callback(new options.exception.invalid("ccm: iv must be at least 7 bytes"));
            }
            /** @type {number} */
            y = 2;
            for (;4 > y && h >>> 8 * y;y++) {
            }
            return y < 15 - winPageY && (y = 15 - winPageY), data = w.clamp(data, 8 * (15 - y)), str = options.mode.ccm.Y(opt_attributes, str, data, prop, tlen, y), out = options.mode.ccm.F(opt_attributes, out, data, str, tlen, y), w.concat(out.data, out.tag);
          },
          /**
           * @param {Object} s
           * @param {number} data
           * @param {?} x
           * @param {Array} m
           * @param {number} tlen
           * @return {?}
           */
          decrypt : function(s, data, x, m, tlen) {
            tlen = tlen || 64;
            m = m || [];
            var w = options.bitArray;
            /** @type {number} */
            var a = w.bitLength(x) / 8;
            var ol = w.bitLength(data);
            var out = w.clamp(data, ol - tlen);
            var ast = w.bitSlice(data, ol - tlen);
            /** @type {number} */
            ol = (ol - tlen) / 8;
            if (7 > a) {
              callback(new options.exception.invalid("ccm: iv must be at least 7 bytes"));
            }
            /** @type {number} */
            data = 2;
            for (;4 > data && ol >>> 8 * data;data++) {
            }
            return data < 15 - a && (data = 15 - a), x = w.clamp(x, 8 * (15 - data)), out = options.mode.ccm.F(s, out, x, ast, tlen, data), s = options.mode.ccm.Y(s, out.data, x, m, tlen, data), w.equal(out.tag, s) || callback(new options.exception.corrupt("ccm: tag doesn't match")), out.data;
          },
          /**
           * @param {Object} t
           * @param {string} d
           * @param {number} data
           * @param {number} k
           * @param {number} tlen
           * @param {Object} attributes
           * @return {?}
           */
          Y : function(t, d, data, k, tlen, attributes) {
            /** @type {Array} */
            var keys = [];
            var self = options.bitArray;
            /** @type {function (?, Array): ?} */
            var o = self.o;
            if (((tlen /= 8) % 2 || (4 > tlen || 16 < tlen)) && callback(new options.exception.invalid("ccm: invalid tag length")), (4294967295 < k.length || 4294967295 < d.length) && callback(new options.exception.bug("ccm: can't deal with 4GiB or more data")), attributes = [self.partial(8, (k.length ? 64 : 0) | tlen - 2 << 2 | attributes - 1)], (attributes = self.concat(attributes, data))[3] |= self.bitLength(d) / 8, attributes = t.encrypt(attributes), k.length) {
              if (65279 >= (data = self.bitLength(k) / 8)) {
                /** @type {Array} */
                keys = [self.partial(16, data)];
              } else {
                if (4294967295 >= data) {
                  keys = self.concat([self.partial(16, 65534)], [data]);
                }
              }
              keys = self.concat(keys, k);
              /** @type {number} */
              k = 0;
              for (;k < keys.length;k += 4) {
                attributes = t.encrypt(o(attributes, keys.slice(k, k + 4).concat([0, 0, 0])));
              }
            }
            /** @type {number} */
            k = 0;
            for (;k < d.length;k += 4) {
              attributes = t.encrypt(o(attributes, d.slice(k, k + 4).concat([0, 0, 0])));
            }
            return self.clamp(attributes, 8 * tlen);
          },
          /**
           * @param {Object} value
           * @param {Array} data
           * @param {?} attributes
           * @param {(Object|string)} input
           * @param {Array} tlen
           * @param {number} y
           * @return {?}
           */
          F : function(value, data, attributes, input, tlen, y) {
            var i;
            var w = options.bitArray;
            /** @type {function (?, Array): ?} */
            i = w.o;
            var iLen = data.length;
            var bl = w.bitLength(data);
            if (attributes = w.concat([w.partial(8, y - 1)], attributes).concat([0, 0, 0]).slice(0, 4), input = w.bitSlice(i(input, value.encrypt(attributes)), 0, tlen), !iLen) {
              return{
                tag : input,
                data : []
              };
            }
            /** @type {number} */
            i = 0;
            for (;i < iLen;i += 4) {
              attributes[3]++;
              tlen = value.encrypt(attributes);
              data[i] ^= tlen[0];
              data[i + 1] ^= tlen[1];
              data[i + 2] ^= tlen[2];
              data[i + 3] ^= tlen[3];
            }
            return{
              tag : input,
              data : w.clamp(data, bl)
            };
          }
        };
        if (options.beware === target) {
          options.beware = {};
        }
        /**
         * @return {undefined}
         */
        options.beware["CBC mode is dangerous because it doesn't protect message integrity."] = function() {
          options.mode.cbc = {
            name : "cbc",
            /**
             * @param {Object} opt_attributes
             * @param {string} str
             * @param {Array} m
             * @param {number} i
             * @return {?}
             */
            encrypt : function(opt_attributes, str, m, i) {
              if (i) {
                if (i.length) {
                  callback(new options.exception.invalid("cbc can't authenticate data"));
                }
              }
              if (128 !== options.bitArray.bitLength(m)) {
                callback(new options.exception.invalid("cbc iv must be 128 bits"));
              }
              var t = options.bitArray;
              /** @type {function (?, Array): ?} */
              var xor = t.o;
              var e = t.bitLength(str);
              /** @type {number} */
              var j = 0;
              /** @type {Array} */
              var ctx = [];
              if (7 & e) {
                callback(new options.exception.invalid("pkcs#5 padding only works for multiples of a byte"));
              }
              /** @type {number} */
              i = 0;
              for (;j + 128 <= e;i += 4, j += 128) {
                m = opt_attributes.encrypt(xor(m, str.slice(i, i + 4)));
                ctx.splice(i, 0, m[0], m[1], m[2], m[3]);
              }
              return e = 16843009 * (16 - (e >> 3 & 15)), m = opt_attributes.encrypt(xor(m, t.concat(str, [e, e, e, e]).slice(i, i + 4))), ctx.splice(i, 0, m[0], m[1], m[2], m[3]), ctx;
            },
            /**
             * @param {Object} val
             * @param {string} data
             * @param {number} arr
             * @param {number} i
             * @return {?}
             */
            decrypt : function(val, data, arr, i) {
              if (i) {
                if (i.length) {
                  callback(new options.exception.invalid("cbc can't authenticate data"));
                }
              }
              if (128 !== options.bitArray.bitLength(arr)) {
                callback(new options.exception.invalid("cbc iv must be 128 bits"));
              }
              if (127 & options.bitArray.bitLength(data) || !data.length) {
                callback(new options.exception.corrupt("cbc ciphertext must be a positive multiple of the block size"));
              }
              var m;
              var test = options.bitArray;
              /** @type {function (?, Array): ?} */
              var _map = test.o;
              /** @type {Array} */
              var x = [];
              /** @type {number} */
              i = 0;
              for (;i < data.length;i += 4) {
                m = data.slice(i, i + 4);
                arr = _map(arr, val.decrypt(m));
                x.splice(i, 0, arr[0], arr[1], arr[2], arr[3]);
                arr = m;
              }
              return(0 == (m = 255 & x[i - 1]) || 16 < m) && callback(new options.exception.corrupt("pkcs#5 padding corrupt")), arr = 16843009 * m, test.equal(test.bitSlice([arr, arr, arr, arr], 0, 8 * m), test.bitSlice(x, 32 * x.length - 8 * m, 32 * x.length)) || callback(new options.exception.corrupt("pkcs#5 padding corrupt")), test.bitSlice(x, 0, 32 * x.length - 8 * m);
            }
          };
        };
        options.mode.ocb2 = {
          name : "ocb2",
          /**
           * @param {Object} opt_attributes
           * @param {string} input
           * @param {?} attributes
           * @param {string} adata
           * @param {number} tlen
           * @param {boolean} premac
           * @return {?}
           */
          encrypt : function(opt_attributes, input, attributes, adata, tlen, premac) {
            if (128 !== options.bitArray.bitLength(attributes)) {
              callback(new options.exception.invalid("ocb iv must be 128 bits"));
            }
            var key;
            /** @type {function (Array): ?} */
            var times2 = options.mode.ocb2.V;
            var w = options.bitArray;
            /** @type {function (?, Array): ?} */
            var xor = w.o;
            /** @type {Array} */
            var checksum = [0, 0, 0, 0];
            attributes = times2(opt_attributes.encrypt(attributes));
            var name;
            /** @type {Array} */
            var classes = [];
            adata = adata || [];
            tlen = tlen || 64;
            /** @type {number} */
            key = 0;
            for (;key + 4 < input.length;key += 4) {
              checksum = xor(checksum, name = input.slice(key, key + 4));
              /** @type {Array} */
              classes = classes.concat(xor(attributes, opt_attributes.encrypt(xor(attributes, name))));
              attributes = times2(attributes);
            }
            return name = input.slice(key), input = w.bitLength(name), key = opt_attributes.encrypt(xor(attributes, [0, 0, 0, input])), name = w.clamp(xor(name.concat([0, 0, 0]), key), input), checksum = xor(checksum, xor(name.concat([0, 0, 0]), key)), checksum = opt_attributes.encrypt(xor(checksum, xor(attributes, times2(attributes)))), adata.length && (checksum = xor(checksum, premac ? adata : options.mode.ocb2.pmac(opt_attributes, adata))), classes.concat(w.concat(name, w.clamp(checksum, tlen)));
          },
          /**
           * @param {Object} prp
           * @param {Object} ciphertext
           * @param {number} attributes
           * @param {string} adata
           * @param {number} tlen
           * @param {boolean} premac
           * @return {?}
           */
          decrypt : function(prp, ciphertext, attributes, adata, tlen, premac) {
            if (128 !== options.bitArray.bitLength(attributes)) {
              callback(new options.exception.invalid("ocb iv must be 128 bits"));
            }
            tlen = tlen || 64;
            var type;
            var bl;
            /** @type {function (Array): ?} */
            var makeArray = options.mode.ocb2.V;
            var w = options.bitArray;
            /** @type {function (?, Array): ?} */
            var xor = w.o;
            /** @type {Array} */
            var data = [0, 0, 0, 0];
            var checkSet = makeArray(prp.encrypt(attributes));
            /** @type {number} */
            var len = options.bitArray.bitLength(ciphertext) - tlen;
            /** @type {Array} */
            var key = [];
            adata = adata || [];
            /** @type {number} */
            attributes = 0;
            for (;attributes + 4 < len / 32;attributes += 4) {
              type = xor(checkSet, prp.decrypt(xor(checkSet, ciphertext.slice(attributes, attributes + 4))));
              data = xor(data, type);
              /** @type {Array} */
              key = key.concat(type);
              checkSet = makeArray(checkSet);
            }
            return bl = len - 32 * attributes, type = prp.encrypt(xor(checkSet, [0, 0, 0, bl])), type = xor(type, w.clamp(ciphertext.slice(attributes), bl).concat([0, 0, 0])), data = xor(data, type), data = prp.encrypt(xor(data, xor(checkSet, makeArray(checkSet)))), adata.length && (data = xor(data, premac ? adata : options.mode.ocb2.pmac(prp, adata))), w.equal(w.clamp(data, tlen), w.bitSlice(ciphertext, len)) || callback(new options.exception.corrupt("ocb: tag doesn't match")), key.concat(w.clamp(type, 
            bl));
          },
          /**
           * @param {Object} prp
           * @param {string} adata
           * @return {?}
           */
          pmac : function(prp, adata) {
            var key;
            /** @type {function (Array): ?} */
            var parseFloat = options.mode.ocb2.V;
            var w = options.bitArray;
            /** @type {function (?, Array): ?} */
            var fn = w.o;
            /** @type {Array} */
            var type = [0, 0, 0, 0];
            var result = fn(result = prp.encrypt([0, 0, 0, 0]), parseFloat(parseFloat(result)));
            /** @type {number} */
            key = 0;
            for (;key + 4 < adata.length;key += 4) {
              result = parseFloat(result);
              type = fn(type, prp.encrypt(fn(result, adata.slice(key, key + 4))));
            }
            return key = adata.slice(key), 128 > w.bitLength(key) && (result = fn(result, parseFloat(result)), key = w.concat(key, [-2147483648, 0, 0, 0])), type = fn(type, key), prp.encrypt(fn(parseFloat(fn(result, parseFloat(result))), type));
          },
          /**
           * @param {Array} array
           * @return {?}
           */
          V : function(array) {
            return[array[0] << 1 ^ array[1] >>> 31, array[1] << 1 ^ array[2] >>> 31, array[2] << 1 ^ array[3] >>> 31, array[3] << 1 ^ 135 * (array[0] >>> 31)];
          }
        };
        options.mode.gcm = {
          name : "gcm",
          /**
           * @param {Object} opt_attributes
           * @param {Object} str
           * @param {Object} tlen
           * @param {Array} string
           * @param {number} seed
           * @return {?}
           */
          encrypt : function(opt_attributes, str, tlen, string, seed) {
            var udataCur = str.slice(0);
            return str = options.bitArray, string = string || [], opt_attributes = options.mode.gcm.F(pdataCur, opt_attributes, udataCur, string, tlen, seed || 128), str.concat(opt_attributes.data, opt_attributes.tag);
          },
          /**
           * @param {Object} d
           * @param {Object} str
           * @param {Object} opt_attributes
           * @param {Array} string
           * @param {number} tlen
           * @return {?}
           */
          decrypt : function(d, str, opt_attributes, string, tlen) {
            var data = str.slice(0);
            var w = options.bitArray;
            var ol = w.bitLength(data);
            return string = string || [], (tlen = tlen || 128) <= ol ? (str = w.bitSlice(data, ol - tlen), data = w.bitSlice(data, 0, ol - tlen)) : (str = data, data = []), d = options.mode.gcm.F(FALSE, d, data, string, opt_attributes, tlen), w.equal(d.tag, str) || callback(new options.exception.corrupt("gcm: tag doesn't match")), d.data;
          },
          /**
           * @param {Array} var_args
           * @param {Object} setting
           * @return {?}
           */
          pa : function(var_args, setting) {
            var sectionLength;
            var i;
            var checkSet;
            var types;
            var s;
            /** @type {function (?, Array): ?} */
            var makeArray = options.bitArray.o;
            /** @type {Array} */
            checkSet = [0, 0, 0, 0];
            types = setting.slice(0);
            /** @type {number} */
            sectionLength = 0;
            for (;128 > sectionLength;sectionLength++) {
              if (i = 0 != (var_args[Math.floor(sectionLength / 32)] & 1 << 31 - sectionLength % 32)) {
                checkSet = makeArray(checkSet, types);
              }
              /** @type {boolean} */
              s = 0 != (1 & types[3]);
              /** @type {number} */
              i = 3;
              for (;0 < i;i--) {
                /** @type {number} */
                types[i] = types[i] >>> 1 | (1 & types[i - 1]) << 31;
              }
              types[0] >>>= 1;
              if (s) {
                types[0] ^= -520093696;
              }
            }
            return checkSet;
          },
          /**
           * @param {Error} path
           * @param {Array} args
           * @param {Array} values
           * @return {?}
           */
          t : function(path, args, values) {
            var i;
            var valuesLen = values.length;
            args = args.slice(0);
            /** @type {number} */
            i = 0;
            for (;i < valuesLen;i += 4) {
              args[0] ^= 4294967295 & values[i];
              args[1] ^= 4294967295 & values[i + 1];
              args[2] ^= 4294967295 & values[i + 2];
              args[3] ^= 4294967295 & values[i + 3];
              args = options.mode.gcm.pa(args, path);
            }
            return args;
          },
          /**
           * @param {Object} value
           * @param {Object} a
           * @param {Array} data
           * @param {Array} args
           * @param {Object} attributes
           * @param {number} r
           * @return {?}
           */
          F : function(value, a, data, args, attributes, r) {
            var key;
            var tmp;
            var enc;
            var i;
            var opt_attributes;
            var l;
            var len;
            var d;
            var w = options.bitArray;
            l = data.length;
            len = w.bitLength(data);
            d = w.bitLength(args);
            tmp = w.bitLength(attributes);
            key = a.encrypt([0, 0, 0, 0]);
            if (96 === tmp) {
              attributes = attributes.slice(0);
              attributes = w.concat(attributes, [1]);
            } else {
              attributes = options.mode.gcm.t(key, [0, 0, 0, 0], attributes);
              attributes = options.mode.gcm.t(key, attributes, [0, 0, Math.floor(tmp / 4294967296), 4294967295 & tmp]);
            }
            tmp = options.mode.gcm.t(key, [0, 0, 0, 0], args);
            opt_attributes = attributes.slice(0);
            args = tmp.slice(0);
            if (!value) {
              args = options.mode.gcm.t(key, tmp, data);
            }
            /** @type {number} */
            i = 0;
            for (;i < l;i += 4) {
              opt_attributes[3]++;
              enc = a.encrypt(opt_attributes);
              data[i] ^= enc[0];
              data[i + 1] ^= enc[1];
              data[i + 2] ^= enc[2];
              data[i + 3] ^= enc[3];
            }
            return data = w.clamp(data, len), value && (args = options.mode.gcm.t(key, tmp, data)), value = [Math.floor(d / 4294967296), 4294967295 & d, Math.floor(len / 4294967296), 4294967295 & len], args = options.mode.gcm.t(key, args, value), enc = a.encrypt(attributes), args[0] ^= enc[0], args[1] ^= enc[1], args[2] ^= enc[2], args[3] ^= enc[3], {
              tag : w.bitSlice(args, 0, r),
              data : data
            };
          }
        };
        /**
         * @param {Object} data
         * @param {Function} Hash
         * @return {undefined}
         */
        options.misc.hmac = function(data, Hash) {
          this.$ = Hash = Hash || options.hash.sha256;
          var y;
          /** @type {Array} */
          var d = [[], []];
          /** @type {number} */
          var height = Hash.prototype.blockSize / 32;
          /** @type {Array} */
          this.C = [new Hash, new Hash];
          if (data.length > height) {
            data = Hash.hash(data);
          }
          /** @type {number} */
          y = 0;
          for (;y < height;y++) {
            /** @type {number} */
            d[0][y] = 909522486 ^ data[y];
            /** @type {number} */
            d[1][y] = 1549556828 ^ data[y];
          }
          this.C[0].update(d[0]);
          this.C[1].update(d[1]);
          this.U = new Hash(this.C[0]);
        };
        /** @type {function (?): ?} */
        options.misc.hmac.prototype.encrypt = options.misc.hmac.prototype.mac = function(msgs) {
          return this.ga && callback(new options.exception.invalid("encrypt on already updated hmac called!")), this.update(msgs), this.digest(msgs);
        };
        /**
         * @return {undefined}
         */
        options.misc.hmac.prototype.reset = function() {
          this.U = new this.$(this.C[0]);
          /** @type {boolean} */
          this.ga = FALSE;
        };
        /**
         * @param {?} data
         * @return {undefined}
         */
        options.misc.hmac.prototype.update = function(data) {
          /** @type {boolean} */
          this.ga = pdataCur;
          this.U.update(data);
        };
        /**
         * @return {?}
         */
        options.misc.hmac.prototype.digest = function() {
          var pdataCur = this.U.finalize();
          return pdataCur = (new this.$(this.C[1])).update(pdataCur).finalize(), this.reset(), pdataCur;
        };
        /**
         * @param {string} str
         * @param {string} qualifier
         * @param {number} high
         * @param {number} key
         * @param {Function} flags
         * @return {?}
         */
        options.misc.pbkdf2 = function(str, qualifier, high, key, flags) {
          high = high || 1E3;
          if (0 > key || 0 > high) {
            callback(options.exception.invalid("invalid params to pbkdf2"));
          }
          if ("string" == typeof str) {
            str = options.codec.utf8String.toBits(str);
          }
          if ("string" == typeof qualifier) {
            qualifier = options.codec.utf8String.toBits(qualifier);
          }
          str = new (flags = flags || options.misc.hmac)(str);
          var attributes;
          var low;
          var i;
          var tmp;
          /** @type {Array} */
          var data = [];
          var w = options.bitArray;
          /** @type {number} */
          tmp = 1;
          for (;32 * data.length < (key || 1);tmp++) {
            flags = attributes = str.encrypt(w.concat(qualifier, [tmp]));
            /** @type {number} */
            low = 1;
            for (;low < high;low++) {
              attributes = str.encrypt(attributes);
              /** @type {number} */
              i = 0;
              for (;i < attributes.length;i++) {
                flags[i] ^= attributes[i];
              }
            }
            /** @type {Array} */
            data = data.concat(flags);
          }
          return key && (data = w.clamp(data, key)), data;
        };
        /**
         * @param {Function} R
         * @return {undefined}
         */
        options.prng = function(R) {
          /** @type {Array} */
          this.j = [new options.hash.sha256];
          /** @type {Array} */
          this.u = [0];
          /** @type {number} */
          this.T = 0;
          this.L = {};
          /** @type {number} */
          this.S = 0;
          this.X = {};
          /** @type {number} */
          this.da = this.n = this.w = this.ma = 0;
          /** @type {Array} */
          this.d = [0, 0, 0, 0, 0, 0, 0, 0];
          /** @type {Array} */
          this.q = [0, 0, 0, 0];
          this.Q = target;
          /** @type {Function} */
          this.R = R;
          /** @type {boolean} */
          this.K = FALSE;
          this.P = {
            progress : {},
            seeded : {}
          };
          /** @type {number} */
          this.B = this.la = 0;
          /** @type {number} */
          this.M = 1;
          /** @type {number} */
          this.O = 2;
          /** @type {number} */
          this.ia = 65536;
          /** @type {Array} */
          this.W = [0, 48, 64, 96, 128, 192, 256, 384, 512, 768, 1024];
          /** @type {number} */
          this.ja = 3E4;
          /** @type {number} */
          this.ha = 80;
        };
        options.prng.prototype = {
          /**
           * @param {number} end
           * @param {number} recurring
           * @return {?}
           */
          randomWords : function(end, recurring) {
            var i;
            var args;
            /** @type {Array} */
            var func = [];
            if ((i = this.isReady(recurring)) === this.B && callback(new options.exception.notReady("generator isn't seeded")), i & this.O) {
              /** @type {boolean} */
              i = !(i & this.M);
              /** @type {Array} */
              args = [];
              var conditionIndex;
              /** @type {number} */
              var w = 0;
              this.da = args[0] = (new Date).valueOf() + this.ja;
              /** @type {number} */
              conditionIndex = 0;
              for (;16 > conditionIndex;conditionIndex++) {
                args.push(4294967296 * Math.random() | 0);
              }
              /** @type {number} */
              conditionIndex = 0;
              for (;conditionIndex < this.j.length && (args = args.concat(this.j[conditionIndex].finalize()), w += this.u[conditionIndex], this.u[conditionIndex] = 0, i || !(this.T & 1 << conditionIndex));conditionIndex++) {
              }
              if (this.T >= 1 << this.j.length) {
                this.j.push(new options.hash.sha256);
                this.u.push(0);
              }
              this.n -= w;
              if (w > this.w) {
                this.w = w;
              }
              this.T++;
              this.d = options.hash.sha256.hash(this.d.concat(args));
              this.Q = new options.cipher.aes(this.d);
              /** @type {number} */
              i = 0;
              for (;4 > i && (this.q[i] = this.q[i] + 1 | 0, !this.q[i]);i++) {
              }
            }
            /** @type {number} */
            i = 0;
            for (;i < end;i += 4) {
              if (0 == (i + 1) % this.ia) {
                go(this);
              }
              args = slice(this);
              func.push(args[0], args[1], args[2], args[3]);
            }
            return go(this), func.slice(0, end);
          },
          /**
           * @param {number} R
           * @param {string} allowZeroParanoia
           * @return {undefined}
           */
          setDefaultParanoia : function(R, allowZeroParanoia) {
            if (0 === R) {
              if ("Setting paranoia=0 will ruin your security; use it only for testing" !== allowZeroParanoia) {
                callback("Setting paranoia=0 will ruin your security; use it only for testing");
              }
            }
            /** @type {number} */
            this.R = R;
          },
          /**
           * @param {number} data
           * @param {number} expectedNumberOfNonCommentArgs
           * @param {?} j
           * @return {undefined}
           */
          addEntropy : function(data, expectedNumberOfNonCommentArgs, j) {
            j = j || "user";
            var id;
            var tmp;
            /** @type {number} */
            var t = (new Date).valueOf();
            var k = this.L[j];
            var oldReady = this.isReady();
            /** @type {number} */
            var shouldremove = 0;
            switch((id = this.X[j]) === target && (id = this.X[j] = this.ma++), k === target && (k = this.L[j] = 0), this.L[j] = (this.L[j] + 1) % this.j.length, typeof data) {
              case "number":
                if (expectedNumberOfNonCommentArgs === target) {
                  /** @type {number} */
                  expectedNumberOfNonCommentArgs = 1;
                }
                this.j[k].update([id, this.S++, 1, expectedNumberOfNonCommentArgs, t, 1, 0 | data]);
                break;
              case "object":
                if ("[object Uint32Array]" === (j = Object.prototype.toString.call(data))) {
                  /** @type {Array} */
                  tmp = [];
                  /** @type {number} */
                  j = 0;
                  for (;j < data.length;j++) {
                    tmp.push(data[j]);
                  }
                  /** @type {Array} */
                  data = tmp;
                } else {
                  if ("[object Array]" !== j) {
                    /** @type {number} */
                    shouldremove = 1;
                  }
                  /** @type {number} */
                  j = 0;
                  for (;j < data.length && !shouldremove;j++) {
                    if ("number" != typeof data[j]) {
                      /** @type {number} */
                      shouldremove = 1;
                    }
                  }
                }
                if (!shouldremove) {
                  if (expectedNumberOfNonCommentArgs === target) {
                    /** @type {number} */
                    j = expectedNumberOfNonCommentArgs = 0;
                    for (;j < data.length;j++) {
                      tmp = data[j];
                      for (;0 < tmp;) {
                        expectedNumberOfNonCommentArgs++;
                        tmp >>>= 1;
                      }
                    }
                  }
                  this.j[k].update([id, this.S++, 2, expectedNumberOfNonCommentArgs, t, data.length].concat(data));
                }
                break;
              case "string":
                if (expectedNumberOfNonCommentArgs === target) {
                  expectedNumberOfNonCommentArgs = data.length;
                }
                this.j[k].update([id, this.S++, 3, expectedNumberOfNonCommentArgs, t, data.length]);
                this.j[k].update(data);
                break;
              default:
                /** @type {number} */
                shouldremove = 1;
            }
            if (shouldremove) {
              callback(new options.exception.bug("random: addEntropy only supports number, array of numbers or string"));
            }
            this.u[k] += expectedNumberOfNonCommentArgs;
            this.n += expectedNumberOfNonCommentArgs;
            if (oldReady === this.B) {
              if (this.isReady() !== this.B) {
                declare("seeded", Math.max(this.w, this.n));
              }
              declare("progress", this.getProgress());
            }
          },
          /**
           * @param {Function} w
           * @return {?}
           */
          isReady : function(w) {
            return w = this.W[w !== target ? w : this.R], this.w && this.w >= w ? this.u[0] > this.ha && (new Date).valueOf() > this.da ? this.O | this.M : this.M : this.n >= w ? this.O | this.B : this.B;
          },
          /**
           * @param {Object} n
           * @return {?}
           */
          getProgress : function(n) {
            return n = this.W[n || this.R], this.w >= n || this.n > n ? 1 : this.n / n;
          },
          /**
           * @return {undefined}
           */
          startCollectors : function() {
            if (!this.K) {
              this.f = {
                loadTimeCollector : agree(this, this.ua),
                mouseCollector : agree(this, this.va),
                keyboardCollector : agree(this, this.sa),
                accelerometerCollector : agree(this, this.ka),
                touchCollector : agree(this, this.xa)
              };
              if (window.addEventListener) {
                window.addEventListener("load", this.f.loadTimeCollector, FALSE);
                window.addEventListener("mousemove", this.f.mouseCollector, FALSE);
                window.addEventListener("keypress", this.f.keyboardCollector, FALSE);
                window.addEventListener("devicemotion", this.f.accelerometerCollector, FALSE);
                window.addEventListener("touchmove", this.f.touchCollector, FALSE);
              } else {
                if (document.attachEvent) {
                  document.attachEvent("onload", this.f.loadTimeCollector);
                  document.attachEvent("onmousemove", this.f.mouseCollector);
                  document.attachEvent("keypress", this.f.keyboardCollector);
                } else {
                  callback(new options.exception.bug("can't attach event"));
                }
              }
              /** @type {boolean} */
              this.K = pdataCur;
            }
          },
          /**
           * @return {undefined}
           */
          stopCollectors : function() {
            if (this.K) {
              if (window.removeEventListener) {
                window.removeEventListener("load", this.f.loadTimeCollector, FALSE);
                window.removeEventListener("mousemove", this.f.mouseCollector, FALSE);
                window.removeEventListener("keypress", this.f.keyboardCollector, FALSE);
                window.removeEventListener("devicemotion", this.f.accelerometerCollector, FALSE);
                window.removeEventListener("touchmove", this.f.touchCollector, FALSE);
              } else {
                if (document.detachEvent) {
                  document.detachEvent("onload", this.f.loadTimeCollector);
                  document.detachEvent("onmousemove", this.f.mouseCollector);
                  document.detachEvent("keypress", this.f.keyboardCollector);
                }
              }
              /** @type {boolean} */
              this.K = FALSE;
            }
          },
          /**
           * @param {string} type
           * @param {Function} listener
           * @return {undefined}
           */
          addEventListener : function(type, listener) {
            /** @type {Function} */
            this.P[type][this.la++] = listener;
          },
          /**
           * @param {string} name
           * @param {?} capture
           * @return {undefined}
           */
          removeEventListener : function(name, capture) {
            var i;
            var cur;
            var prev = this.P[name];
            /** @type {Array} */
            var eventPath = [];
            for (cur in prev) {
              if (prev.hasOwnProperty(cur)) {
                if (prev[cur] === capture) {
                  eventPath.push(cur);
                }
              }
            }
            /** @type {number} */
            i = 0;
            for (;i < eventPath.length;i++) {
              delete prev[cur = eventPath[i]];
            }
          },
          /**
           * @return {undefined}
           */
          sa : function() {
            run(1);
          },
          /**
           * @param {Object} e
           * @return {undefined}
           */
          va : function(e) {
            var _i;
            var y;
            try {
              _i = e.x || (e.clientX || (e.offsetX || 0));
              y = e.y || (e.clientY || (e.offsetY || 0));
            } catch (t) {
              /** @type {number} */
              y = _i = 0;
            }
            if (0 != _i) {
              if (0 != y) {
                options.random.addEntropy([_i, y], 2, "mouse");
              }
            }
            run(0);
          },
          /**
           * @param {Event} event
           * @return {undefined}
           */
          xa : function(event) {
            event = event.touches[0] || event.changedTouches[0];
            options.random.addEntropy([event.pageX || event.clientX, event.pageY || event.clientY], 1, "touch");
            run(0);
          },
          /**
           * @return {undefined}
           */
          ua : function() {
            run(2);
          },
          /**
           * @param {number} info
           * @return {undefined}
           */
          ka : function(info) {
            if (info = info.accelerationIncludingGravity.x || (info.accelerationIncludingGravity.y || info.accelerationIncludingGravity.z), window.orientation) {
              /** @type {number} */
              var pdataCur = window.orientation;
              if ("number" == typeof pdataCur) {
                options.random.addEntropy(pdataCur, 1, "accelerometer");
              }
            }
            if (info) {
              options.random.addEntropy(info, 2, "accelerometer");
            }
            run(0);
          }
        };
        options.random = new options.prng(6);
        t: {
          try {
            var ab;
            var crypt;
            var buf;
            var declarationError;
            if (declarationError = void 0 !== module) {
              var e;
              if (e = module.exports) {
                var value;
                try {
                  value = $sanitize(119);
                } catch (t) {
                  /** @type {null} */
                  value = null;
                }
                e = (crypt = value) && crypt.randomBytes;
              }
              declarationError = e;
            }
            if (declarationError) {
              ab = crypt.randomBytes(128);
              /** @type {Uint32Array} */
              ab = new Uint32Array((new Uint8Array(ab)).buffer);
              options.random.addEntropy(ab, 1024, "crypto['randomBytes']");
            } else {
              if ("undefined" != typeof window && "undefined" != typeof Uint32Array) {
                if (buf = new Uint32Array(32), window.crypto && window.crypto.getRandomValues) {
                  window.crypto.getRandomValues(buf);
                } else {
                  if (!window.msCrypto || !window.msCrypto.getRandomValues) {
                    break t;
                  }
                  window.msCrypto.getRandomValues(buf);
                }
                options.random.addEntropy(buf, 1024, "crypto['getRandomValues']");
              }
            }
          } catch (fmt) {
            if ("undefined" != typeof window) {
              if (window.console) {
                console.log("There was an error collecting entropy from the browser:");
                console.log(fmt);
              }
            }
          }
        }
        options.json = {
          defaults : {
            v : 1,
            iter : 1E3,
            ks : 128,
            ts : 64,
            mode : "ccm",
            adata : "",
            cipher : "aes"
          },
          /**
           * @param {string} password
           * @param {string} ms
           * @param {Object} prop
           * @param {number} p
           * @return {?}
           */
          oa : function(password, ms, prop, p) {
            prop = prop || {};
            p = p || {};
            var attributes;
            var $ = options.json;
            var expectedNumberOfNonCommentArgs = $.p({
              iv : options.random.randomWords(4, 0)
            }, $.defaults);
            return $.p(expectedNumberOfNonCommentArgs, prop), prop = expectedNumberOfNonCommentArgs.adata, "string" == typeof expectedNumberOfNonCommentArgs.salt && (expectedNumberOfNonCommentArgs.salt = options.codec.base64.toBits(expectedNumberOfNonCommentArgs.salt)), "string" == typeof expectedNumberOfNonCommentArgs.iv && (expectedNumberOfNonCommentArgs.iv = options.codec.base64.toBits(expectedNumberOfNonCommentArgs.iv)), (!options.mode[expectedNumberOfNonCommentArgs.mode] || (!options.cipher[expectedNumberOfNonCommentArgs.cipher] || 
            ("string" == typeof password && 100 >= expectedNumberOfNonCommentArgs.iter || (64 !== expectedNumberOfNonCommentArgs.ts && (96 !== expectedNumberOfNonCommentArgs.ts && 128 !== expectedNumberOfNonCommentArgs.ts) || (128 !== expectedNumberOfNonCommentArgs.ks && (192 !== expectedNumberOfNonCommentArgs.ks && 256 !== expectedNumberOfNonCommentArgs.ks) || (2 > expectedNumberOfNonCommentArgs.iv.length || 4 < expectedNumberOfNonCommentArgs.iv.length)))))) && callback(new options.exception.invalid("json encrypt: invalid parameters")), 
            "string" == typeof password ? (password = (attributes = options.misc.cachedPbkdf2(password, expectedNumberOfNonCommentArgs)).key.slice(0, expectedNumberOfNonCommentArgs.ks / 32), expectedNumberOfNonCommentArgs.salt = attributes.salt) : options.ecc && (password instanceof options.ecc.elGamal.publicKey && (attributes = password.kem(), expectedNumberOfNonCommentArgs.kemtag = attributes.tag, password = attributes.key.slice(0, expectedNumberOfNonCommentArgs.ks / 32))), "string" == typeof ms && 
            (ms = options.codec.utf8String.toBits(ms)), "string" == typeof prop && (prop = options.codec.utf8String.toBits(prop)), attributes = new options.cipher[expectedNumberOfNonCommentArgs.cipher](password), $.p(p, expectedNumberOfNonCommentArgs), p.key = password, expectedNumberOfNonCommentArgs.ct = options.mode[expectedNumberOfNonCommentArgs.mode].encrypt(attributes, ms, expectedNumberOfNonCommentArgs.iv, prop, expectedNumberOfNonCommentArgs.ts), expectedNumberOfNonCommentArgs;
          },
          /**
           * @param {?} opt_attributes
           * @param {string} str
           * @param {?} seed
           * @param {Object} object
           * @return {?}
           */
          encrypt : function(opt_attributes, str, seed, object) {
            var self = options.json;
            var key = self.oa.apply(self, arguments);
            return self.encode(key);
          },
          /**
           * @param {string} password
           * @param {Object} prop
           * @param {Object} testName
           * @param {number} expectedNumberOfNonCommentArgs
           * @return {?}
           */
          na : function(password, prop, testName, expectedNumberOfNonCommentArgs) {
            testName = testName || {};
            expectedNumberOfNonCommentArgs = expectedNumberOfNonCommentArgs || {};
            var data;
            var p;
            var $ = options.json;
            return data = (prop = $.p($.p($.p({}, $.defaults), prop), testName, pdataCur)).adata, "string" == typeof prop.salt && (prop.salt = options.codec.base64.toBits(prop.salt)), "string" == typeof prop.iv && (prop.iv = options.codec.base64.toBits(prop.iv)), (!options.mode[prop.mode] || (!options.cipher[prop.cipher] || ("string" == typeof password && 100 >= prop.iter || (64 !== prop.ts && (96 !== prop.ts && 128 !== prop.ts) || (128 !== prop.ks && (192 !== prop.ks && 256 !== prop.ks) || (!prop.iv || 
            (2 > prop.iv.length || 4 < prop.iv.length))))))) && callback(new options.exception.invalid("json decrypt: invalid parameters")), "string" == typeof password ? (password = (p = options.misc.cachedPbkdf2(password, prop)).key.slice(0, prop.ks / 32), prop.salt = p.salt) : options.ecc && (password instanceof options.ecc.elGamal.secretKey && (password = password.unkem(options.codec.base64.toBits(prop.kemtag)).slice(0, prop.ks / 32))), "string" == typeof data && (data = options.codec.utf8String.toBits(data)), 
            p = new options.cipher[prop.cipher](password), data = options.mode[prop.mode].decrypt(p, prop.ct, prop.iv, data, prop.ts), $.p(expectedNumberOfNonCommentArgs, prop), expectedNumberOfNonCommentArgs.key = password, 1 === testName.raw ? data : options.codec.utf8String.fromBits(data);
          },
          /**
           * @param {string} val
           * @param {(number|string)} data
           * @param {Object} key
           * @param {number} input
           * @return {?}
           */
          decrypt : function(val, data, key, input) {
            var test = options.json;
            return test.na(val, test.decode(data), key, input);
          },
          /**
           * @param {Object} obj
           * @return {?}
           */
          encode : function(obj) {
            var i;
            /** @type {string} */
            var s = "{";
            /** @type {string} */
            var summary = "";
            for (i in obj) {
              if (obj.hasOwnProperty(i)) {
                switch(i.match(/^[a-z0-9]+$/i) || callback(new options.exception.invalid("json encode: invalid property name")), s += summary + '"' + i + '":', summary = ",", typeof obj[i]) {
                  case "number":
                  ;
                  case "boolean":
                    s += obj[i];
                    break;
                  case "string":
                    s += '"' + escape(obj[i]) + '"';
                    break;
                  case "object":
                    s += '"' + options.codec.base64.fromBits(obj[i], 0) + '"';
                    break;
                  default:
                    callback(new options.exception.bug("json encode: unsupported type"));
                }
              }
            }
            return s + "}";
          },
          /**
           * @param {string} qs
           * @return {?}
           */
          decode : function(qs) {
            if (!(qs = qs.replace(/\s/g, "")).match(/^\{.*\}$/)) {
              callback(new options.exception.invalid("json decode: this isn't json!"));
            }
            qs = qs.replace(/^\{|\}$/g, "").split(/,/);
            var i;
            var parts;
            var obj = {};
            /** @type {number} */
            i = 0;
            for (;i < qs.length;i++) {
              if (!(parts = qs[i].match(/^\s*(?:(["']?)([a-z][a-z0-9]*)\1)\s*:\s*(?:(-?\d+)|"([a-z0-9+\/%*_.@=\-]*)"|(true|false))$/i))) {
                callback(new options.exception.invalid("json decode: this isn't json!"));
              }
              if (parts[3]) {
                /** @type {number} */
                obj[parts[2]] = parseInt(parts[3], 10);
              } else {
                if (parts[4]) {
                  obj[parts[2]] = parts[2].match(/^(ct|salt|iv)$/) ? options.codec.base64.toBits(parts[4]) : unescape(parts[4]);
                } else {
                  if (parts[5]) {
                    /** @type {boolean} */
                    obj[parts[2]] = "true" === parts[5];
                  }
                }
              }
            }
            return obj;
          },
          /**
           * @param {number} expectedNumberOfNonCommentArgs
           * @param {Object} actual
           * @param {boolean} value
           * @return {?}
           */
          p : function(expectedNumberOfNonCommentArgs, actual, value) {
            if (expectedNumberOfNonCommentArgs === target && (expectedNumberOfNonCommentArgs = {}), actual === target) {
              return expectedNumberOfNonCommentArgs;
            }
            var key;
            for (key in actual) {
              if (actual.hasOwnProperty(key)) {
                if (value) {
                  if (expectedNumberOfNonCommentArgs[key] !== target) {
                    if (expectedNumberOfNonCommentArgs[key] !== actual[key]) {
                      callback(new options.exception.invalid("required parameter overridden"));
                    }
                  }
                }
                expectedNumberOfNonCommentArgs[key] = actual[key];
              }
            }
            return expectedNumberOfNonCommentArgs;
          },
          /**
           * @param {Object} obj
           * @param {Object} match
           * @return {?}
           */
          za : function(obj, match) {
            var k;
            var values = {};
            for (k in obj) {
              if (obj.hasOwnProperty(k)) {
                if (obj[k] !== match[k]) {
                  values[k] = obj[k];
                }
              }
            }
            return values;
          },
          /**
           * @param {?} obj
           * @param {Array} parts
           * @return {?}
           */
          ya : function(obj, parts) {
            var i;
            var o = {};
            /** @type {number} */
            i = 0;
            for (;i < parts.length;i++) {
              if (obj[parts[i]] !== target) {
                o[parts[i]] = obj[parts[i]];
              }
            }
            return o;
          }
        };
        /** @type {function (?, string, ?, Object): ?} */
        options.encrypt = options.json.encrypt;
        /** @type {function (string, (number|string), Object, number): ?} */
        options.decrypt = options.json.decrypt;
        options.misc.wa = {};
        /**
         * @param {string} password
         * @param {Object} obj
         * @return {?}
         */
        options.misc.cachedPbkdf2 = function(password, obj) {
          var c;
          var salt = options.misc.wa;
          return c = (obj = obj || {}).iter || 1E3, (c = (salt = salt[password] = salt[password] || {})[c] = salt[c] || {
            firstSalt : obj.salt && obj.salt.length ? obj.salt.slice(0) : options.random.randomWords(2, 0)
          })[salt = obj.salt === target ? c.firstSalt : obj.salt] = c[salt] || options.misc.pbkdf2(password, salt, obj.iter), {
            key : c[salt].slice(0),
            salt : salt.slice(0)
          };
        };
        /**
         * @param {string} opt_e
         * @return {undefined}
         */
        options.bn = function(opt_e) {
          this.initWith(opt_e);
        };
        options.bn.prototype = {
          radix : 24,
          maxMul : 8,
          /** @type {function (string): undefined} */
          i : options.bn,
          /**
           * @return {?}
           */
          copy : function() {
            return new this.i(this);
          },
          /**
           * @param {string} val
           * @return {?}
           */
          initWith : function(val) {
            var step;
            /** @type {number} */
            var j = 0;
            switch(typeof val) {
              case "object":
                this.limbs = val.limbs.slice(0);
                break;
              case "number":
                /** @type {Array} */
                this.limbs = [val];
                this.normalize();
                break;
              case "string":
                /** @type {string} */
                val = val.replace(/^0x/, "");
                /** @type {Array} */
                this.limbs = [];
                /** @type {number} */
                step = this.radix / 4;
                /** @type {number} */
                j = 0;
                for (;j < val.length;j += step) {
                  this.limbs.push(parseInt(val.substring(Math.max(val.length - j - step, 0), val.length - j), 16));
                }
                break;
              default:
                /** @type {Array} */
                this.limbs = [0];
            }
            return this;
          },
          /**
           * @param {number} recurring
           * @return {?}
           */
          equals : function(recurring) {
            if ("number" == typeof recurring) {
              recurring = new this.i(recurring);
            }
            var i;
            /** @type {number} */
            var r = 0;
            this.fullReduce();
            recurring.fullReduce();
            /** @type {number} */
            i = 0;
            for (;i < this.limbs.length || i < recurring.limbs.length;i++) {
              r |= this.getLimb(i) ^ recurring.getLimb(i);
            }
            return 0 === r;
          },
          /**
           * @param {number} num
           * @return {?}
           */
          getLimb : function(num) {
            return num >= this.limbs.length ? 0 : this.limbs[num];
          },
          /**
           * @param {?} a
           * @return {?}
           */
          greaterEquals : function(a) {
            if ("number" == typeof a) {
              a = new this.i(a);
            }
            var cDigit;
            var nDigit;
            var n;
            /** @type {number} */
            var d = 0;
            /** @type {number} */
            var b = 0;
            /** @type {number} */
            cDigit = Math.max(this.limbs.length, a.limbs.length) - 1;
            for (;0 <= cDigit;cDigit--) {
              d |= (nDigit = this.getLimb(cDigit)) - (n = a.getLimb(cDigit)) & ~(b |= n - nDigit & ~d);
            }
            return(b | ~d) >>> 31;
          },
          /**
           * @return {?}
           */
          toString : function() {
            this.fullReduce();
            var i;
            var right;
            /** @type {string} */
            var left = "";
            var results = this.limbs;
            /** @type {number} */
            i = 0;
            for (;i < this.limbs.length;i++) {
              right = results[i].toString(16);
              for (;i < this.limbs.length - 1 && 6 > right.length;) {
                /** @type {string} */
                right = "0" + right;
              }
              /** @type {string} */
              left = right + left;
            }
            return "0x" + left;
          },
          /**
           * @param {?} i
           * @return {?}
           */
          addM : function(i) {
            if ("object" != typeof i) {
              i = new this.i(i);
            }
            var offsets = this.limbs;
            var codeSegments = i.limbs;
            i = offsets.length;
            for (;i < codeSegments.length;i++) {
              /** @type {number} */
              offsets[i] = 0;
            }
            /** @type {number} */
            i = 0;
            for (;i < codeSegments.length;i++) {
              offsets[i] += codeSegments[i];
            }
            return this;
          },
          /**
           * @return {?}
           */
          doubleM : function() {
            var i;
            var name;
            /** @type {number} */
            var x = 0;
            var FRACBITS = this.radix;
            var radixMask = this.radixMask;
            var attrNames = this.limbs;
            /** @type {number} */
            i = 0;
            for (;i < attrNames.length;i++) {
              name = (name = attrNames[i]) + name + x;
              /** @type {number} */
              attrNames[i] = name & radixMask;
              /** @type {number} */
              x = name >> FRACBITS;
            }
            return x && attrNames.push(x), this;
          },
          /**
           * @return {?}
           */
          halveM : function() {
            var i;
            var inner;
            /** @type {number} */
            var arr = 0;
            var radix = this.radix;
            var parts = this.limbs;
            /** @type {number} */
            i = parts.length - 1;
            for (;0 <= i;i--) {
              inner = parts[i];
              /** @type {number} */
              parts[i] = inner + arr >> 1;
              /** @type {number} */
              arr = (1 & inner) << radix;
            }
            return parts[parts.length - 1] || parts.pop(), this;
          },
          /**
           * @param {number} num
           * @return {?}
           */
          subM : function(num) {
            if ("object" != typeof num) {
              num = new this.i(num);
            }
            var options = this.limbs;
            var digits = num.limbs;
            num = options.length;
            for (;num < digits.length;num++) {
              /** @type {number} */
              options[num] = 0;
            }
            /** @type {number} */
            num = 0;
            for (;num < digits.length;num++) {
              options[num] -= digits[num];
            }
            return this;
          },
          /**
           * @param {?} base
           * @return {?}
           */
          mod : function(base) {
            /** @type {boolean} */
            var e = !this.greaterEquals(new options.bn(0));
            base = (new options.bn(base)).normalize();
            var url = (new options.bn(this)).normalize();
            /** @type {number} */
            var n = 0;
            if (e) {
              url = (new options.bn(0)).subM(url).normalize();
            }
            for (;url.greaterEquals(base);n++) {
              base.doubleM();
            }
            if (e) {
              url = base.sub(url).normalize();
            }
            for (;0 < n;n--) {
              base.halveM();
              if (url.greaterEquals(base)) {
                url.subM(base).normalize();
              }
            }
            return url.trim();
          },
          /**
           * @param {?} target
           * @return {?}
           */
          inverseMod : function(target) {
            var j;
            var obj = new options.bn(1);
            var type = new options.bn(0);
            var t = new options.bn(this);
            var d = new options.bn(target);
            /** @type {number} */
            var k = 1;
            if (!(1 & target.limbs[0])) {
              callback(new options.exception.invalid("inverseMod: p must be odd"));
            }
            do {
              if (1 & t.limbs[0]) {
                if (!t.greaterEquals(d)) {
                  j = t;
                  t = d;
                  d = j;
                  j = obj;
                  obj = type;
                  type = j;
                }
                t.subM(d);
                t.normalize();
                if (!obj.greaterEquals(type)) {
                  obj.addM(target);
                }
                obj.subM(type);
              }
              t.halveM();
              if (1 & obj.limbs[0]) {
                obj.addM(target);
              }
              obj.normalize();
              obj.halveM();
              /** @type {number} */
              j = k = 0;
              for (;j < t.limbs.length;j++) {
                k |= t.limbs[j];
              }
            } while (k);
            return d.equals(1) || callback(new options.exception.invalid("inverseMod: p and x must be relatively prime")), type;
          },
          /**
           * @param {number} dataAndEvents
           * @return {?}
           */
          add : function(dataAndEvents) {
            return this.copy().addM(dataAndEvents);
          },
          /**
           * @param {number} expr
           * @return {?}
           */
          sub : function(expr) {
            return this.copy().subM(expr);
          },
          /**
           * @param {number} num
           * @return {?}
           */
          mul : function(num) {
            if ("number" == typeof num) {
              num = new this.i(num);
            }
            var index;
            var value;
            var iteratee = this.limbs;
            var digits = num.limbs;
            var length = iteratee.length;
            var numModules = digits.length;
            var first = new this.i;
            var pos = first.limbs;
            var maxMul = this.maxMul;
            /** @type {number} */
            index = 0;
            for (;index < this.limbs.length + num.limbs.length + 1;index++) {
              /** @type {number} */
              pos[index] = 0;
            }
            /** @type {number} */
            index = 0;
            for (;index < length;index++) {
              value = iteratee[index];
              /** @type {number} */
              num = 0;
              for (;num < numModules;num++) {
                pos[index + num] += value * digits[num];
              }
              if (!--maxMul) {
                maxMul = this.maxMul;
                first.cnormalize();
              }
            }
            return first.cnormalize().reduce();
          },
          /**
           * @return {?}
           */
          square : function() {
            return this.mul(this);
          },
          /**
           * @param {Array} b
           * @return {?}
           */
          power : function(b) {
            if ("number" == typeof b) {
              /** @type {Array} */
              b = [b];
            } else {
              if (b.limbs !== target) {
                b = b.normalize().limbs;
              }
            }
            var bi;
            var radix;
            var text = new this.i(1);
            var cx = this;
            /** @type {number} */
            bi = 0;
            for (;bi < b.length;bi++) {
              /** @type {number} */
              radix = 0;
              for (;radix < this.radix;radix++) {
                if (b[bi] & 1 << radix) {
                  text = text.mul(cx);
                }
                cx = cx.square();
              }
            }
            return text;
          },
          /**
           * @param {Object} node
           * @param {?} base
           * @return {?}
           */
          mulmod : function(node, base) {
            return this.mod(base).mul(node.mod(base)).mod(base);
          },
          /**
           * @param {string} inplace
           * @param {?} args
           * @return {?}
           */
          powermod : function(inplace, args) {
            var result = new options.bn(1);
            var value = new options.bn(this);
            var environment = new options.bn(inplace);
            for (;1 & environment.limbs[0] && (result = result.mulmod(value, args)), environment.halveM(), !environment.equals(0);) {
              value = value.mulmod(value, args);
            }
            return result.normalize().reduce();
          },
          /**
           * @return {?}
           */
          trim : function() {
            var cur;
            var eventPath = this.limbs;
            do {
              cur = eventPath.pop();
            } while (eventPath.length && 0 === cur);
            return eventPath.push(cur), this;
          },
          /**
           * @return {?}
           */
          reduce : function() {
            return this;
          },
          /**
           * @return {?}
           */
          fullReduce : function() {
            return this.normalize();
          },
          /**
           * @return {?}
           */
          normalize : function() {
            var i;
            /** @type {number} */
            var val = 0;
            var placeVal = this.placeVal;
            var sign = this.ipv;
            var values = this.limbs;
            var valuesLen = values.length;
            var STACK_ATTACHED = this.radixMask;
            /** @type {number} */
            i = 0;
            for (;i < valuesLen || 0 !== val && -1 !== val;i++) {
              /** @type {number} */
              val = ((val = (values[i] || 0) + val) - (values[i] = val & STACK_ATTACHED)) * sign;
            }
            return-1 === val && (values[i - 1] -= placeVal), this;
          },
          /**
           * @return {?}
           */
          cnormalize : function() {
            var i;
            /** @type {number} */
            var h = 0;
            var ifrHeight = this.ipv;
            var codeSegments = this.limbs;
            var selectorCount = codeSegments.length;
            var mask = this.radixMask;
            /** @type {number} */
            i = 0;
            for (;i < selectorCount - 1;i++) {
              /** @type {number} */
              h = ((h = codeSegments[i] + h) - (codeSegments[i] = h & mask)) * ifrHeight;
            }
            return codeSegments[i] += h, this;
          },
          /**
           * @param {string} str
           * @return {?}
           */
          toBits : function(str) {
            this.fullReduce();
            str = str || (this.exponent || this.bitLength());
            /** @type {number} */
            var cDigit = Math.floor((str - 1) / 24);
            var w = options.bitArray;
            /** @type {Array} */
            var out = [w.partial((str + 7 & -8) % this.radix || this.radix, this.getLimb(cDigit))];
            cDigit--;
            for (;0 <= cDigit;cDigit--) {
              out = w.concat(out, [w.partial(Math.min(this.radix, str), this.getLimb(cDigit))]);
              str -= this.radix;
            }
            return out;
          },
          /**
           * @return {?}
           */
          bitLength : function() {
            this.fullReduce();
            /** @type {number} */
            var radix = this.radix * (this.limbs.length - 1);
            var body = this.limbs[this.limbs.length - 1];
            for (;body;body >>>= 1) {
              radix++;
            }
            return radix + 7 & -8;
          }
        };
        /**
         * @param {Object} data
         * @return {?}
         */
        options.bn.fromBits = function(data) {
          var ul = new this;
          /** @type {Array} */
          var out = [];
          var callback = options.bitArray;
          var self = this.prototype;
          /** @type {number} */
          var to = Math.min(this.bitLength || 4294967296, callback.bitLength(data));
          var i = to % self.radix || self.radix;
          out[0] = callback.extract(data, 0, i);
          for (;i < to;i += self.radix) {
            out.unshift(callback.extract(data, i, self.radix));
          }
          return ul.limbs = out, ul;
        };
        /** @type {number} */
        options.bn.prototype.ipv = 1 / (options.bn.prototype.placeVal = Math.pow(2, options.bn.prototype.radix));
        /** @type {number} */
        options.bn.prototype.radixMask = (1 << options.bn.prototype.radix) - 1;
        /**
         * @param {number} opt_attributes
         * @param {Array} fileExtensions
         * @return {?}
         */
        options.bn.pseudoMersennePrime = function(opt_attributes, fileExtensions) {
          /**
           * @param {string} val
           * @return {undefined}
           */
          function self(val) {
            this.initWith(val);
          }
          var i;
          var off;
          var args = self.prototype = new options.bn;
          /** @type {number} */
          i = args.modOffset = Math.ceil(off = opt_attributes / args.radix);
          /** @type {number} */
          args.exponent = opt_attributes;
          /** @type {Array} */
          args.offset = [];
          /** @type {Array} */
          args.factor = [];
          /** @type {number} */
          args.minOffset = i;
          /** @type {number} */
          args.fullMask = 0;
          /** @type {Array} */
          args.fullOffset = [];
          /** @type {Array} */
          args.fullFactor = [];
          args.modulus = self.modulus = new options.bn(Math.pow(2, opt_attributes));
          /** @type {number} */
          args.fullMask = 0 | -Math.pow(2, opt_attributes % args.radix);
          /** @type {number} */
          i = 0;
          for (;i < fileExtensions.length;i++) {
            /** @type {number} */
            args.offset[i] = Math.floor(fileExtensions[i][0] / args.radix - off);
            /** @type {number} */
            args.fullOffset[i] = Math.ceil(fileExtensions[i][0] / args.radix - off);
            /** @type {number} */
            args.factor[i] = fileExtensions[i][1] * Math.pow(0.5, opt_attributes - fileExtensions[i][0] + args.offset[i] * args.radix);
            /** @type {number} */
            args.fullFactor[i] = fileExtensions[i][1] * Math.pow(0.5, opt_attributes - fileExtensions[i][0] + args.fullOffset[i] * args.radix);
            args.modulus.addM(new options.bn(Math.pow(2, fileExtensions[i][0]) * fileExtensions[i][1]));
            /** @type {number} */
            args.minOffset = Math.min(args.minOffset, -args.offset[i]);
          }
          return args.i = self, args.modulus.cnormalize(), args.reduce = function() {
            var minOffset;
            var x;
            var ratio;
            var i;
            var max = this.modOffset;
            var items = this.limbs;
            var offset = this.offset;
            var cnl = this.offset.length;
            var flexWidths = this.factor;
            minOffset = this.minOffset;
            for (;items.length > max;) {
              ratio = items.pop();
              i = items.length;
              /** @type {number} */
              x = 0;
              for (;x < cnl;x++) {
                items[i + offset[x]] -= flexWidths[x] * ratio;
              }
              if (!--minOffset) {
                items.push(0);
                this.cnormalize();
                minOffset = this.minOffset;
              }
            }
            return this.cnormalize(), this;
          }, args.fa = -1 === args.fullMask ? args.reduce : function() {
            var i;
            var fact;
            var stack = this.limbs;
            /** @type {number} */
            var sp = stack.length - 1;
            if (this.reduce(), sp === this.modOffset - 1) {
              /** @type {number} */
              fact = stack[sp] & this.fullMask;
              stack[sp] -= fact;
              /** @type {number} */
              i = 0;
              for (;i < this.fullOffset.length;i++) {
                stack[sp + this.fullOffset[i]] -= this.fullFactor[i] * fact;
              }
              this.normalize();
            }
          }, args.fullReduce = function() {
            var fact;
            var i;
            this.fa();
            this.addM(this.modulus);
            this.addM(this.modulus);
            this.normalize();
            this.fa();
            i = this.limbs.length;
            for (;i < this.modOffset;i++) {
              /** @type {number} */
              this.limbs[i] = 0;
            }
            fact = this.greaterEquals(this.modulus);
            /** @type {number} */
            i = 0;
            for (;i < this.limbs.length;i++) {
              this.limbs[i] -= this.modulus.limbs[i] * fact;
            }
            return this.cnormalize(), this;
          }, args.inverse = function() {
            return this.power(this.modulus.sub(2));
          }, self.fromBits = options.bn.fromBits, self;
        };
        /** @type {function (number, Array): ?} */
        var throttledUpdate = options.bn.pseudoMersennePrime;
        options.bn.prime = {
          p127 : throttledUpdate(127, [[0, -1]]),
          p25519 : throttledUpdate(255, [[0, -19]]),
          p192k : throttledUpdate(192, [[32, -1], [12, -1], [8, -1], [7, -1], [6, -1], [3, -1], [0, -1]]),
          p224k : throttledUpdate(224, [[32, -1], [12, -1], [11, -1], [9, -1], [7, -1], [4, -1], [1, -1], [0, -1]]),
          p256k : throttledUpdate(256, [[32, -1], [9, -1], [8, -1], [7, -1], [6, -1], [4, -1], [0, -1]]),
          p192 : throttledUpdate(192, [[0, -1], [64, -1]]),
          p224 : throttledUpdate(224, [[0, 1], [96, -1]]),
          p256 : throttledUpdate(256, [[0, -1], [96, 1], [192, 1], [224, -1]]),
          p384 : throttledUpdate(384, [[0, -1], [32, 1], [96, -1], [128, -1]]),
          p521 : throttledUpdate(521, [[0, -1]])
        };
        /**
         * @param {number} obj
         * @param {number} value
         * @return {?}
         */
        options.bn.random = function(obj, value) {
          if ("object" != typeof obj) {
            obj = new options.bn(obj);
          }
          var ret;
          var index;
          var size = obj.limbs.length;
          var ns = obj.limbs[size - 1] + 1;
          var arr = new options.bn;
          for (;;) {
            do {
              if (0 > (ret = options.random.randomWords(size, value))[size - 1]) {
                ret[size - 1] += 4294967296;
              }
            } while (Math.floor(ret[size - 1] / ns) === Math.floor(4294967296 / ns));
            ret[size - 1] %= ns;
            /** @type {number} */
            index = 0;
            for (;index < size - 1;index++) {
              ret[index] &= obj.radixMask;
            }
            if (arr.limbs = ret, !arr.greaterEquals(obj)) {
              return arr;
            }
          }
        };
        options.ecc = {};
        /**
         * @param {?} p
         * @param {number} x
         * @param {Object} y
         * @return {undefined}
         */
        options.ecc.point = function(p, x, y) {
          if (x === target) {
            /** @type {boolean} */
            this.isIdentity = pdataCur;
          } else {
            if (x instanceof options.bn) {
              x = new p.field(x);
            }
            if (y instanceof options.bn) {
              y = new p.field(y);
            }
            /** @type {number} */
            this.x = x;
            /** @type {Object} */
            this.y = y;
            /** @type {boolean} */
            this.isIdentity = FALSE;
          }
          this.curve = p;
        };
        options.ecc.point.prototype = {
          /**
           * @return {?}
           */
          toJac : function() {
            return new options.ecc.pointJac(this.curve, this.x, this.y, new this.curve.field(1));
          },
          /**
           * @param {number} value
           * @return {?}
           */
          mult : function(value) {
            return this.toJac().mult(value, this).toAffine();
          },
          /**
           * @param {Array} dataAndEvents
           * @param {Array} num
           * @param {Array} prop
           * @return {?}
           */
          mult2 : function(dataAndEvents, num, prop) {
            return this.toJac().mult2(dataAndEvents, this, num, prop).toAffine();
          },
          /**
           * @return {?}
           */
          multiples : function() {
            var eventPath;
            var e;
            var previous;
            if (this.ca === target) {
              previous = this.toJac().doubl();
              /** @type {Array} */
              eventPath = this.ca = [new options.ecc.point(this.curve), this, previous.toAffine()];
              /** @type {number} */
              e = 3;
              for (;16 > e;e++) {
                previous = previous.add(this);
                eventPath.push(previous.toAffine());
              }
            }
            return this.ca;
          },
          /**
           * @return {?}
           */
          isValid : function() {
            return this.y.square().equals(this.curve.b.add(this.x.mul(this.curve.a.add(this.x.square()))));
          },
          /**
           * @return {?}
           */
          toBits : function() {
            return options.bitArray.concat(this.x.toBits(), this.y.toBits());
          }
        };
        /**
         * @param {?} curve
         * @param {number} value
         * @param {?} yp
         * @param {number} z1
         * @return {undefined}
         */
        options.ecc.pointJac = function(curve, value, yp, z1) {
          if (value === target) {
            /** @type {boolean} */
            this.isIdentity = pdataCur;
          } else {
            /** @type {number} */
            this.x = value;
            this.y = yp;
            /** @type {number} */
            this.z = z1;
            /** @type {boolean} */
            this.isIdentity = FALSE;
          }
          this.curve = curve;
        };
        options.ecc.pointJac.prototype = {
          /**
           * @param {number} dataAndEvents
           * @return {?}
           */
          add : function(dataAndEvents) {
            var t;
            var base;
            var n;
            var cDigit;
            return this.curve !== dataAndEvents.curve && callback("sjcl['ecc']['add'](): Points must be on the same curve to add them!"), this.isIdentity ? dataAndEvents.toJac() : dataAndEvents.isIdentity ? this : (t = this.z.square(), (base = dataAndEvents.x.mul(t).subM(this.x)).equals(0) ? this.y.equals(dataAndEvents.y.mul(t.mul(this.z))) ? this.doubl() : new options.ecc.pointJac(this.curve) : (t = dataAndEvents.y.mul(t.mul(this.z)).subM(this.y), n = base.square(), dataAndEvents = t.square(), cDigit = 
            base.square().mul(base).addM(this.x.add(this.x).mul(n)), dataAndEvents = dataAndEvents.subM(cDigit), t = this.x.mul(n).subM(dataAndEvents).mul(t), n = this.y.mul(base.square().mul(base)), t = t.subM(n), base = this.z.mul(base), new options.ecc.pointJac(this.curve, dataAndEvents, t, base)));
          },
          /**
           * @return {?}
           */
          doubl : function() {
            if (this.isIdentity) {
              return this;
            }
            var cDigit = (dataAndEvents = this.y.square()).mul(this.x.mul(4));
            var chr2 = dataAndEvents.square().mul(8);
            var dataAndEvents = this.z.square();
            var four = this.curve.a.toString() == (new options.bn(-3)).toString() ? this.x.sub(dataAndEvents).mul(3).mul(this.x.add(dataAndEvents)) : this.x.square().mul(3).add(dataAndEvents.square().mul(this.curve.a));
            return dataAndEvents = four.square().subM(cDigit).subM(cDigit), cDigit = cDigit.sub(dataAndEvents).mul(four).subM(chr2), chr2 = this.y.add(this.y).mul(this.z), new options.ecc.pointJac(this.curve, dataAndEvents, cDigit, chr2);
          },
          /**
           * @return {?}
           */
          toAffine : function() {
            if (this.isIdentity || this.z.equals(0)) {
              return new options.ecc.point(this.curve);
            }
            var cDigit = this.z.inverse();
            var diff = cDigit.square();
            return new options.ecc.point(this.curve, this.x.mul(diff).fullReduce(), this.y.mul(diff.mul(cDigit)).fullReduce());
          },
          /**
           * @param {Array} value
           * @param {?} dataAndEvents
           * @return {?}
           */
          mult : function(value, dataAndEvents) {
            if ("number" == typeof value) {
              /** @type {Array} */
              value = [value];
            } else {
              if (value.limbs !== target) {
                value = value.normalize().limbs;
              }
            }
            var i;
            var betashift;
            var v1 = (new options.ecc.point(this.curve)).toJac();
            var clone = dataAndEvents.multiples();
            /** @type {number} */
            i = value.length - 1;
            for (;0 <= i;i--) {
              /** @type {number} */
              betashift = options.bn.prototype.radix - 4;
              for (;0 <= betashift;betashift -= 4) {
                v1 = v1.doubl().doubl().doubl().doubl().add(clone[value[i] >> betashift & 15]);
              }
            }
            return v1;
          },
          /**
           * @param {Array} a
           * @param {(Array|Element)} num
           * @param {Array} b
           * @param {number} attribute
           * @return {?}
           */
          mult2 : function(a, num, b, attribute) {
            if ("number" == typeof a) {
              /** @type {Array} */
              a = [a];
            } else {
              if (a.limbs !== target) {
                a = a.normalize().limbs;
              }
            }
            if ("number" == typeof b) {
              /** @type {Array} */
              b = [b];
            } else {
              if (b.limbs !== target) {
                b = b.normalize().limbs;
              }
            }
            var i;
            var doubl = (new options.ecc.point(this.curve)).toJac();
            num = num.multiples();
            var bits;
            var M;
            var clone = attribute.multiples();
            /** @type {number} */
            attribute = Math.max(a.length, b.length) - 1;
            for (;0 <= attribute;attribute--) {
              /** @type {number} */
              bits = 0 | a[attribute];
              /** @type {number} */
              M = 0 | b[attribute];
              /** @type {number} */
              i = options.bn.prototype.radix - 4;
              for (;0 <= i;i -= 4) {
                doubl = doubl.doubl().doubl().doubl().doubl().add(num[bits >> i & 15]).add(clone[M >> i & 15]);
              }
            }
            return doubl;
          },
          /**
           * @return {?}
           */
          isValid : function() {
            var p1 = (p2 = this.z.square()).square();
            var p2 = p1.mul(p2);
            return this.y.square().equals(this.curve.b.mul(p2).add(this.x.mul(this.curve.a.mul(p1).add(this.x.square()))));
          }
        };
        /**
         * @param {Object} Vec2
         * @param {?} tension
         * @param {Object} a
         * @param {Object} v1
         * @param {Object} x
         * @param {Object} u
         * @return {undefined}
         */
        options.ecc.curve = function(Vec2, tension, a, v1, x, u) {
          /** @type {Object} */
          this.field = Vec2;
          this.r = new options.bn(tension);
          this.a = new Vec2(a);
          this.b = new Vec2(v1);
          this.G = new options.ecc.point(this, new Vec2(x), new Vec2(u));
        };
        /**
         * @param {Object} data
         * @return {?}
         */
        options.ecc.curve.prototype.fromBits = function(data) {
          var jQuery = options.bitArray;
          /** @type {number} */
          var ret = this.field.prototype.exponent + 7 & -8;
          return(data = new options.ecc.point(this, this.field.fromBits(jQuery.bitSlice(data, 0, ret)), this.field.fromBits(jQuery.bitSlice(data, ret, 2 * ret)))).isValid() || callback(new options.exception.corrupt("not on the curve!")), data;
        };
        options.ecc.curves = {
          c192 : new options.ecc.curve(options.bn.prime.p192, "0xffffffffffffffffffffffff99def836146bc9b1b4d22831", -3, "0x64210519e59c80e70fa7e9ab72243049feb8deecc146b9b1", "0x188da80eb03090f67cbf20eb43a18800f4ff0afd82ff1012", "0x07192b95ffc8da78631011ed6b24cdd573f977a11e794811"),
          c224 : new options.ecc.curve(options.bn.prime.p224, "0xffffffffffffffffffffffffffff16a2e0b8f03e13dd29455c5c2a3d", -3, "0xb4050a850c04b3abf54132565044b0b7d7bfd8ba270b39432355ffb4", "0xb70e0cbd6bb4bf7f321390b94a03c1d356c21122343280d6115c1d21", "0xbd376388b5f723fb4c22dfe6cd4375a05a07476444d5819985007e34"),
          c256 : new options.ecc.curve(options.bn.prime.p256, "0xffffffff00000000ffffffffffffffffbce6faada7179e84f3b9cac2fc632551", -3, "0x5ac635d8aa3a93e7b3ebbd55769886bc651d06b0cc53b0f63bce3c3e27d2604b", "0x6b17d1f2e12c4247f8bce6e563a440f277037d812deb33a0f4a13945d898c296", "0x4fe342e2fe1a7f9b8ee7eb4a7c0f9e162bce33576b315ececbb6406837bf51f5"),
          c384 : new options.ecc.curve(options.bn.prime.p384, "0xffffffffffffffffffffffffffffffffffffffffffffffffc7634d81f4372ddf581a0db248b0a77aecec196accc52973", -3, "0xb3312fa7e23ee7e4988e056be3f82d19181d9c6efe8141120314088f5013875ac656398d8a2ed19d2a85c8edd3ec2aef", "0xaa87ca22be8b05378eb1c71ef320ad746e1d3b628ba79b9859f741e082542a385502f25dbf55296c3a545e3872760ab7", "0x3617de4a96262c6f5d9e98bf9292dc29f8f41dbd289a147ce9da3113b5f0b8c00a60b1ce1d7e819d7a431d7c90ea0e5f"),
          k192 : new options.ecc.curve(options.bn.prime.p192k, "0xfffffffffffffffffffffffe26f2fc170f69466a74defd8d", 0, 3, "0xdb4ff10ec057e9ae26b07d0280b7f4341da5d1b1eae06c7d", "0x9b2f2f6d9c5628a7844163d015be86344082aa88d95e2f9d"),
          k224 : new options.ecc.curve(options.bn.prime.p224k, "0x010000000000000000000000000001dce8d2ec6184caf0a971769fb1f7", 0, 5, "0xa1455b334df099df30fc28a169a467e9e47075a90f7e650eb6b7a45c", "0x7e089fed7fba344282cafbd6f7e319f7c0b0bd59e2ca4bdb556d61a5"),
          k256 : new options.ecc.curve(options.bn.prime.p256k, "0xfffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 0, 7, "0x79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", "0x483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8")
        };
        options.ecc.basicKey = {
          /**
           * @param {Node} client
           * @param {Object} msgs
           * @return {undefined}
           */
          publicKey : function(client, msgs) {
            /** @type {Node} */
            this.l = client;
            this.s = client.r.bitLength();
            this.I = msgs instanceof Array ? client.fromBits(msgs) : msgs;
            /**
             * @return {?}
             */
            this.get = function() {
              var x = this.I.toBits();
              var pkgfile = options.bitArray.bitLength(x);
              return{
                x : options.bitArray.bitSlice(x, 0, pkgfile / 2),
                y : x = options.bitArray.bitSlice(x, pkgfile / 2)
              };
            };
          },
          /**
           * @param {Node} self
           * @param {Function} dataAndEvents
           * @return {undefined}
           */
          secretKey : function(self, dataAndEvents) {
            /** @type {Node} */
            this.l = self;
            this.s = self.r.bitLength();
            /** @type {Function} */
            this.H = dataAndEvents;
            /**
             * @return {?}
             */
            this.get = function() {
              return this.H.toBits();
            };
          }
        };
        /**
         * @param {string} i
         * @return {?}
         */
        options.ecc.basicKey.generateKeys = function(i) {
          return function(start, isXML, chars) {
            return "number" == typeof(start = start || 256) && ((start = options.ecc.curves["c" + start]) === target && callback(new options.exception.invalid("no such curve"))), chars = chars || options.bn.random(start.r, isXML), isXML = start.G.mult(chars), {
              pub : new options.ecc[i].publicKey(start, isXML),
              sec : new options.ecc[i].secretKey(start, chars)
            };
          };
        };
        options.ecc.elGamal = {
          generateKeys : options.ecc.basicKey.generateKeys("elGamal"),
          /**
           * @param {?} dataAndEvents
           * @param {?} deepDataAndEvents
           * @return {undefined}
           */
          publicKey : function(dataAndEvents, deepDataAndEvents) {
            options.ecc.basicKey.publicKey.apply(this, arguments);
          },
          /**
           * @param {?} dataAndEvents
           * @param {?} deepDataAndEvents
           * @return {undefined}
           */
          secretKey : function(dataAndEvents, deepDataAndEvents) {
            options.ecc.basicKey.secretKey.apply(this, arguments);
          }
        };
        options.ecc.elGamal.publicKey.prototype = {
          /**
           * @param {number} isXML
           * @return {?}
           */
          kem : function(isXML) {
            isXML = options.bn.random(this.l.r, isXML);
            var Tag = this.l.G.mult(isXML).toBits();
            return{
              key : options.hash.sha256.hash(this.I.mult(isXML).toBits()),
              tag : Tag
            };
          }
        };
        options.ecc.elGamal.secretKey.prototype = {
          /**
           * @param {Object} msgs
           * @return {?}
           */
          unkem : function(msgs) {
            return options.hash.sha256.hash(this.l.fromBits(msgs).mult(this.H).toBits());
          },
          /**
           * @param {Object} FXStage
           * @return {?}
           */
          dh : function(FXStage) {
            return options.hash.sha256.hash(FXStage.I.mult(this.H).toBits());
          },
          /**
           * @param {Object} FXStage
           * @return {?}
           */
          dhJavaEc : function(FXStage) {
            return FXStage.I.mult(this.H).x.toBits();
          }
        };
        options.ecc.ecdsa = {
          generateKeys : options.ecc.basicKey.generateKeys("ecdsa")
        };
        /**
         * @param {?} dataAndEvents
         * @param {?} deepDataAndEvents
         * @return {undefined}
         */
        options.ecc.ecdsa.publicKey = function(dataAndEvents, deepDataAndEvents) {
          options.ecc.basicKey.publicKey.apply(this, arguments);
        };
        options.ecc.ecdsa.publicKey.prototype = {
          /**
           * @param {Object} msgs
           * @param {Object} done
           * @param {boolean} value
           * @return {?}
           */
          verify : function(msgs, done, value) {
            if (options.bitArray.bitLength(msgs) > this.s) {
              msgs = options.bitArray.clamp(msgs, this.s);
            }
            var d = options.bitArray;
            var base = this.l.r;
            var node = this.s;
            var recurring = options.bn.fromBits(d.bitSlice(done, 0, node));
            var cDigit = (d = options.bn.fromBits(d.bitSlice(done, node, 2 * node)), value ? d : d.inverseMod(base));
            if (node = options.bn.fromBits(msgs).mul(cDigit).mod(base), cDigit = recurring.mul(cDigit).mod(base), node = this.l.G.mult2(node, cDigit, this.I).x, recurring.equals(0) || (d.equals(0) || (recurring.greaterEquals(base) || (d.greaterEquals(base) || !node.equals(recurring))))) {
              if (value === target) {
                return this.verify(msgs, done, pdataCur);
              }
              callback(new options.exception.corrupt("signature didn't check out"));
            }
            return pdataCur;
          }
        };
        /**
         * @param {?} dataAndEvents
         * @param {?} deepDataAndEvents
         * @return {undefined}
         */
        options.ecc.ecdsa.secretKey = function(dataAndEvents, deepDataAndEvents) {
          options.ecc.basicKey.secretKey.apply(this, arguments);
        };
        options.ecc.ecdsa.secretKey.prototype = {
          /**
           * @param {Object} inplace
           * @param {number} value
           * @param {(Array|Function|string)} callback
           * @param {number} num
           * @return {?}
           */
          sign : function(inplace, value, callback, num) {
            if (options.bitArray.bitLength(inplace) > this.s) {
              inplace = options.bitArray.clamp(inplace, this.s);
            }
            var base = this.l.r;
            var errStr = base.bitLength();
            return num = num || options.bn.random(base.sub(1), value).add(1), value = this.l.G.mult(num).x.mod(base), inplace = options.bn.fromBits(inplace).add(value.mul(this.H)), callback = callback ? inplace.inverseMod(base).mul(num).mod(base) : inplace.mul(num.inverseMod(base)).mod(base), options.bitArray.concat(value.toBits(errStr), callback.toBits(errStr));
          }
        };
        options.keyexchange.srp = {
          /**
           * @param {Object} inplace
           * @param {string} deepDataAndEvents
           * @param {?} walkers
           * @param {Object} node
           * @return {?}
           */
          makeVerifier : function(inplace, deepDataAndEvents, walkers, node) {
            return inplace = options.keyexchange.srp.makeX(inplace, deepDataAndEvents, walkers), inplace = options.bn.fromBits(inplace), node.g.powermod(inplace, node.N);
          },
          /**
           * @param {string} data
           * @param {string} deepDataAndEvents
           * @param {?} obj
           * @return {?}
           */
          makeX : function(data, deepDataAndEvents, obj) {
            return data = options.hash.sha1.hash(data + ":" + deepDataAndEvents), options.hash.sha1.hash(options.bitArray.concat(obj, data));
          },
          /**
           * @param {(number|string)} string
           * @return {?}
           */
          knownGroup : function(string) {
            return "string" != typeof string && (string = string.toString()), options.keyexchange.srp.Z || options.keyexchange.srp.qa(), options.keyexchange.srp.ba[string];
          },
          Z : FALSE,
          /**
           * @return {undefined}
           */
          qa : function() {
            var i;
            var node;
            /** @type {number} */
            i = 0;
            for (;i < options.keyexchange.srp.aa.length;i++) {
              node = options.keyexchange.srp.aa[i].toString();
              (node = options.keyexchange.srp.ba[node]).N = new options.bn(node.N);
              node.g = new options.bn(node.g);
            }
            /** @type {boolean} */
            options.keyexchange.srp.Z = pdataCur;
          },
          aa : [1024, 1536, 2048],
          ba : {
            1024 : {
              N : "EEAF0AB9ADB38DD69C33F80AFA8FC5E86072618775FF3C0B9EA2314C9C256576D674DF7496EA81D3383B4813D692C6E0E0D5D8E250B98BE48E495C1D6089DAD15DC7D7B46154D6B6CE8EF4AD69B15D4982559B297BCF1885C529F566660E57EC68EDBC3C05726CC02FD4CBF4976EAA9AFD5138FE8376435B9FC61D2FC0EB06E3",
              g : 2
            },
            1536 : {
              N : "9DEF3CAFB939277AB1F12A8617A47BBBDBA51DF499AC4C80BEEEA9614B19CC4D5F4F5F556E27CBDE51C6A94BE4607A291558903BA0D0F84380B655BB9A22E8DCDF028A7CEC67F0D08134B1C8B97989149B609E0BE3BAB63D47548381DBC5B1FC764E3F4B53DD9DA1158BFD3E2B9C8CF56EDF019539349627DB2FD53D24B7C48665772E437D6C7F8CE442734AF7CCB7AE837C264AE3A9BEB87F8A2FE9B8B5292E5A021FFF5E91479E8CE7A28C2442C6F315180F93499A234DCF76E3FED135F9BB",
              g : 2
            },
            2048 : {
              N : "AC6BDB41324A9A9BF166DE5E1389582FAF72B6651987EE07FC3192943DB56050A37329CBB4A099ED8193E0757767A13DD52312AB4B03310DCD7F48A9DA04FD50E8083969EDB767B0CF6095179A163AB3661A05FBD5FAAAE82918A9962F0B93B855F97993EC975EEAA80D740ADBF4FF747359D041D5C33EA71D281E446B14773BCA97B43A23FB801676BD207A436C6481F1D2B9078717461A5B9D32E688F87748544523B524B0D57D5EA77A2775D2ECFA032CFBDBF52FB3786160279004E57AE6AF874E7303CE53299CCC041C7BC308D82A5698F3A8D0C38271AE35F8E9DBFBB694B5C803D89F7AE435DE236D525F54759B65E372FCD68EF20FA7111F9E4AFF73",
              g : 2
            }
          }
        };
      }, function(module, dataAndEvents, callback) {
        (function(arg) {
          /**
           * @param {Object} context
           * @return {?}
           */
          var runInContext = function(context) {
            return context && (context.Math == Math && context);
          };
          module.exports = runInContext("object" == typeof globalThis && globalThis) || (runInContext("object" == typeof window && window) || (runInContext("object" == typeof self && self) || (runInContext("object" == typeof arg && arg) || Function("return this")())));
        }).call(this, callback(23));
      }, function(module, dataAndEvents, require) {
        var range = require(3);
        var flag = require(45);
        var inspect = require(7);
        var Block = require(49);
        var a = require(50);
        var nodes = require(73);
        var obj = flag("wks");
        var x = range.Symbol;
        var recurse = nodes ? x : x && x.withoutSetter || Block;
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          return inspect(obj, expectedNumberOfNonCommentArgs) || (a && inspect(x, expectedNumberOfNonCommentArgs) ? obj[expectedNumberOfNonCommentArgs] = x[expectedNumberOfNonCommentArgs] : obj[expectedNumberOfNonCommentArgs] = recurse("Symbol." + expectedNumberOfNonCommentArgs)), obj[expectedNumberOfNonCommentArgs];
        };
      }, function(module, dataAndEvents) {
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          try {
            return!!expectedNumberOfNonCommentArgs();
          } catch (t) {
            return true;
          }
        };
      }, function(module, dataAndEvents, require) {
        var getActual = require(13);
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          if (!getActual(expectedNumberOfNonCommentArgs)) {
            throw TypeError(String(expectedNumberOfNonCommentArgs) + " is not an object");
          }
          return expectedNumberOfNonCommentArgs;
        };
      }, function(module, dataAndEvents) {
        /** @type {function (this:Object, *): boolean} */
        var has = {}.hasOwnProperty;
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Function} proto
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, proto) {
          return has.call(expectedNumberOfNonCommentArgs, proto);
        };
      }, function(module, dataAndEvents, require) {
        var names = require(3);
        var inspect = require(36).f;
        var flag = require(10);
        var callback = require(15);
        var capitalize = require(31);
        var debug = require(78);
        var getActual = require(57);
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Object} object
         * @return {undefined}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, object) {
          var a;
          var key;
          var desc;
          var obj;
          var elem;
          var name = expectedNumberOfNonCommentArgs.target;
          var n = expectedNumberOfNonCommentArgs.global;
          var hasExt = expectedNumberOfNonCommentArgs.stat;
          if (a = n ? names : hasExt ? names[name] || capitalize(name, {}) : (names[name] || {}).prototype) {
            for (key in object) {
              if (obj = object[key], desc = expectedNumberOfNonCommentArgs.noTargetGet ? (elem = inspect(a, key)) && elem.value : a[key], !getActual(n ? key : name + (hasExt ? "." : "#") + key, expectedNumberOfNonCommentArgs.forced) && void 0 !== desc) {
                if (typeof obj == typeof desc) {
                  continue;
                }
                debug(obj, desc);
              }
              if (expectedNumberOfNonCommentArgs.sham || desc && desc.sham) {
                flag(obj, "sham", true);
              }
              callback(a, key, obj, expectedNumberOfNonCommentArgs);
            }
          }
        };
      }, function(dataAndEvents, fnc, callback) {
        (function(video) {
          /** @type {null} */
          var pdataOld = null;
          /** @type {null} */
          var events = null;
          if ("undefined" != typeof WorkerGlobalScope && self instanceof WorkerGlobalScope) {
            if (pdataOld = self.crypto || self.msCrypto) {
              events = pdataOld.subtle;
            }
          } else {
            if (pdataOld = video.crypto || video.msCrypto) {
              events = pdataOld.subtle;
            }
          }
          fnc.a = {
            WebCrypto : pdataOld,
            SubtleCrypto : events
          };
        }).call(this, callback(23));
      }, function(module, dataAndEvents, require) {
        var Block = require(11);
        var object = require(12);
        var group = require(24);
        /** @type {function (number, Object, boolean): ?} */
        module.exports = Block ? function(expectedNumberOfNonCommentArgs, methodName, value) {
          return object.f(expectedNumberOfNonCommentArgs, methodName, group(1, value));
        } : function(expectedNumberOfNonCommentArgs, proto, value) {
          return expectedNumberOfNonCommentArgs[proto] = value, expectedNumberOfNonCommentArgs;
        };
      }, function(module, dataAndEvents, require) {
        var getActual = require(5);
        /** @type {boolean} */
        module.exports = !getActual(function() {
          return 7 != Object.defineProperty({}, 1, {
            /**
             * @return {?}
             */
            get : function() {
              return 7;
            }
          })[1];
        });
      }, function(dataAndEvents, state, require) {
        var url = require(11);
        var Block = require(47);
        var getName = require(6);
        var expect = require(48);
        /** @type {function (Object, string, Object): Object} */
        var defineProperty = Object.defineProperty;
        /** @type {Function} */
        state.f = url ? defineProperty : function(expectedNumberOfNonCommentArgs, name, desc) {
          if (getName(expectedNumberOfNonCommentArgs), name = expect(name, true), getName(desc), Block) {
            try {
              return defineProperty(expectedNumberOfNonCommentArgs, name, desc);
            } catch (t) {
            }
          }
          if ("get" in desc || "set" in desc) {
            throw TypeError("Accessors not supported");
          }
          return "value" in desc && (expectedNumberOfNonCommentArgs[name] = desc.value), expectedNumberOfNonCommentArgs;
        };
      }, function(module, dataAndEvents) {
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          return "object" == typeof expectedNumberOfNonCommentArgs ? null !== expectedNumberOfNonCommentArgs : "function" == typeof expectedNumberOfNonCommentArgs;
        };
      }, function(module, dataAndEvents, expression) {
        var modifyProps = expression(37);
        var obj = expression(3);
        /**
         * @param {number} opts
         * @return {?}
         */
        var inspect = function(opts) {
          return "function" == typeof opts ? opts : void 0;
        };
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Function} name
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, name) {
          return arguments.length < 2 ? inspect(modifyProps[expectedNumberOfNonCommentArgs]) || inspect(obj[expectedNumberOfNonCommentArgs]) : modifyProps[expectedNumberOfNonCommentArgs] && modifyProps[expectedNumberOfNonCommentArgs][name] || obj[expectedNumberOfNonCommentArgs] && obj[expectedNumberOfNonCommentArgs][name];
        };
      }, function(module, dataAndEvents, require) {
        var Block = require(3);
        var assert = require(10);
        var inspect = require(7);
        var next = require(31);
        var getActual = require(33);
        var nodes = require(18);
        var len = nodes.get;
        var register = nodes.enforce;
        /** @type {Array.<string>} */
        var s = String(String).split("String");
        (module.exports = function(expectedNumberOfNonCommentArgs, name, value, db) {
          /** @type {boolean} */
          var u = !!db && !!db.unsafe;
          /** @type {boolean} */
          var c = !!db && !!db.enumerable;
          /** @type {boolean} */
          var l = !!db && !!db.noTargetGet;
          if ("function" == typeof value) {
            if (!("string" != typeof name)) {
              if (!inspect(value, "name")) {
                assert(value, "name", name);
              }
            }
            /** @type {string} */
            register(value).source = s.join("string" == typeof name ? name : "");
          }
          if (expectedNumberOfNonCommentArgs !== Block) {
            if (u) {
              if (!l) {
                if (expectedNumberOfNonCommentArgs[name]) {
                  /** @type {boolean} */
                  c = true;
                }
              }
            } else {
              delete expectedNumberOfNonCommentArgs[name];
            }
            if (c) {
              /** @type {boolean} */
              expectedNumberOfNonCommentArgs[name] = value;
            } else {
              assert(expectedNumberOfNonCommentArgs, name, value);
            }
          } else {
            if (c) {
              /** @type {boolean} */
              expectedNumberOfNonCommentArgs[name] = value;
            } else {
              next(name, value);
            }
          }
        })(Function.prototype, "toString", function() {
          return "function" == typeof this && len(this).source || getActual(this);
        });
      }, function(module, dataAndEvents) {
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          if ("function" != typeof expectedNumberOfNonCommentArgs) {
            throw TypeError(String(expectedNumberOfNonCommentArgs) + " is not a function");
          }
          return expectedNumberOfNonCommentArgs;
        };
      }, function(module, dataAndEvents) {
        /** @type {boolean} */
        module.exports = false;
      }, function($, dataAndEvents, require) {
        var remove;
        var query;
        var walk;
        var helper = require(74);
        var Block = require(3);
        var getActual = require(13);
        var assert = require(10);
        var filter = require(7);
        var input = require(34);
        var nodes = require(35);
        var SimulatedScope = Block.WeakMap;
        if (helper) {
          var scope = new SimulatedScope;
          var getter = scope.get;
          var _$watch = scope.has;
          var fn = scope.set;
          /**
           * @param {?} first
           * @param {?} expectedHashCode
           * @return {?}
           */
          remove = function(first, expectedHashCode) {
            return fn.call(scope, first, expectedHashCode), expectedHashCode;
          };
          /**
           * @param {?} first
           * @return {?}
           */
          query = function(first) {
            return getter.call(scope, first) || {};
          };
          /**
           * @param {?} tree
           * @return {?}
           */
          walk = function(tree) {
            return _$watch.call(scope, tree);
          };
        } else {
          var i = input("state");
          /** @type {boolean} */
          nodes[i] = true;
          /**
           * @param {?} first
           * @param {?} expectedHashCode
           * @return {?}
           */
          remove = function(first, expectedHashCode) {
            return assert(first, i, expectedHashCode), expectedHashCode;
          };
          /**
           * @param {?} first
           * @return {?}
           */
          query = function(first) {
            return filter(first, i) ? first[i] : {};
          };
          /**
           * @param {?} tree
           * @return {?}
           */
          walk = function(tree) {
            return filter(tree, i);
          };
        }
        $.exports = {
          /** @type {function (?, ?): ?} */
          set : remove,
          /** @type {function (?): ?} */
          get : query,
          /** @type {function (?): ?} */
          has : walk,
          /**
           * @param {?} nodes
           * @return {?}
           */
          enforce : function(nodes) {
            return walk(nodes) ? query(nodes) : remove(nodes, {});
          },
          /**
           * @param {string} method
           * @return {?}
           */
          getterFor : function(method) {
            return function(nodes) {
              var n;
              if (!getActual(nodes) || (n = query(nodes)).type !== method) {
                throw TypeError("Incompatible receiver, " + method + " required");
              }
              return n;
            };
          }
        };
      }, function(module, dataAndEvents) {
        /** @type {function (this:*): string} */
        var ostring = {}.toString;
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          return ostring.call(expectedNumberOfNonCommentArgs).slice(8, -1);
        };
      }, function(module, dataAndEvents) {
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          if (null == expectedNumberOfNonCommentArgs) {
            throw TypeError("Can't call method on " + expectedNumberOfNonCommentArgs);
          }
          return expectedNumberOfNonCommentArgs;
        };
      }, function(module, dataAndEvents) {
        module.exports = {};
      }, function(mod, dataAndEvents, $sanitize) {
        var __bind = $sanitize(16);
        /**
         * @param {?} promise
         * @return {undefined}
         */
        var Deferred = function(promise) {
          var text;
          var doneResults;
          this.promise = new promise(function(textAlt, data) {
            if (void 0 !== text || void 0 !== doneResults) {
              throw TypeError("Bad Promise constructor");
            }
            text = textAlt;
            doneResults = data;
          });
          this.resolve = __bind(text);
          this.reject = __bind(doneResults);
        };
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        mod.exports.f = function(expectedNumberOfNonCommentArgs) {
          return new Deferred(expectedNumberOfNonCommentArgs);
        };
      }, function(module, dataAndEvents) {
        var dom;
        dom = function() {
          return this;
        }();
        try {
          dom = dom || (new Function("return this"))();
        } catch (t) {
          if ("object" == typeof window) {
            /** @type {Window} */
            dom = window;
          }
        }
        module.exports = dom;
      }, function(module, dataAndEvents) {
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Object} object
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, object) {
          return{
            enumerable : !(1 & expectedNumberOfNonCommentArgs),
            configurable : !(2 & expectedNumberOfNonCommentArgs),
            writable : !(4 & expectedNumberOfNonCommentArgs),
            value : object
          };
        };
      }, function(module, dataAndEvents) {
        /** @type {function (*): number} */
        var ceil = Math.ceil;
        /** @type {function (*): number} */
        var floor = Math.floor;
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          return isNaN(expectedNumberOfNonCommentArgs = +expectedNumberOfNonCommentArgs) ? 0 : (expectedNumberOfNonCommentArgs > 0 ? floor : ceil)(expectedNumberOfNonCommentArgs);
        };
      }, function(module, dataAndEvents, require) {
        var format = require(54);
        var getActual = require(20);
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          return format(getActual(expectedNumberOfNonCommentArgs));
        };
      }, function(module, dataAndEvents, require) {
        var getActual = require(6);
        var objDisplay = require(95);
        var assert = require(38);
        var test = require(43);
        var inspect = require(96);
        var log = require(97);
        /**
         * @param {?} descriptor
         * @param {Object} e
         * @return {undefined}
         */
        var Promise = function(descriptor, e) {
          this.stopped = descriptor;
          /** @type {Object} */
          this.result = e;
        };
        /**
         * @param {string} gotoEnd
         * @return {?}
         */
        (module.exports = function(expectedNumberOfNonCommentArgs, proto, data, err, chai) {
          var it;
          var expected;
          var i;
          var l;
          var obj;
          var ostring;
          var current;
          var callback = test(proto, data, err ? 2 : 1);
          if (chai) {
            /** @type {number} */
            it = expectedNumberOfNonCommentArgs;
          } else {
            if ("function" != typeof(expected = inspect(expectedNumberOfNonCommentArgs))) {
              throw TypeError("Target is not iterable");
            }
            if (objDisplay(expected)) {
              /** @type {number} */
              i = 0;
              l = assert(expectedNumberOfNonCommentArgs.length);
              for (;l > i;i++) {
                if ((obj = err ? callback(getActual(current = expectedNumberOfNonCommentArgs[i])[0], current[1]) : callback(expectedNumberOfNonCommentArgs[i])) && obj instanceof Promise) {
                  return obj;
                }
              }
              return new Promise(false);
            }
            it = expected.call(expectedNumberOfNonCommentArgs);
          }
          ostring = it.next;
          for (;!(current = ostring.call(it)).done;) {
            if ("object" == typeof(obj = log(it, callback, current.value, err)) && (obj && obj instanceof Promise)) {
              return obj;
            }
          }
          return new Promise(false);
        }).stop = function(gotoEnd) {
          return new Promise(true, gotoEnd);
        };
      }, function(module, dataAndEvents) {
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          try {
            return{
              error : false,
              value : expectedNumberOfNonCommentArgs()
            };
          } catch (origValue) {
            return{
              error : true,
              value : origValue
            };
          }
        };
      }, function(module, dataAndEvents) {
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          /** @type {number} */
          var i = 0;
          /** @type {number} */
          var j = expectedNumberOfNonCommentArgs.length - 1;
          for (;i < j;++i, --j) {
            var o = expectedNumberOfNonCommentArgs[j];
            expectedNumberOfNonCommentArgs[j] = expectedNumberOfNonCommentArgs[i];
            expectedNumberOfNonCommentArgs[i] = o;
          }
          return expectedNumberOfNonCommentArgs;
        };
      }, function(module, dataAndEvents, $sanitize) {
        var result = {};
        /** @type {string} */
        result[$sanitize(4)("toStringTag")] = "z";
        /** @type {boolean} */
        module.exports = "[object z]" === String(result);
      }, function(module, dataAndEvents, require) {
        var ctor = require(3);
        var getActual = require(10);
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Function} proto
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, proto) {
          try {
            getActual(ctor, expectedNumberOfNonCommentArgs, proto);
          } catch (r) {
            /** @type {Function} */
            ctor[expectedNumberOfNonCommentArgs] = proto;
          }
          return proto;
        };
      }, function(module, dataAndEvents, require) {
        var collection = require(3);
        var assert = require(13);
        var e = collection.document;
        var s = assert(e) && assert(e.createElement);
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          return s ? e.createElement(expectedNumberOfNonCommentArgs) : {};
        };
      }, function(module, dataAndEvents, require) {
        var mod = require(46);
        var ostring = Function.toString;
        if ("function" != typeof mod.inspectSource) {
          /**
           * @param {number} expectedNumberOfNonCommentArgs
           * @return {?}
           */
          mod.inspectSource = function(expectedNumberOfNonCommentArgs) {
            return ostring.call(expectedNumberOfNonCommentArgs);
          };
        }
        /** @type {function (number): ?} */
        module.exports = mod.inspectSource;
      }, function(module, dataAndEvents, require) {
        var sorter = require(45);
        var getActual = require(49);
        var key = sorter("keys");
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          return key[expectedNumberOfNonCommentArgs] || (key[expectedNumberOfNonCommentArgs] = getActual(expectedNumberOfNonCommentArgs));
        };
      }, function(module, dataAndEvents) {
        module.exports = {};
      }, function(dataAndEvents, entry, format) {
        var f = format(11);
        var query = format(53);
        var buildParams = format(24);
        var msg = format(26);
        var get_mangled = format(48);
        var dataAttr = format(7);
        var cohortString = format(47);
        /** @type {function (Object, string): (ObjectPropertyDescriptor|undefined)} */
        var getOwnPropertyDescriptor = Object.getOwnPropertyDescriptor;
        /** @type {Function} */
        entry.f = f ? getOwnPropertyDescriptor : function(expectedNumberOfNonCommentArgs, name) {
          if (expectedNumberOfNonCommentArgs = msg(expectedNumberOfNonCommentArgs), name = get_mangled(name, true), cohortString) {
            try {
              return getOwnPropertyDescriptor(expectedNumberOfNonCommentArgs, name);
            } catch (t) {
            }
          }
          if (dataAttr(expectedNumberOfNonCommentArgs, name)) {
            return buildParams(!query.f.call(expectedNumberOfNonCommentArgs, name), expectedNumberOfNonCommentArgs[name]);
          }
        };
      }, function(module, dataAndEvents, topic) {
        var out = topic(3);
        module.exports = out;
      }, function(module, dataAndEvents, require) {
        var getActual = require(25);
        /** @type {function (...[*]): number} */
        var nativeMin = Math.min;
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          return expectedNumberOfNonCommentArgs > 0 ? nativeMin(getActual(expectedNumberOfNonCommentArgs), 9007199254740991) : 0;
        };
      }, function(module, dataAndEvents) {
        /** @type {Array} */
        module.exports = ["constructor", "hasOwnProperty", "isPrototypeOf", "propertyIsEnumerable", "toLocaleString", "toString", "valueOf"];
      }, function(mod, dataAndEvents, require) {
        var hasKey = require(7);
        var getActual = require(59);
        var typeOf = require(34);
        var xhr = require(84);
        var type = typeOf("IE_PROTO");
        var objectProto = Object.prototype;
        /** @type {Function} */
        mod.exports = xhr ? Object.getPrototypeOf : function(expectedNumberOfNonCommentArgs) {
          return expectedNumberOfNonCommentArgs = getActual(expectedNumberOfNonCommentArgs), hasKey(expectedNumberOfNonCommentArgs, type) ? expectedNumberOfNonCommentArgs[type] : "function" == typeof expectedNumberOfNonCommentArgs.constructor && expectedNumberOfNonCommentArgs instanceof expectedNumberOfNonCommentArgs.constructor ? expectedNumberOfNonCommentArgs.constructor.prototype : expectedNumberOfNonCommentArgs instanceof Object ? objectProto : null;
        };
      }, function(module, dataAndEvents, require) {
        var next;
        var createObject = require(6);
        var getActual = require(85);
        var methods = require(39);
        var qs = require(35);
        var nodes = require(61);
        var query = require(32);
        var i = require(34)("IE_PROTO");
        /**
         * @return {undefined}
         */
        var ctor = function() {
        };
        /**
         * @param {string} range
         * @return {?}
         */
        var fn = function(range) {
          return "<script>" + range + "\x3c/script>";
        };
        /**
         * @return {?}
         */
        var init = function() {
          try {
            next = document.domain && new ActiveXObject("htmlfile");
          } catch (t) {
          }
          var out;
          var el;
          init = next ? function(out) {
            out.write(fn(""));
            out.close();
            var YObject = out.parentWindow.Object;
            return out = null, YObject;
          }(next) : ((el = query("iframe")).style.display = "none", nodes.appendChild(el), el.src = String("javascript:"), (out = el.contentWindow.document).open(), out.write(fn("document.F=Object")), out.close(), out.F);
          var i = methods.length;
          for (;i--;) {
            delete init.prototype[methods[i]];
          }
          return init();
        };
        /** @type {boolean} */
        qs[i] = true;
        /** @type {function ((Object|null), (Object|null)=): Object} */
        module.exports = Object.create || function(expectedNumberOfNonCommentArgs, proto) {
          var newObject;
          return null !== expectedNumberOfNonCommentArgs ? (ctor.prototype = createObject(expectedNumberOfNonCommentArgs), newObject = new ctor, ctor.prototype = null, newObject[i] = expectedNumberOfNonCommentArgs) : newObject = init(), void 0 === proto ? newObject : getActual(newObject, proto);
        };
      }, function(module, dataAndEvents, require) {
        var setDescriptor = require(12).f;
        var getActual = require(7);
        var rvar = require(4)("toStringTag");
        /**
         * @param {Function} expectedNumberOfNonCommentArgs
         * @param {Object} object
         * @param {boolean} value
         * @return {undefined}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, object, value) {
          if (expectedNumberOfNonCommentArgs) {
            if (!getActual(expectedNumberOfNonCommentArgs = value ? expectedNumberOfNonCommentArgs : expectedNumberOfNonCommentArgs.prototype, rvar)) {
              setDescriptor(expectedNumberOfNonCommentArgs, rvar, {
                configurable : true,
                value : object
              });
            }
          }
        };
      }, function(module, dataAndEvents, require) {
        var getActual = require(16);
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Function} object
         * @param {boolean} value
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, object, value) {
          if (getActual(expectedNumberOfNonCommentArgs), void 0 === object) {
            return expectedNumberOfNonCommentArgs;
          }
          switch(value) {
            case 0:
              return function() {
                return expectedNumberOfNonCommentArgs.call(object);
              };
            case 1:
              return function(mapper) {
                return expectedNumberOfNonCommentArgs.call(object, mapper);
              };
            case 2:
              return function(mapper, graphics) {
                return expectedNumberOfNonCommentArgs.call(object, mapper, graphics);
              };
            case 3:
              return function(mapper, graphics, capture) {
                return expectedNumberOfNonCommentArgs.call(object, mapper, graphics, capture);
              };
          }
          return function() {
            return expectedNumberOfNonCommentArgs.apply(object, arguments);
          };
        };
      }, function(f, dataAndEvents, require) {
        var getName = require(14);
        f.exports = getName("navigator", "userAgent") || "";
      }, function(module, dataAndEvents, require) {
        var binary = require(17);
        var Block = require(46);
        (module.exports = function(expectedNumberOfNonCommentArgs, actual) {
          return Block[expectedNumberOfNonCommentArgs] || (Block[expectedNumberOfNonCommentArgs] = void 0 !== actual ? actual : {});
        })("versions", []).push({
          version : "3.6.3",
          mode : binary ? "pure" : "global",
          copyright : "\u00c2\u00a9 2020 Denis Pushkarev (zloirock.ru)"
        });
      }, function(module, dataAndEvents, require) {
        var hooks = require(3);
        var SelectorMatcher = require(31);
        var JsDiff = hooks["__core-js_shared__"] || SelectorMatcher("__core-js_shared__", {});
        module.exports = JsDiff;
      }, function(module, dataAndEvents, func) {
        var actual = func(11);
        var lambda = func(5);
        var unwrap = func(32);
        /** @type {boolean} */
        module.exports = !actual && !lambda(function() {
          return 7 != Object.defineProperty(unwrap("div"), "a", {
            /**
             * @return {?}
             */
            get : function() {
              return 7;
            }
          }).a;
        });
      }, function(module, dataAndEvents, require) {
        var getActual = require(13);
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {boolean} proto
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, proto) {
          if (!getActual(expectedNumberOfNonCommentArgs)) {
            return expectedNumberOfNonCommentArgs;
          }
          var valueOf;
          var str;
          if (proto && ("function" == typeof(valueOf = expectedNumberOfNonCommentArgs.toString) && !getActual(str = valueOf.call(expectedNumberOfNonCommentArgs)))) {
            return str;
          }
          if ("function" == typeof(valueOf = expectedNumberOfNonCommentArgs.valueOf) && !getActual(str = valueOf.call(expectedNumberOfNonCommentArgs))) {
            return str;
          }
          if (!proto && ("function" == typeof(valueOf = expectedNumberOfNonCommentArgs.toString) && !getActual(str = valueOf.call(expectedNumberOfNonCommentArgs)))) {
            return str;
          }
          throw TypeError("Can't convert object to primitive value");
        };
      }, function(module, dataAndEvents) {
        /** @type {number} */
        var count = 0;
        /** @type {number} */
        var id = Math.random();
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          return "Symbol(" + String(void 0 === expectedNumberOfNonCommentArgs ? "" : expectedNumberOfNonCommentArgs) + ")_" + (++count + id).toString(36);
        };
      }, function(module, dataAndEvents, require) {
        var getActual = require(5);
        /** @type {boolean} */
        module.exports = !!Object.getOwnPropertySymbols && !getActual(function() {
          return!String(Symbol());
        });
      }, function(module, dataAndEvents, require) {
        var Block = require(30);
        var forOwn = require(19);
        var camelKey = require(4)("toStringTag");
        /** @type {boolean} */
        var content = "Arguments" == forOwn(function() {
          return arguments;
        }());
        module.exports = Block ? forOwn : function(expectedNumberOfNonCommentArgs) {
          var object;
          var data;
          var idx;
          return void 0 === expectedNumberOfNonCommentArgs ? "Undefined" : null === expectedNumberOfNonCommentArgs ? "Null" : "string" == typeof(data = function($cookies, key) {
            try {
              return $cookies[key];
            } catch (t) {
            }
          }(object = Object(expectedNumberOfNonCommentArgs), camelKey)) ? data : content ? forOwn(object) : "Object" == (idx = forOwn(object)) && "function" == typeof object.callee ? "Arguments" : idx;
        };
      }, function(module, dataAndEvents, require) {
        var inspect = require(8);
        var flag = require(83);
        var toString = require(40);
        var parse = require(62);
        var compile = require(42);
        var getActual = require(10);
        var debug = require(15);
        var sorter = require(4);
        var nodes = require(17);
        var target = require(21);
        var Block = require(58);
        var ar = Block.IteratorPrototype;
        var retval = Block.BUGGY_SAFARI_ITERATORS;
        var key = sorter("iterator");
        /**
         * @return {?}
         */
        var copy = function() {
          return this;
        };
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Object} object
         * @param {Function} data
         * @param {?} msg
         * @param {string} method
         * @param {(Function|string)} chai
         * @param {?} includeAll
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, object, data, msg, method, chai, includeAll) {
          flag(data, object, msg);
          var str;
          var types;
          var type;
          /**
           * @param {Function} key
           * @return {?}
           */
          var callback = function(key) {
            if (key === method && fn) {
              return fn;
            }
            if (!retval && key in obj) {
              return obj[key];
            }
            switch(key) {
              case "keys":
              ;
              case "values":
              ;
              case "entries":
                return function() {
                  return new data(this, key);
                };
            }
            return function() {
              return new data(this);
            };
          };
          /** @type {string} */
          var name = object + " Iterator";
          /** @type {boolean} */
          var lists = false;
          var obj = expectedNumberOfNonCommentArgs.prototype;
          var values = obj[key] || (obj["@@iterator"] || method && obj[method]);
          var fn = !retval && values || callback(method);
          var conditional = "Array" == object && obj.entries || values;
          if (conditional && (str = toString(conditional.call(new expectedNumberOfNonCommentArgs)), ar !== Object.prototype && (str.next && (nodes || (toString(str) === ar || (parse ? parse(str, ar) : "function" != typeof str[key] && getActual(str, key, copy))), compile(str, name, true, true), nodes && (target[name] = copy)))), "values" == method && (values && ("values" !== values.name && (lists = true, fn = function() {
            return values.call(this);
          }))), nodes && !includeAll || (obj[key] === fn || getActual(obj, key, fn)), target[object] = fn, method) {
            if (types = {
              values : callback("values"),
              keys : chai ? fn : callback("keys"),
              entries : callback("entries")
            }, includeAll) {
              for (type in types) {
                if (!(!retval && (!lists && type in obj))) {
                  debug(obj, type, types[type]);
                }
              }
            } else {
              inspect({
                target : object,
                proto : true,
                forced : retval || lists
              }, types);
            }
          }
          return types;
        };
      }, function(dataAndEvents, entry, deepDataAndEvents) {
        /** @type {function (this:Object, string): boolean} */
        var html = {}.propertyIsEnumerable;
        /** @type {function (Object, string): (ObjectPropertyDescriptor|undefined)} */
        var getOwnPropertyDescriptor = Object.getOwnPropertyDescriptor;
        /** @type {boolean} */
        var isFunction = getOwnPropertyDescriptor && !html.call({
          1 : 2
        }, 1);
        /** @type {Function} */
        entry.f = isFunction ? function(expectedNumberOfNonCommentArgs) {
          /** @type {(ObjectPropertyDescriptor|undefined)} */
          var property = getOwnPropertyDescriptor(this, expectedNumberOfNonCommentArgs);
          return!!property && property.enumerable;
        } : html;
      }, function(module, dataAndEvents, require) {
        var nodes = require(5);
        var getActual = require(19);
        /** @type {function (this:string, *=, number=): Array.<string>} */
        var split = "".split;
        /** @type {Function} */
        module.exports = nodes(function() {
          return!Object("z").propertyIsEnumerable(0);
        }) ? function(expectedNumberOfNonCommentArgs) {
          return "String" == getActual(expectedNumberOfNonCommentArgs) ? split.call(expectedNumberOfNonCommentArgs, "") : Object(expectedNumberOfNonCommentArgs);
        } : Object;
      }, function(module, dataAndEvents, require) {
        var inspect = require(7);
        var getActual = require(26);
        var callback = require(81).indexOf;
        var j = require(35);
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Object} proto
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, proto) {
          var key;
          var actual = getActual(expectedNumberOfNonCommentArgs);
          /** @type {number} */
          var maxScanLen = 0;
          /** @type {Array} */
          var ret = [];
          for (key in actual) {
            if (!inspect(j, key)) {
              if (inspect(actual, key)) {
                ret.push(key);
              }
            }
          }
          for (;proto.length > maxScanLen;) {
            if (inspect(actual, key = proto[maxScanLen++])) {
              if (!~callback(ret, key)) {
                ret.push(key);
              }
            }
          }
          return ret;
        };
      }, function(dataAndEvents, object) {
        object.f = Object.getOwnPropertySymbols;
      }, function(module, dataAndEvents, require) {
        var getName = require(5);
        /** @type {RegExp} */
        var r20 = /#|\.prototype\./;
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Function} value
         * @return {?}
         */
        var parse = function(expectedNumberOfNonCommentArgs, value) {
          var next = args[promote(expectedNumberOfNonCommentArgs)];
          return next == end || next != current && ("function" == typeof value ? getName(value) : !!value);
        };
        /** @type {function (number): ?} */
        var promote = parse.normalize = function(style) {
          return String(style).replace(r20, ".").toLowerCase();
        };
        var args = parse.data = {};
        /** @type {string} */
        var current = parse.NATIVE = "N";
        /** @type {string} */
        var end = parse.POLYFILL = "P";
        /** @type {function (number, Function): ?} */
        module.exports = parse;
      }, function(module, dataAndEvents, require) {
        var text;
        var textAlt;
        var body;
        var tighten = require(40);
        var inspect = require(10);
        var indexOf = require(7);
        var sorter = require(4);
        var Block = require(17);
        var key = sorter("iterator");
        /** @type {boolean} */
        var BUGGY_SAFARI_ITERATORS = false;
        if ([].keys) {
          if ("next" in (body = [].keys())) {
            if ((textAlt = tighten(tighten(body))) !== Object.prototype) {
              text = textAlt;
            }
          } else {
            /** @type {boolean} */
            BUGGY_SAFARI_ITERATORS = true;
          }
        }
        if (null == text) {
          text = {};
        }
        if (!Block) {
          if (!indexOf(text, key)) {
            inspect(text, key, function() {
              return this;
            });
          }
        }
        module.exports = {
          IteratorPrototype : text,
          BUGGY_SAFARI_ITERATORS : BUGGY_SAFARI_ITERATORS
        };
      }, function(module, dataAndEvents, require) {
        var getActual = require(20);
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          return Object(getActual(expectedNumberOfNonCommentArgs));
        };
      }, function(module, dataAndEvents, require) {
        var getActual = require(55);
        var args = require(39);
        /** @type {function (Object): Array.<string>} */
        module.exports = Object.keys || function(expectedNumberOfNonCommentArgs) {
          return getActual(expectedNumberOfNonCommentArgs, args);
        };
      }, function(module, dataAndEvents, require) {
        var factory = require(14);
        module.exports = factory("document", "documentElement");
      }, function(module, dataAndEvents, require) {
        var getActual = require(6);
        var getName = require(86);
        module.exports = Object.setPrototypeOf || ("__proto__" in {} ? function() {
          var set;
          /** @type {boolean} */
          var op = false;
          var xs = {};
          try {
            (set = Object.getOwnPropertyDescriptor(Object.prototype, "__proto__").set).call(xs, []);
            /** @type {boolean} */
            op = xs instanceof Array;
          } catch (t) {
          }
          return function(obj, value) {
            return getActual(obj), getName(value), op ? set.call(obj, value) : obj.__proto__ = value, obj;
          };
        }() : void 0);
      }, function(module, dataAndEvents, require) {
        var global = require(3);
        module.exports = global.Promise;
      }, function(module, dataAndEvents, require) {
        var getActual = require(6);
        var inspect = require(16);
        var prop = require(4)("species");
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Object} proto
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, proto) {
          var a;
          var obj = getActual(expectedNumberOfNonCommentArgs).constructor;
          return void 0 === obj || null == (a = getActual(obj)[prop]) ? proto : inspect(a);
        };
      }, function(c, dataAndEvents, require) {
        var callback;
        var channel;
        var thisObj;
        var global = require(3);
        var core = require(5);
        var getActual = require(19);
        var makeIterator = require(43);
        var xml = require(61);
        var inspect = require(32);
        var Block = require(66);
        var l = global.location;
        var pass = global.setImmediate;
        var clear = global.clearImmediate;
        var process = global.process;
        var MessageChannel = global.MessageChannel;
        var _ = global.Dispatch;
        /** @type {number} */
        var a = 0;
        var data = {};
        /**
         * @param {number} property
         * @return {undefined}
         */
        var resolve = function(property) {
          if (data.hasOwnProperty(property)) {
            var fn = data[property];
            delete data[property];
            fn();
          }
        };
        /**
         * @param {number} value
         * @return {?}
         */
        var expect = function(value) {
          return function() {
            resolve(value);
          };
        };
        /**
         * @param {MessageEvent} e
         * @return {undefined}
         */
        var completed = function(e) {
          resolve(e.data);
        };
        /**
         * @param {string} o
         * @return {undefined}
         */
        var request = function(o) {
          global.postMessage(o + "", l.protocol + "//" + l.host);
        };
        if (!(pass && clear)) {
          /**
           * @param {?} fn
           * @return {?}
           */
          pass = function(fn) {
            /** @type {Array} */
            var ta = [];
            /** @type {number} */
            var i = 1;
            for (;arguments.length > i;) {
              ta.push(arguments[i++]);
            }
            return data[++a] = function() {
              ("function" == typeof fn ? fn : Function(fn)).apply(void 0, ta);
            }, callback(a), a;
          };
          /**
           * @param {?} first
           * @return {undefined}
           */
          clear = function(first) {
            delete data[first];
          };
          if ("process" == getActual(process)) {
            /**
             * @param {number} stream
             * @return {undefined}
             */
            callback = function(stream) {
              process.nextTick(expect(stream));
            };
          } else {
            if (_ && _.now) {
              /**
               * @param {number} result
               * @return {undefined}
               */
              callback = function(result) {
                _.now(expect(result));
              };
            } else {
              if (MessageChannel && !Block) {
                thisObj = (channel = new MessageChannel).port2;
                /** @type {function (MessageEvent): undefined} */
                channel.port1.onmessage = completed;
                callback = makeIterator(thisObj.postMessage, thisObj, 1);
              } else {
                if (!global.addEventListener || ("function" != typeof postMessage || (global.importScripts || core(request)))) {
                  /** @type {function (number): undefined} */
                  callback = "onreadystatechange" in inspect("script") ? function(result) {
                    /**
                     * @return {undefined}
                     */
                    xml.appendChild(inspect("script")).onreadystatechange = function() {
                      xml.removeChild(this);
                      resolve(result);
                    };
                  } : function(func) {
                    setTimeout(expect(func), 0);
                  };
                } else {
                  /** @type {function (string): undefined} */
                  callback = request;
                  global.addEventListener("message", completed, false);
                }
              }
            }
          }
        }
        c.exports = {
          set : pass,
          clear : clear
        };
      }, function(module, dataAndEvents, getName) {
        var name = getName(44);
        /** @type {boolean} */
        module.exports = /(iphone|ipod|ipad).*applewebkit/i.test(name);
      }, function(module, dataAndEvents, require) {
        var inspect = require(6);
        var getActual = require(13);
        var argv = require(22);
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Object} proto
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, proto) {
          if (inspect(expectedNumberOfNonCommentArgs), getActual(proto) && proto.constructor === expectedNumberOfNonCommentArgs) {
            return proto;
          }
          var invokeDfd = argv.f(expectedNumberOfNonCommentArgs);
          return(0, invokeDfd.resolve)(proto), invokeDfd.promise;
        };
      }, function(dataAndEvents, deepDataAndEvents, require) {
        var getActual = require(8);
        var fn = require(16);
        var v = require(22);
        var xhr = require(28);
        var isArray = require(27);
        getActual({
          target : "Promise",
          stat : true
        }, {
          /**
           * @param {?} promises
           * @return {?}
           */
          allSettled : function(promises) {
            var self = this;
            var result = v.f(self);
            var resolve = result.resolve;
            var iterator = result.reject;
            var x = xhr(function() {
              var callback = fn(self.resolve);
              /** @type {Array} */
              var expectedNumberOfNonCommentArgs = [];
              /** @type {number} */
              var rightId = 0;
              /** @type {number} */
              var u = 1;
              isArray(promises, function(node) {
                /** @type {number} */
                var id = rightId++;
                /** @type {boolean} */
                var a = false;
                expectedNumberOfNonCommentArgs.push(void 0);
                u++;
                callback.call(self, node).then(function(x) {
                  if (!a) {
                    /** @type {boolean} */
                    a = true;
                    expectedNumberOfNonCommentArgs[id] = {
                      status : "fulfilled",
                      value : x
                    };
                    if (!--u) {
                      resolve(expectedNumberOfNonCommentArgs);
                    }
                  }
                }, function(err) {
                  if (!a) {
                    /** @type {boolean} */
                    a = true;
                    expectedNumberOfNonCommentArgs[id] = {
                      status : "rejected",
                      reason : err
                    };
                    if (!--u) {
                      resolve(expectedNumberOfNonCommentArgs);
                    }
                  }
                });
              });
              if (!--u) {
                resolve(expectedNumberOfNonCommentArgs);
              }
            });
            return x.error && iterator(x.value), result.promise;
          }
        });
      }, function(module, key, callback) {
        (function(dataAndEvents) {
          var factory;
          var JsDiff;
          if (!(void 0 === (JsDiff = "function" == typeof(factory = function() {
            var single = void 0 !== dataAndEvents ? dataAndEvents : self;
            if (void 0 !== single.TextEncoder && void 0 !== single.TextDecoder) {
              return{
                TextEncoder : single.TextEncoder,
                TextDecoder : single.TextDecoder
              };
            }
            /** @type {Array} */
            var ea = ["utf8", "utf-8", "unicode-1-1-utf-8"];
            return{
              /**
               * @param {string} value
               * @return {undefined}
               */
              TextEncoder : function(value) {
                if (ea.indexOf(value) < 0 && null != value) {
                  throw new RangeError("Invalid encoding type. Only utf-8 is supported");
                }
                /** @type {string} */
                this.encoding = "utf-8";
                /**
                 * @param {number} string
                 * @return {?}
                 */
                this.encode = function(string) {
                  if ("string" != typeof string) {
                    throw new TypeError("passed argument must be of type string");
                  }
                  /** @type {string} */
                  var message = unescape(encodeURIComponent(string));
                  /** @type {Uint8Array} */
                  var buf = new Uint8Array(message.length);
                  return message.split("").forEach(function(a, off) {
                    /** @type {number} */
                    buf[off] = a.charCodeAt(0);
                  }), buf;
                };
              },
              /**
               * @param {string} value
               * @param {Object} options
               * @return {undefined}
               */
              TextDecoder : function(value, options) {
                if (ea.indexOf(value) < 0 && null != value) {
                  throw new RangeError("Invalid encoding type. Only utf-8 is supported");
                }
                if (this.encoding = "utf-8", this.ignoreBOM = false, this.fatal = void 0 !== options && ("fatal" in options && options.fatal), "boolean" != typeof this.fatal) {
                  throw new TypeError("fatal flag must be boolean");
                }
                /**
                 * @param {number} bytes
                 * @param {Object} data
                 * @return {?}
                 */
                this.decode = function(bytes, data) {
                  if (void 0 === bytes) {
                    return "";
                  }
                  if ("boolean" != typeof(void 0 !== data && ("stream" in data && data.stream))) {
                    throw new TypeError("stream option must be boolean");
                  }
                  if (ArrayBuffer.isView(bytes)) {
                    /** @type {Uint8Array} */
                    var dataArray = new Uint8Array(bytes.buffer);
                    /** @type {Array} */
                    var buf = new Array(dataArray.length);
                    return dataArray.forEach(function(c, off) {
                      /** @type {string} */
                      buf[off] = String.fromCharCode(c);
                    }), decodeURIComponent(escape(buf.join("")));
                  }
                  throw new TypeError("passed argument must be an array buffer view");
                };
              }
            };
          }) ? factory.apply(key, []) : factory))) {
            module.exports = JsDiff;
          }
        }).call(this, callback(23));
      }, function(module, dataAndEvents, topic) {
        var out = topic(71);
        topic(103);
        topic(104);
        topic(105);
        topic(106);
        module.exports = out;
      }, function(module, dataAndEvents, require) {
        require(72);
        require(76);
        require(87);
        require(91);
        require(68);
        require(102);
        var global = require(37);
        module.exports = global.Promise;
      }, function(dataAndEvents, deepDataAndEvents, func) {
        var actual = func(30);
        var makeInherit = func(15);
        var newResult = func(75);
        if (!actual) {
          makeInherit(Object.prototype, "toString", newResult, {
            unsafe : true
          });
        }
      }, function(module, dataAndEvents, fun) {
        var exports = fun(50);
        module.exports = exports && (!Symbol.sham && "symbol" == typeof Symbol.iterator);
      }, function(module, dataAndEvents, require) {
        var expect = require(3);
        var next = require(33);
        var name = expect.WeakMap;
        /** @type {boolean} */
        module.exports = "function" == typeof name && /native code/.test(next(name));
      }, function(module, dataAndEvents, require) {
        var Block = require(30);
        var getActual = require(51);
        /** @type {Function} */
        module.exports = Block ? {}.toString : function() {
          return "[object " + getActual(this) + "]";
        };
      }, function(dataAndEvents, deepDataAndEvents, require) {
        var fn = require(77).charAt;
        var style = require(18);
        var is = require(52);
        var setStyle = style.set;
        var toObject = style.getterFor("String Iterator");
        is(String, "String", function(string) {
          setStyle(this, {
            type : "String Iterator",
            string : String(string),
            index : 0
          });
        }, function() {
          var result;
          var self = toObject(this);
          var list = self.string;
          var index = self.index;
          return index >= list.length ? {
            value : void 0,
            done : true
          } : (result = fn(list, index), self.index += result.length, {
            value : result,
            done : false
          });
        });
      }, function(module, dataAndEvents, getCallback) {
        var cb = getCallback(25);
        var callback = getCallback(20);
        /**
         * @param {boolean} i
         * @return {?}
         */
        var write = function(i) {
          return function(value, evt) {
            var el;
            var s;
            /** @type {string} */
            var source = String(callback(value));
            var index = cb(evt);
            /** @type {number} */
            var len = source.length;
            return index < 0 || index >= len ? i ? "" : void 0 : (el = source.charCodeAt(index)) < 55296 || (el > 56319 || (index + 1 === len || ((s = source.charCodeAt(index + 1)) < 56320 || s > 57343))) ? i ? source.charAt(index) : el : i ? source.slice(index, index + 2) : s - 56320 + (el - 55296 << 10) + 65536;
          };
        };
        module.exports = {
          codeAt : write(false),
          charAt : write(true)
        };
      }, function(module, dataAndEvents, require) {
        var inspect = require(7);
        var getActual = require(79);
        var cfg = require(36);
        var a = require(12);
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {number} object
         * @return {undefined}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, object) {
          var codeSegments = getActual(object);
          var f = a.f;
          var callback = cfg.f;
          /** @type {number} */
          var i = 0;
          for (;i < codeSegments.length;i++) {
            var depth = codeSegments[i];
            if (!inspect(expectedNumberOfNonCommentArgs, depth)) {
              f(expectedNumberOfNonCommentArgs, depth, callback(object, depth));
            }
          }
        };
      }, function(module, dataAndEvents, require) {
        var getActual = require(14);
        var ret = require(80);
        var a = require(56);
        var inspect = require(6);
        module.exports = getActual("Reflect", "ownKeys") || function(expectedNumberOfNonCommentArgs) {
          var r = ret.f(inspect(expectedNumberOfNonCommentArgs));
          var f = a.f;
          return f ? r.concat(f(expectedNumberOfNonCommentArgs)) : r;
        };
      }, function(dataAndEvents, entry, toArray) {
        var dataAttr = toArray(55);
        var camelKey = toArray(39).concat("length", "prototype");
        /** @type {function (Object): Array.<string>} */
        entry.f = Object.getOwnPropertyNames || function(expectedNumberOfNonCommentArgs) {
          return dataAttr(expectedNumberOfNonCommentArgs, camelKey);
        };
      }, function(mod, dataAndEvents, require) {
        var getName = require(26);
        var getActual = require(38);
        var check = require(82);
        /**
         * @param {boolean} recurring
         * @return {?}
         */
        var guard = function(recurring) {
          return function(value, searchElement, key) {
            var target;
            var t = getName(value);
            var actual = getActual(t.length);
            var k = check(key, actual);
            if (recurring && searchElement != searchElement) {
              for (;actual > k;) {
                if ((target = t[k++]) != target) {
                  return true;
                }
              }
            } else {
              for (;actual > k;k++) {
                if ((recurring || k in t) && t[k] === searchElement) {
                  return recurring || (k || 0);
                }
              }
            }
            return!recurring && -1;
          };
        };
        mod.exports = {
          includes : guard(true),
          indexOf : guard(false)
        };
      }, function(module, dataAndEvents, require) {
        var getActual = require(25);
        /** @type {function (...[*]): number} */
        var nativeMax = Math.max;
        /** @type {function (...[*]): number} */
        var nativeMin = Math.min;
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Function} proto
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, proto) {
          var fromIndex = getActual(expectedNumberOfNonCommentArgs);
          return fromIndex < 0 ? nativeMax(fromIndex + proto, 0) : nativeMin(fromIndex, proto);
        };
      }, function(module, dataAndEvents, require) {
        var basePrototype = require(58).IteratorPrototype;
        var getActual = require(41);
        var next = require(24);
        var inspect = require(42);
        var nodes = require(21);
        /**
         * @return {?}
         */
        var result = function() {
          return this;
        };
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Function} proto
         * @param {boolean} data
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, proto, data) {
          /** @type {string} */
          var depth = proto + " Iterator";
          return expectedNumberOfNonCommentArgs.prototype = getActual(basePrototype, {
            next : next(1, data)
          }), inspect(expectedNumberOfNonCommentArgs, depth, false, true), nodes[depth] = result, expectedNumberOfNonCommentArgs;
        };
      }, function(module, dataAndEvents, require) {
        var getActual = require(5);
        /** @type {boolean} */
        module.exports = !getActual(function() {
          /**
           * @return {undefined}
           */
          function C() {
          }
          return C.prototype.constructor = null, Object.getPrototypeOf(new C) !== C.prototype;
        });
      }, function(module, dataAndEvents, require) {
        var Block = require(11);
        var callback = require(12);
        var getActual = require(6);
        var forOwn = require(60);
        /** @type {Function} */
        module.exports = Block ? Object.defineProperties : function(expectedNumberOfNonCommentArgs, object) {
          getActual(expectedNumberOfNonCommentArgs);
          var key;
          var keys = forOwn(object);
          var len = keys.length;
          /** @type {number} */
          var j = 0;
          for (;len > j;) {
            callback.f(expectedNumberOfNonCommentArgs, key = keys[j++], object[key]);
          }
          return expectedNumberOfNonCommentArgs;
        };
      }, function(module, dataAndEvents, require) {
        var getActual = require(13);
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          if (!getActual(expectedNumberOfNonCommentArgs) && null !== expectedNumberOfNonCommentArgs) {
            throw TypeError("Can't set " + String(expectedNumberOfNonCommentArgs) + " as a prototype");
          }
          return expectedNumberOfNonCommentArgs;
        };
      }, function(dataAndEvents, deepDataAndEvents, require) {
        var events = require(3);
        var types = require(88);
        var data = require(89);
        var set = require(10);
        var getName = require(4);
        var key = getName("iterator");
        var name = getName("toStringTag");
        var value = data.values;
        var type;
        for (type in types) {
          var constructor = events[type];
          var cache = constructor && constructor.prototype;
          if (cache) {
            if (cache[key] !== value) {
              try {
                set(cache, key, value);
              } catch (t) {
                cache[key] = value;
              }
            }
            if (cache[name] || set(cache, name, type), types[type]) {
              var prop;
              for (prop in data) {
                if (cache[prop] !== data[prop]) {
                  try {
                    set(cache, prop, data[prop]);
                  } catch (t) {
                    cache[prop] = data[prop];
                  }
                }
              }
            }
          }
        }
      }, function(module, dataAndEvents) {
        module.exports = {
          CSSRuleList : 0,
          CSSStyleDeclaration : 0,
          CSSValueList : 0,
          ClientRectList : 0,
          DOMRectList : 0,
          DOMStringList : 0,
          DOMTokenList : 1,
          DataTransferItemList : 0,
          FileList : 0,
          HTMLAllCollection : 0,
          HTMLCollection : 0,
          HTMLFormElement : 0,
          HTMLSelectElement : 0,
          MediaList : 0,
          MimeTypeArray : 0,
          NamedNodeMap : 0,
          NodeList : 1,
          PaintRequestList : 0,
          Plugin : 0,
          PluginArray : 0,
          SVGLengthList : 0,
          SVGNumberList : 0,
          SVGPathSegList : 0,
          SVGPointList : 0,
          SVGStringList : 0,
          SVGTransformList : 0,
          SourceBufferList : 0,
          StyleSheetList : 0,
          TextTrackCueList : 0,
          TextTrackList : 0,
          TouchList : 0
        };
      }, function(module, dataAndEvents, require) {
        var getActual = require(26);
        var isArray = require(90);
        var nodes = require(21);
        var style = require(18);
        var factory = require(52);
        var setStyle = style.set;
        var ease = style.getterFor("Array Iterator");
        module.exports = factory(Array, "Array", function(obj, kind) {
          setStyle(this, {
            type : "Array Iterator",
            target : getActual(obj),
            index : 0,
            kind : kind
          });
        }, function() {
          var e = ease(this);
          var source = e.target;
          var key = e.kind;
          /** @type {number} */
          var i = e.index++;
          return!source || i >= source.length ? (e.target = void 0, {
            value : void 0,
            done : true
          }) : "keys" == key ? {
            value : i,
            done : false
          } : "values" == key ? {
            value : source[i],
            done : false
          } : {
            value : [i, source[i]],
            done : false
          };
        }, "values");
        nodes.Arguments = nodes.Array;
        isArray("keys");
        isArray("values");
        isArray("entries");
      }, function(module, dataAndEvents, _) {
        var wrapped = _(4);
        var today = _(41);
        var _this = _(12);
        var optgroup = wrapped("unscopables");
        var x = Array.prototype;
        if (null == x[optgroup]) {
          _this.f(x, optgroup, {
            configurable : true,
            value : today(null)
          });
        }
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {undefined}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          /** @type {boolean} */
          x[optgroup][expectedNumberOfNonCommentArgs] = true;
        };
      }, function(dataAndEvents, deepDataAndEvents, require) {
        var cb;
        var poll;
        var ret;
        var then;
        var nodes = require(8);
        var Block = require(17);
        var self = require(3);
        var factory = require(14);
        var object = require(63);
        var typeOf = require(15);
        var createObject = require(92);
        var render = require(42);
        var visitor = require(93);
        var isKind = require(13);
        var resolve = require(16);
        var group = require(94);
        var getActual = require(19);
        var getName = require(33);
        var cast = require(27);
        var helper = require(98);
        var indexOf = require(64);
        var setter = require(65).set;
        var assert = require(99);
        var callback = require(67);
        var mixIn = require(100);
        var cfg = require(22);
        var call = require(28);
        var Ember = require(18);
        var swap = require(57);
        var flag = require(4);
        var inspect = require(101);
        var obj = flag("species");
        /** @type {string} */
        var view = "Promise";
        var get = Ember.get;
        var setStyle = Ember.set;
        var toObject = Ember.getterFor(view);
        var value = object;
        var a = self.TypeError;
        var doc = self.document;
        var process = self.process;
        var handler = factory("fetch");
        var fn = cfg.f;
        var org = fn;
        /** @type {boolean} */
        var domain = "process" == getActual(process);
        /** @type {boolean} */
        var K = !!(doc && (doc.createEvent && self.dispatchEvent));
        var target = swap(view, function() {
          if (getName(value) === String(value)) {
            if (66 === inspect) {
              return true;
            }
            if (!domain && "function" != typeof PromiseRejectionEvent) {
              return true;
            }
          }
          if (Block && !value.prototype.finally) {
            return true;
          }
          if (inspect >= 51 && /native code/.test(value)) {
            return false;
          }
          var me = value.resolve(1);
          /**
           * @param {?} onComplete
           * @return {undefined}
           */
          var finish = function(onComplete) {
            onComplete(function() {
            }, function() {
            });
          };
          return(me.constructor = {})[obj] = finish, !(me.then(function() {
          }) instanceof finish);
        });
        var result = target || !helper(function(isXML) {
          value.all(isXML).catch(function() {
          });
        });
        /**
         * @param {?} val
         * @return {?}
         */
        var isArray = function(val) {
          var then;
          return!(!isKind(val) || "function" != typeof(then = val.then)) && then;
        };
        /**
         * @param {Function} e
         * @param {Object} data
         * @param {boolean} recurring
         * @return {undefined}
         */
        var check = function(e, data, recurring) {
          if (!data.notified) {
            /** @type {boolean} */
            data.notified = true;
            var items = data.reactions;
            assert(function() {
              var raw = data.value;
              /** @type {boolean} */
              var caseSensitive = 1 == data.state;
              /** @type {number} */
              var index = 0;
              for (;items.length > index;) {
                var expectedNumberOfNonCommentArgs;
                var cb;
                var c;
                var result = items[index++];
                var text = caseSensitive ? result.ok : result.fail;
                var resolve = result.resolve;
                var callback = result.reject;
                var domain = result.domain;
                try {
                  if (text) {
                    if (!caseSensitive) {
                      if (2 === data.rejection) {
                        finish(e, data);
                      }
                      /** @type {number} */
                      data.rejection = 1;
                    }
                    if (true === text) {
                      expectedNumberOfNonCommentArgs = raw;
                    } else {
                      if (domain) {
                        domain.enter();
                      }
                      expectedNumberOfNonCommentArgs = text(raw);
                      if (domain) {
                        domain.exit();
                        /** @type {boolean} */
                        c = true;
                      }
                    }
                    if (expectedNumberOfNonCommentArgs === result.promise) {
                      callback(a("Promise-chain cycle"));
                    } else {
                      if (cb = isArray(expectedNumberOfNonCommentArgs)) {
                        cb.call(expectedNumberOfNonCommentArgs, resolve, callback);
                      } else {
                        resolve(expectedNumberOfNonCommentArgs);
                      }
                    }
                  } else {
                    callback(raw);
                  }
                } catch (STOP) {
                  if (domain) {
                    if (!c) {
                      domain.exit();
                    }
                  }
                  callback(STOP);
                }
              }
              /** @type {Array} */
              data.reactions = [];
              /** @type {boolean} */
              data.notified = false;
              if (recurring) {
                if (!data.rejection) {
                  done(e, data);
                }
              }
            });
          }
        };
        /**
         * @param {string} type
         * @param {Object} target
         * @param {?} options
         * @return {undefined}
         */
        var triggerEvent = function(type, target, options) {
          var event;
          var getXYfromEvent;
          if (K) {
            /** @type {Object} */
            (event = doc.createEvent("Event")).promise = target;
            event.reason = options;
            event.initEvent(type, false, true);
            self.dispatchEvent(event);
          } else {
            event = {
              promise : target,
              reason : options
            };
          }
          if (getXYfromEvent = self["on" + type]) {
            getXYfromEvent(event);
          } else {
            if ("unhandledrejection" === type) {
              mixIn("Unhandled promise rejection", options);
            }
          }
        };
        /**
         * @param {Function} data
         * @param {Object} u
         * @return {undefined}
         */
        var done = function(data, u) {
          setter.call(self, function() {
            var object;
            var args = u.value;
            if (map(u) && (object = call(function() {
              if (domain) {
                process.emit("unhandledRejection", args, data);
              } else {
                triggerEvent("unhandledrejection", data, args);
              }
            }), u.rejection = domain || map(u) ? 2 : 1, object.error)) {
              throw object.value;
            }
          });
        };
        /**
         * @param {Object} v
         * @return {?}
         */
        var map = function(v) {
          return 1 !== v.rejection && !v.parent;
        };
        /**
         * @param {Function} err
         * @param {Object} buffer
         * @return {undefined}
         */
        var finish = function(err, buffer) {
          setter.call(self, function() {
            if (domain) {
              process.emit("rejectionHandled", err);
            } else {
              triggerEvent("rejectionhandled", err, buffer.value);
            }
          });
        };
        /**
         * @param {Function} callback
         * @param {Function} value
         * @param {?} o
         * @param {Object} context
         * @return {?}
         */
        var $ = function(callback, value, o, context) {
          return function(arg) {
            callback(value, o, arg, context);
          };
        };
        /**
         * @param {Function} event
         * @param {Object} d
         * @param {Function} f
         * @param {Object} i
         * @return {undefined}
         */
        var next = function(event, d, f, i) {
          if (!d.done) {
            /** @type {boolean} */
            d.done = true;
            if (i) {
              /** @type {Object} */
              d = i;
            }
            /** @type {Function} */
            d.value = f;
            /** @type {number} */
            d.state = 2;
            check(event, d, true);
          }
        };
        /**
         * @param {Function} name
         * @param {Object} callback
         * @param {Function} t
         * @param {Object} expected
         * @return {undefined}
         */
        var test = function(name, callback, t, expected) {
          if (!callback.done) {
            /** @type {boolean} */
            callback.done = true;
            if (expected) {
              /** @type {Object} */
              callback = expected;
            }
            try {
              if (name === t) {
                throw a("Promise can't be resolved itself");
              }
              var self = isArray(t);
              if (self) {
                assert(function() {
                  var res = {
                    done : false
                  };
                  try {
                    self.call(t, $(test, name, res, callback), $(next, name, res, callback));
                  } catch (fromIndex) {
                    next(name, res, fromIndex, callback);
                  }
                });
              } else {
                /** @type {Function} */
                callback.value = t;
                /** @type {number} */
                callback.state = 1;
                check(name, callback, false);
              }
            } catch (fromIndex) {
              next(name, {
                done : false
              }, fromIndex, callback);
            }
          }
        };
        if (target) {
          /**
           * @param {number} expectedNumberOfNonCommentArgs
           * @return {undefined}
           */
          value = function(expectedNumberOfNonCommentArgs) {
            group(this, value, view);
            resolve(expectedNumberOfNonCommentArgs);
            cb.call(this);
            var res = get(this);
            try {
              expectedNumberOfNonCommentArgs($(test, this, res), $(next, this, res));
            } catch (fromIndex) {
              next(this, res, fromIndex);
            }
          };
          (cb = function(stats) {
            setStyle(this, {
              type : view,
              done : false,
              notified : false,
              parent : false,
              reactions : [],
              rejection : false,
              state : 0,
              value : void 0
            });
          }).prototype = createObject(value.prototype, {
            /**
             * @param {Object} opt_attributes
             * @param {Object} onReject
             * @return {?}
             */
            then : function(opt_attributes, onReject) {
              var self = toObject(this);
              var response = fn(indexOf(this, value));
              return response.ok = "function" != typeof opt_attributes || opt_attributes, response.fail = "function" == typeof onReject && onReject, response.domain = domain ? process.domain : void 0, self.parent = true, self.reactions.push(response), 0 != self.state && check(this, self, false), response.promise;
            },
            /**
             * @param {Object} callback
             * @return {?}
             */
            catch : function(callback) {
              return this.then(void 0, callback);
            }
          });
          /**
           * @return {undefined}
           */
          poll = function() {
            var results = new cb;
            var res = get(results);
            this.promise = results;
            this.resolve = $(test, results, res);
            this.reject = $(next, results, res);
          };
          /** @type {function (number): ?} */
          cfg.f = fn = function(expectedNumberOfNonCommentArgs) {
            return expectedNumberOfNonCommentArgs === value || expectedNumberOfNonCommentArgs === ret ? new poll(expectedNumberOfNonCommentArgs) : org(expectedNumberOfNonCommentArgs);
          };
          if (!Block) {
            if (!("function" != typeof object)) {
              then = object.prototype.then;
              typeOf(object.prototype, "then", function(attributes, onReject) {
                var self = this;
                return(new value(function(mapper, reject) {
                  then.call(self, mapper, reject);
                })).then(attributes, onReject);
              }, {
                unsafe : true
              });
              if ("function" == typeof handler) {
                nodes({
                  global : true,
                  enumerable : true,
                  forced : true
                }, {
                  /**
                   * @param {?} pool
                   * @return {?}
                   */
                  fetch : function(pool) {
                    return callback(value, handler.apply(self, arguments));
                  }
                });
              }
            }
          }
        }
        nodes({
          global : true,
          wrap : true,
          forced : target
        }, {
          Promise : value
        });
        render(value, view, false, true);
        visitor(view);
        ret = factory(view);
        nodes({
          target : view,
          stat : true,
          forced : target
        }, {
          /**
           * @param {number} expectedNumberOfNonCommentArgs
           * @return {?}
           */
          reject : function(expectedNumberOfNonCommentArgs) {
            var callback = fn(this);
            return callback.reject.call(void 0, expectedNumberOfNonCommentArgs), callback.promise;
          }
        });
        nodes({
          target : view,
          stat : true,
          forced : Block || target
        }, {
          /**
           * @param {number} expectedNumberOfNonCommentArgs
           * @return {?}
           */
          resolve : function(expectedNumberOfNonCommentArgs) {
            return callback(Block && this === ret ? value : this, expectedNumberOfNonCommentArgs);
          }
        });
        nodes({
          target : view,
          stat : true,
          forced : result
        }, {
          /**
           * @param {?} value
           * @return {?}
           */
          all : function(value) {
            var view = this;
            var result = fn(view);
            var resolver = result.resolve;
            var reject = result.reject;
            var r = call(function() {
              var fn = resolve(view.resolve);
              /** @type {Array} */
              var expectedNumberOfNonCommentArgs = [];
              /** @type {number} */
              var s = 0;
              /** @type {number} */
              var a = 1;
              cast(value, function(locals) {
                /** @type {number} */
                var unlock = s++;
                /** @type {boolean} */
                var c = false;
                expectedNumberOfNonCommentArgs.push(void 0);
                a++;
                fn.call(view, locals).then(function(data) {
                  if (!c) {
                    /** @type {boolean} */
                    c = true;
                    expectedNumberOfNonCommentArgs[unlock] = data;
                    if (!--a) {
                      resolver(expectedNumberOfNonCommentArgs);
                    }
                  }
                }, reject);
              });
              if (!--a) {
                resolver(expectedNumberOfNonCommentArgs);
              }
            });
            return r.error && reject(r.value), result.promise;
          },
          /**
           * @param {?} array
           * @return {?}
           */
          race : function(array) {
            var expectedNumberOfNonCommentArgs = this;
            var result = fn(expectedNumberOfNonCommentArgs);
            var reject = result.reject;
            var r = call(function() {
              var fn = resolve(expectedNumberOfNonCommentArgs.resolve);
              cast(array, function(locals) {
                fn.call(expectedNumberOfNonCommentArgs, locals).then(result.resolve, reject);
              });
            });
            return r.error && reject(r.value), result.promise;
          }
        });
      }, function(module, dataAndEvents, $sanitize) {
        var dataAttr = $sanitize(15);
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Object} object
         * @param {boolean} data
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, object, data) {
          var name;
          for (name in object) {
            dataAttr(expectedNumberOfNonCommentArgs, name, object[name], data);
          }
          return expectedNumberOfNonCommentArgs;
        };
      }, function(module, dataAndEvents, require) {
        var inspect = require(14);
        var argv = require(12);
        var getActual = require(4);
        var Block = require(11);
        var optgroup = getActual("species");
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {undefined}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          var str = inspect(expectedNumberOfNonCommentArgs);
          var len = argv.f;
          if (Block) {
            if (str) {
              if (!str[optgroup]) {
                len(str, optgroup, {
                  configurable : true,
                  /**
                   * @return {?}
                   */
                  get : function() {
                    return this;
                  }
                });
              }
            }
          }
        };
      }, function(module, dataAndEvents) {
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Function} proto
         * @param {string} value
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, proto, value) {
          if (!(expectedNumberOfNonCommentArgs instanceof proto)) {
            throw TypeError("Incorrect " + (value ? value + " " : "") + "invocation");
          }
          return expectedNumberOfNonCommentArgs;
        };
      }, function(module, dataAndEvents, require) {
        var getName = require(4);
        var nodes = require(21);
        var name = getName("iterator");
        var ap = Array.prototype;
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          return void 0 !== expectedNumberOfNonCommentArgs && (nodes.Array === expectedNumberOfNonCommentArgs || ap[name] === expectedNumberOfNonCommentArgs);
        };
      }, function(module, dataAndEvents, require) {
        var getActual = require(51);
        var Block = require(21);
        var expression = require(4)("iterator");
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs) {
          if (null != expectedNumberOfNonCommentArgs) {
            return expectedNumberOfNonCommentArgs[expression] || (expectedNumberOfNonCommentArgs["@@iterator"] || Block[getActual(expectedNumberOfNonCommentArgs)]);
          }
        };
      }, function(module, dataAndEvents, require) {
        var inspect = require(6);
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Function} callback
         * @param {Object} data
         * @param {?} err
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, callback, data, err) {
          try {
            return err ? callback(inspect(data)[0], data[1]) : callback(data);
          } catch (e) {
            var closure = expectedNumberOfNonCommentArgs.return;
            throw void 0 !== closure && inspect(closure.call(expectedNumberOfNonCommentArgs)), e;
          }
        };
      }, function(module, dataAndEvents, $sanitize) {
        var prop = $sanitize(4)("iterator");
        /** @type {boolean} */
        var property = false;
        try {
          /** @type {number} */
          var o = 0;
          var args = {
            /**
             * @return {?}
             */
            next : function() {
              return{
                done : !!o++
              };
            },
            /**
             * @return {undefined}
             */
            return : function() {
              /** @type {boolean} */
              property = true;
            }
          };
          /**
           * @return {?}
           */
          args[prop] = function() {
            return this;
          };
          Array.from(args, function() {
            throw 2;
          });
        } catch (t) {
        }
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Function} object
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, object) {
          if (!object && !property) {
            return false;
          }
          /** @type {boolean} */
          var str = false;
          try {
            var originalEvent = {};
            /**
             * @return {?}
             */
            originalEvent[prop] = function() {
              return{
                /**
                 * @return {?}
                 */
                next : function() {
                  return{
                    done : str = true
                  };
                }
              };
            };
            expectedNumberOfNonCommentArgs(originalEvent);
          } catch (t) {
          }
          return str;
        };
      }, function(module, dataAndEvents, require) {
        var flush;
        var current;
        var ret;
        var send;
        var iterations;
        var node;
        var value;
        var then;
        var expectedNumberOfNonCommentArgs = require(3);
        var requestAnimationFrame = require(36).f;
        var getActual = require(19);
        var setter = require(65).set;
        var Block = require(66);
        var BrowserMutationObserver = expectedNumberOfNonCommentArgs.MutationObserver || expectedNumberOfNonCommentArgs.WebKitMutationObserver;
        var process = expectedNumberOfNonCommentArgs.process;
        var p1 = expectedNumberOfNonCommentArgs.Promise;
        /** @type {boolean} */
        var b = "process" == getActual(process);
        var id = requestAnimationFrame(expectedNumberOfNonCommentArgs, "queueMicrotask");
        var exports = id && id.value;
        if (!exports) {
          /**
           * @return {undefined}
           */
          flush = function() {
            var d;
            var keys;
            if (b) {
              if (d = process.domain) {
                d.exit();
              }
            }
            for (;current;) {
              keys = current.fn;
              current = current.next;
              try {
                keys();
              } catch (t) {
                throw current ? send() : ret = void 0, t;
              }
            }
            ret = void 0;
            if (d) {
              d.enter();
            }
          };
          if (b) {
            /**
             * @return {undefined}
             */
            send = function() {
              process.nextTick(flush);
            };
          } else {
            if (BrowserMutationObserver && !Block) {
              /** @type {boolean} */
              iterations = true;
              /** @type {Text} */
              node = document.createTextNode("");
              (new BrowserMutationObserver(flush)).observe(node, {
                characterData : true
              });
              /**
               * @return {undefined}
               */
              send = function() {
                /** @type {boolean} */
                node.data = iterations = !iterations;
              };
            } else {
              if (p1 && p1.resolve) {
                value = p1.resolve(void 0);
                then = value.then;
                /**
                 * @return {undefined}
                 */
                send = function() {
                  then.call(value, flush);
                };
              } else {
                /**
                 * @return {undefined}
                 */
                send = function() {
                  setter.call(expectedNumberOfNonCommentArgs, flush);
                };
              }
            }
          }
        }
        module.exports = exports || function(expectedNumberOfNonCommentArgs) {
          var next = {
            fn : expectedNumberOfNonCommentArgs,
            next : void 0
          };
          if (ret) {
            ret.next = next;
          }
          if (!current) {
            current = next;
            send();
          }
          ret = next;
        };
      }, function(module, dataAndEvents, Event) {
        var self = Event(3);
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Function} actual
         * @return {undefined}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, actual) {
          var test = self.console;
          if (test) {
            if (test.error) {
              if (1 === arguments.length) {
                test.error(expectedNumberOfNonCommentArgs);
              } else {
                test.error(expectedNumberOfNonCommentArgs, actual);
              }
            }
          }
        };
      }, function(module, dataAndEvents, require) {
        var groupedSelectors;
        var selector;
        var global = require(3);
        var expect = require(44);
        var doc = global.process;
        var docElement = doc && doc.versions;
        var uHostName = docElement && docElement.v8;
        if (uHostName) {
          selector = (groupedSelectors = uHostName.split("."))[0] + groupedSelectors[1];
        } else {
          if (expect) {
            if (!(groupedSelectors = expect.match(/Edge\/(\d+)/)) || groupedSelectors[1] >= 74) {
              if (groupedSelectors = expect.match(/Chrome\/(\d+)/)) {
                selector = groupedSelectors[1];
              }
            }
          }
        }
        module.exports = selector && +selector;
      }, function(dataAndEvents, deepDataAndEvents, require) {
        var getActual = require(8);
        var Block = require(17);
        var b = require(63);
        var nodes = require(5);
        var helper = require(14);
        var inspect = require(64);
        var expect = require(67);
        var createObject = require(15);
        getActual({
          target : "Promise",
          proto : true,
          real : true,
          forced : !!b && nodes(function() {
            b.prototype.finally.call({
              /**
               * @return {undefined}
               */
              then : function() {
              }
            }, function() {
            });
          })
        }, {
          /**
           * @param {Function} callback
           * @return {?}
           */
          finally : function(callback) {
            var str = inspect(this, helper("Promise"));
            /** @type {boolean} */
            var fn = "function" == typeof callback;
            return this.then(fn ? function(dataAndEvents) {
              return expect(str, callback()).then(function() {
                return dataAndEvents;
              });
            } : callback, fn ? function(dataAndEvents) {
              return expect(str, callback()).then(function() {
                throw dataAndEvents;
              });
            } : callback);
          }
        });
        if (!Block) {
          if (!("function" != typeof b)) {
            if (!b.prototype.finally) {
              createObject(b.prototype, "finally", helper("Promise").prototype.finally);
            }
          }
        }
      }, function(dataAndEvents, deepDataAndEvents, require) {
        var getActual = require(8);
        var Block = require(11);
        var tmpl = require(40);
        var fn = require(62);
        var create = require(41);
        var module = require(12);
        var helper = require(24);
        var Event = require(27);
        var expect = require(10);
        var request = require(18);
        var callback = request.set;
        var transferFlags = request.getterFor("AggregateError");
        /**
         * @param {string} type
         * @param {string} message
         * @return {?}
         */
        var error = function(type, message) {
          var nodes = this;
          if (!(nodes instanceof error)) {
            return new error(type, message);
          }
          if (fn) {
            nodes = fn(new Error(message), tmpl(nodes));
          }
          /** @type {Array} */
          var errors = [];
          return Event(type, errors.push, errors), Block ? callback(nodes, {
            errors : errors,
            type : "AggregateError"
          }) : nodes.errors = errors, void 0 !== message && expect(nodes, "message", String(message)), nodes;
        };
        error.prototype = create(Error.prototype, {
          constructor : helper(5, error),
          message : helper(5, ""),
          name : helper(5, "AggregateError")
        });
        if (Block) {
          module.f(error.prototype, "errors", {
            /**
             * @return {?}
             */
            get : function() {
              return transferFlags(this).errors;
            },
            configurable : true
          });
        }
        getActual({
          global : true
        }, {
          /** @type {function (string, string): ?} */
          AggregateError : error
        });
      }, function(dataAndEvents, deepDataAndEvents, $sanitize) {
        $sanitize(68);
      }, function(dataAndEvents, deepDataAndEvents, topic) {
        var throttledUpdate = topic(8);
        var out = topic(22);
        var MAP = topic(28);
        throttledUpdate({
          target : "Promise",
          stat : true
        }, {
          /**
           * @param {?} t
           * @return {?}
           */
          try : function(t) {
            var response = out.f(this);
            var g = MAP(t);
            return(g.error ? response.reject : response.resolve)(g.value), response.promise;
          }
        });
      }, function(dataAndEvents, deepDataAndEvents, require) {
        var nodes = require(8);
        var qMap = require(16);
        var helper = require(14);
        var argv = require(22);
        var f = require(28);
        var getActual = require(27);
        nodes({
          target : "Promise",
          stat : true
        }, {
          /**
           * @param {?} obj
           * @return {?}
           */
          any : function(obj) {
            var expectedNumberOfNonCommentArgs = this;
            var original = argv.f(expectedNumberOfNonCommentArgs);
            var button = original.resolve;
            var iterator = original.reject;
            var x = f(function() {
              var native_method = qMap(expectedNumberOfNonCommentArgs.resolve);
              /** @type {Array} */
              var tempData = [];
              /** @type {number} */
              var R = 0;
              /** @type {number} */
              var f = 1;
              /** @type {boolean} */
              var h = false;
              getActual(obj, function(mapper) {
                /** @type {number} */
                var unlock = R++;
                /** @type {boolean} */
                var u = false;
                tempData.push(void 0);
                f++;
                native_method.call(expectedNumberOfNonCommentArgs, mapper).then(function(expectedNumberOfNonCommentArgs) {
                  if (!u) {
                    if (!h) {
                      /** @type {boolean} */
                      h = true;
                      button(expectedNumberOfNonCommentArgs);
                    }
                  }
                }, function(data) {
                  if (!u) {
                    if (!h) {
                      /** @type {boolean} */
                      u = true;
                      tempData[unlock] = data;
                      if (!--f) {
                        iterator(new (helper("AggregateError"))(tempData, "No one promise resolved"));
                      }
                    }
                  }
                });
              });
              if (!--f) {
                iterator(new (helper("AggregateError"))(tempData, "No one promise resolved"));
              }
            });
            return x.error && iterator(x.value), original.promise;
          }
        });
      }, function(module, dataAndEvents, topic) {
        var out = topic(108);
        module.exports = out;
      }, function(module, dataAndEvents, require) {
        require(109);
        var factory = require(113);
        module.exports = factory("String", "padStart");
      }, function(dataAndEvents, deepDataAndEvents, require) {
        var getActual = require(8);
        var oldStart = require(110).start;
        getActual({
          target : "String",
          proto : true,
          forced : require(112)
        }, {
          /**
           * @param {number} curr
           * @return {?}
           */
          padStart : function(curr) {
            return oldStart(this, curr, arguments.length > 1 ? arguments[1] : void 0);
          }
        });
      }, function(module, dataAndEvents, getCallback) {
        var cb = getCallback(38);
        var ostring = getCallback(111);
        var callback = getCallback(20);
        /** @type {function (*): number} */
        var ceil = Math.ceil;
        /**
         * @param {boolean} recurring
         * @return {?}
         */
        var write = function(recurring) {
          return function(value, outErr, string) {
            var width;
            var r;
            /** @type {string} */
            var t = String(callback(value));
            /** @type {number} */
            var left = t.length;
            /** @type {string} */
            var it = void 0 === string ? " " : String(string);
            var right = cb(outErr);
            return right <= left || "" == it ? t : (width = right - left, (r = ostring.call(it, ceil(width / it.length))).length > width && (r = r.slice(0, width)), recurring ? t + r : r + t);
          };
        };
        module.exports = {
          start : write(false),
          end : write(true)
        };
      }, function(module, dataAndEvents, require) {
        var getActual = require(25);
        var nodes = require(20);
        module.exports = "".repeat || function(expectedNumberOfNonCommentArgs) {
          /** @type {string} */
          var s = String(nodes(this));
          /** @type {string} */
          var slarge = "";
          var actual = getActual(expectedNumberOfNonCommentArgs);
          if (actual < 0 || actual == 1 / 0) {
            throw RangeError("Wrong number of repetitions");
          }
          for (;actual > 0;(actual >>>= 1) && (s += s)) {
            if (1 & actual) {
              slarge += s;
            }
          }
          return slarge;
        };
      }, function(module, dataAndEvents, getName) {
        var name = getName(44);
        /** @type {boolean} */
        module.exports = /Version\/10\.\d+(\.\d+)?( Mobile\/\w+)? Safari\//.test(name);
      }, function(module, dataAndEvents, keys) {
        var props = keys(3);
        var ondata = keys(43);
        var call = Function.call;
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Function} proto
         * @param {boolean} data
         * @return {?}
         */
        module.exports = function(expectedNumberOfNonCommentArgs, proto, data) {
          return ondata(call, props[expectedNumberOfNonCommentArgs].prototype[proto], data);
        };
      }, function(module, dataAndEvents, topic) {
        var out = topic(115);
        module.exports = out;
      }, function(module, dataAndEvents, expression) {
        expression(116);
        var obj = expression(37);
        module.exports = obj.Object.assign;
      }, function(dataAndEvents, deepDataAndEvents, require) {
        var getActual = require(8);
        var Block = require(117);
        getActual({
          target : "Object",
          stat : true,
          forced : Object.assign !== Block
        }, {
          assign : Block
        });
      }, function(module, dataAndEvents, require) {
        var Block = require(11);
        var getActual = require(5);
        var expect = require(60);
        var a = require(56);
        var action = require(53);
        var getObject = require(59);
        var inspect = require(54);
        var match = Object.assign;
        /** @type {function (Object, string, Object): Object} */
        var defineProperty = Object.defineProperty;
        module.exports = !match || getActual(function() {
          if (Block && 1 !== match({
            b : 1
          }, match(defineProperty({}, "a", {
            enumerable : true,
            /**
             * @return {undefined}
             */
            get : function() {
              defineProperty(this, "b", {
                value : 3,
                enumerable : false
              });
            }
          }), {
            b : 2
          })).b) {
            return true;
          }
          var prop = {};
          var testName = {};
          var p = Symbol();
          return prop[p] = 7, "abcdefghijklmnopqrst".split("").forEach(function(testname) {
            /** @type {string} */
            testName[testname] = testname;
          }), 7 != match({}, prop)[p] || "abcdefghijklmnopqrst" != expect(match({}, testName)).join("");
        }) ? function(expectedNumberOfNonCommentArgs, object) {
          var cache = getObject(expectedNumberOfNonCommentArgs);
          /** @type {number} */
          var argLength = arguments.length;
          /** @type {number} */
          var current = 1;
          var f = a.f;
          var handler = action.f;
          for (;argLength > current;) {
            var prop;
            var element = inspect(arguments[current++]);
            var second = f ? expect(element).concat(f(element)) : expect(element);
            var l = second.length;
            /** @type {number} */
            var i = 0;
            for (;l > i;) {
              prop = second[i++];
              if (!(Block && !handler.call(element, prop))) {
                cache[prop] = element[prop];
              }
            }
          }
          return cache;
        } : match;
      }, function(module, dataAndEvents, deepDataAndEvents) {
        var html5 = function(self) {
          /**
           * @param {Function} object
           * @param {Function} recurring
           * @param {string} name
           * @param {Array} data
           * @return {?}
           */
          function cb(object, recurring, name, data) {
            var Type = recurring && recurring.prototype instanceof superClass ? recurring : superClass;
            /** @type {Object} */
            var self = Object.create(Type.prototype);
            var pdataCur = new Class(data || []);
            return self._invoke = function(config, path, data) {
              /** @type {string} */
              var arLength = "suspendedStart";
              return function(type, key) {
                if ("executing" === arLength) {
                  throw new Error("Generator is already running");
                }
                if ("completed" === arLength) {
                  if ("throw" === type) {
                    throw key;
                  }
                  return{
                    value : void 0,
                    done : true
                  };
                }
                /** @type {string} */
                data.method = type;
                data.arg = key;
                for (;;) {
                  var name = data.delegate;
                  if (name) {
                    var v = next(name, data);
                    if (v) {
                      if (v === value) {
                        continue;
                      }
                      return v;
                    }
                  }
                  if ("next" === data.method) {
                    data.sent = data._sent = data.arg;
                  } else {
                    if ("throw" === data.method) {
                      if ("suspendedStart" === arLength) {
                        throw arLength = "completed", data.arg;
                      }
                      data.dispatchException(data.arg);
                    } else {
                      if ("return" === data.method) {
                        data.abrupt("return", data.arg);
                      }
                    }
                  }
                  /** @type {string} */
                  arLength = "executing";
                  var item = debug(config, path, data);
                  if ("normal" === item.type) {
                    if (arLength = data.done ? "completed" : "suspendedYield", item.arg === value) {
                      continue;
                    }
                    return{
                      value : item.arg,
                      done : data.done
                    };
                  }
                  if ("throw" === item.type) {
                    /** @type {string} */
                    arLength = "completed";
                    /** @type {string} */
                    data.method = "throw";
                    data.arg = item.arg;
                  }
                }
              };
            }(object, name, pdataCur), self;
          }
          /**
           * @param {Function} func
           * @param {?} arg
           * @param {Object} text
           * @return {?}
           */
          function debug(func, arg, text) {
            try {
              return{
                type : "normal",
                arg : func.call(arg, text)
              };
            } catch (param) {
              return{
                type : "throw",
                arg : param
              };
            }
          }
          /**
           * @return {undefined}
           */
          function superClass() {
          }
          /**
           * @return {undefined}
           */
          function method() {
          }
          /**
           * @return {undefined}
           */
          function prop() {
          }
          /**
           * @param {Object} obj
           * @return {undefined}
           */
          function onComplete(obj) {
            ["next", "throw", "return"].forEach(function(val) {
              /**
               * @param {Object} event
               * @return {?}
               */
              obj[val] = function(event) {
                return this._invoke(val, event);
              };
            });
          }
          /**
           * @param {Object} d
           * @return {undefined}
           */
          function handler(d) {
            var promise;
            /**
             * @param {string} ms
             * @param {Object} action
             * @return {?}
             */
            this._invoke = function(ms, action) {
              /**
               * @return {?}
               */
              function reject() {
                return new Promise(function(next_callback, opt_e) {
                  !function parse(str, body, callback, val) {
                    var data = debug(d[str], d, body);
                    if ("throw" !== data.type) {
                      var a = data.arg;
                      var expectedNumberOfNonCommentArgs = a.value;
                      return expectedNumberOfNonCommentArgs && ("object" == typeof expectedNumberOfNonCommentArgs && hasOwnProperty.call(expectedNumberOfNonCommentArgs, "__await")) ? Promise.resolve(expectedNumberOfNonCommentArgs.__await).then(function(w) {
                        parse("next", w, callback, val);
                      }, function(w) {
                        parse("throw", w, callback, val);
                      }) : Promise.resolve(expectedNumberOfNonCommentArgs).then(function(e) {
                        a.value = e;
                        callback(a);
                      }, function(w) {
                        return parse("throw", w, callback, val);
                      });
                    }
                    val(data.arg);
                  }(ms, action, next_callback, opt_e);
                });
              }
              return promise = promise ? promise.then(reject, reject) : reject();
            };
          }
          /**
           * @param {?} options
           * @param {Object} data
           * @return {?}
           */
          function next(options, data) {
            var obj = options.iterator[data.method];
            if (void 0 === obj) {
              if (data.delegate = null, "throw" === data.method) {
                if (options.iterator.return && (data.method = "return", data.arg = void 0, next(options, data), "throw" === data.method)) {
                  return value;
                }
                /** @type {string} */
                data.method = "throw";
                /** @type {TypeError} */
                data.arg = new TypeError("The iterator does not provide a 'throw' method");
              }
              return value;
            }
            var self = debug(obj, options.iterator, data.arg);
            if ("throw" === self.type) {
              return data.method = "throw", data.arg = self.arg, data.delegate = null, value;
            }
            var v = self.arg;
            return v ? v.done ? (data[options.resultName] = v.value, data.next = options.nextLoc, "return" !== data.method && (data.method = "next", data.arg = void 0), data.delegate = null, value) : v : (data.method = "throw", data.arg = new TypeError("iterator result is not an object"), data.delegate = null, value);
          }
          /**
           * @param {Array} property
           * @return {undefined}
           */
          function addProperty(property) {
            var copies = {
              tryLoc : property[0]
            };
            if (1 in property) {
              copies.catchLoc = property[1];
            }
            if (2 in property) {
              copies.finallyLoc = property[2];
              copies.afterLoc = property[3];
            }
            this.tryEntries.push(copies);
          }
          /**
           * @param {?} httpServer
           * @return {undefined}
           */
          function start(httpServer) {
            var me = httpServer.completion || {};
            /** @type {string} */
            me.type = "normal";
            delete me.arg;
            httpServer.completion = me;
          }
          /**
           * @param {Array} key
           * @return {undefined}
           */
          function Class(key) {
            /** @type {Array} */
            this.tryEntries = [{
              tryLoc : "root"
            }];
            key.forEach(addProperty, this);
            this.reset(true);
          }
          /**
           * @param {Array} value
           * @return {?}
           */
          function val(value) {
            if (value) {
              var values = value[key];
              if (values) {
                return values.call(value);
              }
              if ("function" == typeof value.next) {
                return value;
              }
              if (!isNaN(value.length)) {
                /** @type {number} */
                var index = -1;
                /**
                 * @return {?}
                 */
                var list = function s() {
                  for (;++index < value.length;) {
                    if (hasOwnProperty.call(value, index)) {
                      return s.value = value[index], s.done = false, s;
                    }
                  }
                  return s.value = void 0, s.done = true, s;
                };
                return list.next = list;
              }
            }
            return{
              /** @type {function (): ?} */
              next : state
            };
          }
          /**
           * @return {?}
           */
          function state() {
            return{
              value : void 0,
              done : true
            };
          }
          var ObjProto = Object.prototype;
          /** @type {function (this:Object, *): boolean} */
          var hasOwnProperty = ObjProto.hasOwnProperty;
          var container = "function" == typeof Symbol ? Symbol : {};
          var key = container.iterator || "@@iterator";
          var i = container.asyncIterator || "@@asyncIterator";
          var p = container.toStringTag || "@@toStringTag";
          /** @type {function (Function, Function, string, Array): ?} */
          self.wrap = cb;
          var value = {};
          var expectedNumberOfNonCommentArgs = {};
          /**
           * @return {?}
           */
          expectedNumberOfNonCommentArgs[key] = function() {
            return this;
          };
          /** @type {function (Object): (Object|null)} */
          var getPrototypeOf = Object.getPrototypeOf;
          /** @type {(Object|null)} */
          var ctor = getPrototypeOf && getPrototypeOf(getPrototypeOf(val([])));
          if (ctor) {
            if (ctor !== ObjProto) {
              if (hasOwnProperty.call(ctor, key)) {
                /** @type {Object} */
                expectedNumberOfNonCommentArgs = ctor;
              }
            }
          }
          /** @type {Object} */
          var base = prop.prototype = superClass.prototype = Object.create(expectedNumberOfNonCommentArgs);
          return method.prototype = base.constructor = prop, prop.constructor = method, prop[p] = method.displayName = "GeneratorFunction", self.isGeneratorFunction = function(recurring) {
            /** @type {(Function|boolean|null)} */
            var func = "function" == typeof recurring && recurring.constructor;
            return!!func && (func === method || "GeneratorFunction" === (func.displayName || func.name));
          }, self.mark = function(expectedNumberOfNonCommentArgs) {
            return Object.setPrototypeOf ? Object.setPrototypeOf(expectedNumberOfNonCommentArgs, prop) : (expectedNumberOfNonCommentArgs.__proto__ = prop, p in expectedNumberOfNonCommentArgs || (expectedNumberOfNonCommentArgs[p] = "GeneratorFunction")), expectedNumberOfNonCommentArgs.prototype = Object.create(base), expectedNumberOfNonCommentArgs;
          }, self.awrap = function(dataAndEvents) {
            return{
              __await : dataAndEvents
            };
          }, onComplete(handler.prototype), handler.prototype[i] = function() {
            return this;
          }, self.AsyncIterator = handler, self.async = function(str, recurring, ids, x) {
            var stream = new handler(cb(str, recurring, ids, x));
            return self.isGeneratorFunction(recurring) ? stream : stream.next().then(function(d) {
              return d.done ? d.value : stream.next();
            });
          }, onComplete(base), base[p] = "Generator", base[key] = function() {
            return this;
          }, base.toString = function() {
            return "[object Generator]";
          }, self.keys = function(expectedNumberOfNonCommentArgs) {
            /** @type {Array} */
            var eventPath = [];
            var cur;
            for (cur in expectedNumberOfNonCommentArgs) {
              eventPath.push(cur);
            }
            return eventPath.reverse(), function init() {
              for (;eventPath.length;) {
                var result = eventPath.pop();
                if (result in expectedNumberOfNonCommentArgs) {
                  return init.value = result, init.done = false, init;
                }
              }
              return init.done = true, init;
            };
          }, self.values = val, Class.prototype = {
            /** @type {function (Array): undefined} */
            constructor : Class,
            /**
             * @param {boolean} dataAndEvents
             * @return {undefined}
             */
            reset : function(dataAndEvents) {
              if (this.prev = 0, this.next = 0, this.sent = this._sent = void 0, this.done = false, this.delegate = null, this.method = "next", this.arg = void 0, this.tryEntries.forEach(start), !dataAndEvents) {
                var header;
                for (header in this) {
                  if ("t" === header.charAt(0)) {
                    if (hasOwnProperty.call(this, header)) {
                      if (!isNaN(+header.slice(1))) {
                        this[header] = void 0;
                      }
                    }
                  }
                }
              }
            },
            /**
             * @return {?}
             */
            stop : function() {
              /** @type {boolean} */
              this.done = true;
              var me = this.tryEntries[0].completion;
              if ("throw" === me.type) {
                throw me.arg;
              }
              return this.rval;
            },
            /**
             * @param {?} arg
             * @return {?}
             */
            dispatchException : function(arg) {
              /**
               * @param {?} el
               * @param {boolean} signal_eof
               * @return {?}
               */
              function next(el, signal_eof) {
                return op.type = "throw", op.arg = arg, opts.next = el, signal_eof && (opts.method = "next", opts.arg = void 0), !!signal_eof;
              }
              if (this.done) {
                throw arg;
              }
              var opts = this;
              /** @type {number} */
              var s = this.tryEntries.length - 1;
              for (;s >= 0;--s) {
                var self = this.tryEntries[s];
                var op = self.completion;
                if ("root" === self.tryLoc) {
                  return next("end");
                }
                if (self.tryLoc <= this.prev) {
                  /** @type {boolean} */
                  var format = hasOwnProperty.call(self, "catchLoc");
                  /** @type {boolean} */
                  var useFormat = hasOwnProperty.call(self, "finallyLoc");
                  if (format && useFormat) {
                    if (this.prev < self.catchLoc) {
                      return next(self.catchLoc, true);
                    }
                    if (this.prev < self.finallyLoc) {
                      return next(self.finallyLoc);
                    }
                  } else {
                    if (format) {
                      if (this.prev < self.catchLoc) {
                        return next(self.catchLoc, true);
                      }
                    } else {
                      if (!useFormat) {
                        throw new Error("try statement without catch or finally");
                      }
                      if (this.prev < self.finallyLoc) {
                        return next(self.finallyLoc);
                      }
                    }
                  }
                }
              }
            },
            /**
             * @param {string} type
             * @param {?} arg
             * @return {?}
             */
            abrupt : function(type, arg) {
              /** @type {number} */
              var unlock = this.tryEntries.length - 1;
              for (;unlock >= 0;--unlock) {
                var cache = this.tryEntries[unlock];
                if (cache.tryLoc <= this.prev && (hasOwnProperty.call(cache, "finallyLoc") && this.prev < cache.finallyLoc)) {
                  var item = cache;
                  break;
                }
              }
              if (item) {
                if ("break" === type || "continue" === type) {
                  if (item.tryLoc <= arg) {
                    if (arg <= item.finallyLoc) {
                      /** @type {null} */
                      item = null;
                    }
                  }
                }
              }
              var data = item ? item.completion : {};
              return data.type = type, data.arg = arg, item ? (this.method = "next", this.next = item.finallyLoc, value) : this.complete(data);
            },
            /**
             * @param {?} data
             * @param {Object} next
             * @return {?}
             */
            complete : function(data, next) {
              if ("throw" === data.type) {
                throw data.arg;
              }
              return "break" === data.type || "continue" === data.type ? this.next = data.arg : "return" === data.type ? (this.rval = this.arg = data.arg, this.method = "return", this.next = "end") : "normal" === data.type && (next && (this.next = next)), value;
            },
            /**
             * @param {?} onComplete
             * @return {?}
             */
            finish : function(onComplete) {
              /** @type {number} */
              var unlock = this.tryEntries.length - 1;
              for (;unlock >= 0;--unlock) {
                var httpServer = this.tryEntries[unlock];
                if (httpServer.finallyLoc === onComplete) {
                  return this.complete(httpServer.completion, httpServer.afterLoc), start(httpServer), value;
                }
              }
            },
            /**
             * @param {Function} opt_attributes
             * @return {?}
             */
            catch : function(opt_attributes) {
              /** @type {number} */
              var unlock = this.tryEntries.length - 1;
              for (;unlock >= 0;--unlock) {
                var httpServer = this.tryEntries[unlock];
                if (httpServer.tryLoc === opt_attributes) {
                  var me = httpServer.completion;
                  if ("throw" === me.type) {
                    var arg = me.arg;
                    start(httpServer);
                  }
                  return arg;
                }
              }
              throw new Error("illegal catch attempt");
            },
            /**
             * @param {Object} isXML
             * @param {string} dataAndEvents
             * @param {?} deepDataAndEvents
             * @return {?}
             */
            delegateYield : function(isXML, dataAndEvents, deepDataAndEvents) {
              return this.delegate = {
                iterator : val(isXML),
                resultName : dataAndEvents,
                nextLoc : deepDataAndEvents
              }, "next" === this.method && (this.arg = void 0), value;
            }
          }, self;
        }(module.exports);
        try {
          regeneratorRuntime = html5;
        } catch (t) {
          Function("r", "regeneratorRuntime = r")(html5);
        }
      }, function(dataAndEvents, deepDataAndEvents) {
      }, function(dataAndEvents, global, $) {
        (function(dataAndEvents) {
          /**
           * @return {?}
           */
          function fn() {
            return self.TYPED_ARRAY_SUPPORT ? 2147483647 : 1073741823;
          }
          /**
           * @param {string} value
           * @param {number} arg
           * @return {?}
           */
          function require(value, arg) {
            if (fn() < arg) {
              throw new RangeError("Invalid typed array length");
            }
            return self.TYPED_ARRAY_SUPPORT ? (value = new Uint8Array(arg)).__proto__ = self.prototype : (null === value && (value = new self(arg)), value.length = arg), value;
          }
          /**
           * @param {number} value
           * @param {string} data
           * @param {string} isXML
           * @return {?}
           */
          function self(value, data, isXML) {
            if (!(self.TYPED_ARRAY_SUPPORT || this instanceof self)) {
              return new self(value, data, isXML);
            }
            if ("number" == typeof value) {
              if ("string" == typeof data) {
                throw new Error("If encoding is specified then the first argument must be a string");
              }
              return invoke(this, value);
            }
            return write(this, value, data, isXML);
          }
          /**
           * @param {string} string
           * @param {?} b
           * @param {string} d
           * @param {?} value
           * @return {?}
           */
          function write(string, b, d, value) {
            if ("number" == typeof b) {
              throw new TypeError('"value" argument must not be a number');
            }
            return "undefined" != typeof ArrayBuffer && b instanceof ArrayBuffer ? function(accumulator, value, offset, number) {
              if (value.byteLength, offset < 0 || value.byteLength < offset) {
                throw new RangeError("'offset' is out of bounds");
              }
              if (value.byteLength < offset + (number || 0)) {
                throw new RangeError("'length' is out of bounds");
              }
              return value = void 0 === offset && void 0 === number ? new Uint8Array(value) : void 0 === number ? new Uint8Array(value, offset) : new Uint8Array(value, offset, number), self.TYPED_ARRAY_SUPPORT ? (accumulator = value).__proto__ = self.prototype : accumulator = callback(accumulator, value), accumulator;
            }(string, b, d, value) : "string" == typeof b ? function(source, data, i) {
              if ("string" == typeof i && "" !== i || (i = "utf8"), !self.isEncoding(i)) {
                throw new TypeError('"encoding" must be a valid string encoding');
              }
              /** @type {number} */
              var id = 0 | decode(data, i);
              var n = (source = require(source, id)).write(data, i);
              return n !== id && (source = source.slice(0, n)), source;
            }(string, b, d) : function(data, object) {
              if (self.isBuffer(object)) {
                /** @type {number} */
                var end = 0 | cb(object.length);
                return 0 === (data = require(data, end)).length || object.copy(data, 0, 0, end), data;
              }
              if (object) {
                if ("undefined" != typeof ArrayBuffer && object.buffer instanceof ArrayBuffer || "length" in object) {
                  return "number" != typeof object.length || (length = object.length) != length ? require(data, 0) : callback(data, object);
                }
                if ("Buffer" === object.type && size(object.data)) {
                  return callback(data, object.data);
                }
              }
              var length;
              throw new TypeError("First argument must be a string, Buffer, ArrayBuffer, Array, or array-like object.");
            }(string, b);
          }
          /**
           * @param {number} arg
           * @return {undefined}
           */
          function forEach(arg) {
            if ("number" != typeof arg) {
              throw new TypeError('"size" argument must be a number');
            }
            if (arg < 0) {
              throw new RangeError('"size" argument must not be negative');
            }
          }
          /**
           * @param {(Array|string)} b
           * @param {number} obj
           * @return {?}
           */
          function invoke(b, obj) {
            if (forEach(obj), b = require(b, obj < 0 ? 0 : 0 | cb(obj)), !self.TYPED_ARRAY_SUPPORT) {
              /** @type {number} */
              var bi = 0;
              for (;bi < obj;++bi) {
                /** @type {number} */
                b[bi] = 0;
              }
            }
            return b;
          }
          /**
           * @param {(Array|string)} data
           * @param {Array} self
           * @return {?}
           */
          function callback(data, self) {
            /** @type {number} */
            var key = self.length < 0 ? 0 : 0 | cb(self.length);
            data = require(data, key);
            /** @type {number} */
            var i = 0;
            for (;i < key;i += 1) {
              /** @type {number} */
              data[i] = 255 & self[i];
            }
            return data;
          }
          /**
           * @param {number} label
           * @return {?}
           */
          function cb(label) {
            if (label >= fn()) {
              throw new RangeError("Attempt to allocate Buffer larger than maximum size: 0x" + fn().toString(16) + " bytes");
            }
            return 0 | label;
          }
          /**
           * @param {string} data
           * @param {string} encoding
           * @return {?}
           */
          function decode(data, encoding) {
            if (self.isBuffer(data)) {
              return data.length;
            }
            if ("undefined" != typeof ArrayBuffer && ("function" == typeof ArrayBuffer.isView && (ArrayBuffer.isView(data) || data instanceof ArrayBuffer))) {
              return data.byteLength;
            }
            if ("string" != typeof data) {
              /** @type {string} */
              data = "" + data;
            }
            var pending = data.length;
            if (0 === pending) {
              return 0;
            }
            /** @type {boolean} */
            var n = false;
            for (;;) {
              switch(encoding) {
                case "ascii":
                ;
                case "latin1":
                ;
                case "binary":
                  return pending;
                case "utf8":
                ;
                case "utf-8":
                ;
                case void 0:
                  return encode(data).length;
                case "ucs2":
                ;
                case "ucs-2":
                ;
                case "utf16le":
                ;
                case "utf-16le":
                  return 2 * pending;
                case "hex":
                  return pending >>> 1;
                case "base64":
                  return get(data).length;
                default:
                  if (n) {
                    return encode(data).length;
                  }
                  /** @type {string} */
                  encoding = ("" + encoding).toLowerCase();
                  /** @type {boolean} */
                  n = true;
              }
            }
          }
          /**
           * @param {string} encoding
           * @param {number} value
           * @param {number} val
           * @return {?}
           */
          function query(encoding, value, val) {
            /** @type {boolean} */
            var n = false;
            if ((void 0 === value || value < 0) && (value = 0), value > this.length) {
              return "";
            }
            if ((void 0 === val || val > this.length) && (val = this.length), val <= 0) {
              return "";
            }
            if ((val >>>= 0) <= (value >>>= 0)) {
              return "";
            }
            if (!encoding) {
              /** @type {string} */
              encoding = "utf8";
            }
            for (;;) {
              switch(encoding) {
                case "hex":
                  return wrap(this, value, val);
                case "utf8":
                ;
                case "utf-8":
                  return bind(this, value, val);
                case "ascii":
                  return compare(this, value, val);
                case "latin1":
                ;
                case "binary":
                  return set(this, value, val);
                case "base64":
                  return startsWith(this, value, val);
                case "ucs2":
                ;
                case "ucs-2":
                ;
                case "utf16le":
                ;
                case "utf-16le":
                  return mixin(this, value, val);
                default:
                  if (n) {
                    throw new TypeError("Unknown encoding: " + encoding);
                  }
                  /** @type {string} */
                  encoding = (encoding + "").toLowerCase();
                  /** @type {boolean} */
                  n = true;
              }
            }
          }
          /**
           * @param {Array} results
           * @param {number} prop
           * @param {number} i
           * @return {undefined}
           */
          function sprintf(results, prop, i) {
            var result = results[prop];
            results[prop] = results[i];
            results[i] = result;
          }
          /**
           * @param {string} array
           * @param {number} value
           * @param {number} index
           * @param {number} key
           * @param {boolean} recurring
           * @return {?}
           */
          function _find(array, value, index, key, recurring) {
            if (0 === array.length) {
              return-1;
            }
            if ("string" == typeof index ? (key = index, index = 0) : index > 2147483647 ? index = 2147483647 : index < -2147483648 && (index = -2147483648), index = +index, isNaN(index) && (index = recurring ? 0 : array.length - 1), index < 0 && (index = array.length + index), index >= array.length) {
              if (recurring) {
                return-1;
              }
              /** @type {number} */
              index = array.length - 1;
            } else {
              if (index < 0) {
                if (!recurring) {
                  return-1;
                }
                /** @type {number} */
                index = 0;
              }
            }
            if ("string" == typeof value && (value = self.from(value, key)), self.isBuffer(value)) {
              return 0 === value.length ? -1 : iterator(array, value, index, key, recurring);
            }
            if ("number" == typeof value) {
              return value &= 255, self.TYPED_ARRAY_SUPPORT && "function" == typeof Uint8Array.prototype.indexOf ? recurring ? Uint8Array.prototype.indexOf.call(array, value, index) : Uint8Array.prototype.lastIndexOf.call(array, value, index) : iterator(array, [value], index, key, recurring);
            }
            throw new TypeError("val must be string, number or Buffer");
          }
          /**
           * @param {string} x
           * @param {?} value
           * @param {number} i
           * @param {string} obj
           * @param {boolean} recurring
           * @return {?}
           */
          function iterator(x, value, i, obj, recurring) {
            /**
             * @param {Object} obj
             * @param {number} value
             * @return {?}
             */
            function iterator(obj, value) {
              return 1 === factor ? obj[value] : obj.readUInt16BE(value * factor);
            }
            var max;
            /** @type {number} */
            var factor = 1;
            var n = x.length;
            var j = value.length;
            if (void 0 !== obj && ("ucs2" === (obj = String(obj).toLowerCase()) || ("ucs-2" === obj || ("utf16le" === obj || "utf-16le" === obj)))) {
              if (x.length < 2 || value.length < 2) {
                return-1;
              }
              /** @type {number} */
              factor = 2;
              n /= 2;
              j /= 2;
              i /= 2;
            }
            if (recurring) {
              /** @type {number} */
              var aux = -1;
              /** @type {number} */
              max = i;
              for (;max < n;max++) {
                if (iterator(x, max) === iterator(value, -1 === aux ? 0 : max - aux)) {
                  if (-1 === aux && (aux = max), max - aux + 1 === j) {
                    return aux * factor;
                  }
                } else {
                  if (-1 !== aux) {
                    max -= max - aux;
                  }
                  /** @type {number} */
                  aux = -1;
                }
              }
            } else {
              if (i + j > n) {
                /** @type {number} */
                i = n - j;
              }
              /** @type {number} */
              max = i;
              for (;max >= 0;max--) {
                /** @type {boolean} */
                var h = true;
                /** @type {number} */
                var udataCur = 0;
                for (;udataCur < j;udataCur++) {
                  if (iterator(x, max + udataCur) !== iterator(value, udataCur)) {
                    /** @type {boolean} */
                    h = false;
                    break;
                  }
                }
                if (h) {
                  return max;
                }
              }
            }
            return-1;
          }
          /**
           * @param {(Array|number)} buffer
           * @param {string} values
           * @param {number} size
           * @param {number} end
           * @return {?}
           */
          function parse(buffer, values, size, end) {
            /** @type {number} */
            size = Number(size) || 0;
            /** @type {number} */
            var j = buffer.length - size;
            if (end) {
              if ((end = Number(end)) > j) {
                /** @type {number} */
                end = j;
              }
            } else {
              /** @type {number} */
              end = j;
            }
            var valuesLen = values.length;
            if (valuesLen % 2 != 0) {
              throw new TypeError("Invalid hex string");
            }
            if (end > valuesLen / 2) {
              /** @type {number} */
              end = valuesLen / 2;
            }
            /** @type {number} */
            var x = 0;
            for (;x < end;++x) {
              /** @type {number} */
              var num2 = parseInt(values.substr(2 * x, 2), 16);
              if (isNaN(num2)) {
                return x;
              }
              /** @type {number} */
              buffer[size + x] = num2;
            }
            return x;
          }
          /**
           * @param {?} key
           * @param {string} data
           * @param {number} point
           * @param {number} selector
           * @return {?}
           */
          function store(key, data, point, selector) {
            return indexOf(encode(data, key.length - point), key, point, selector);
          }
          /**
           * @param {?} key
           * @param {string} data
           * @param {number} values
           * @param {number} expected
           * @return {?}
           */
          function validate(key, data, values, expected) {
            return indexOf(function(chars) {
              /** @type {Array} */
              var valid = [];
              /** @type {number} */
              var y = 0;
              for (;y < chars.length;++y) {
                valid.push(255 & chars.charCodeAt(y));
              }
              return valid;
            }(data), key, values, expected);
          }
          /**
           * @param {?} path
           * @param {string} message
           * @param {number} value
           * @param {number} expected
           * @return {?}
           */
          function push(path, message, value, expected) {
            return validate(path, message, value, expected);
          }
          /**
           * @param {?} c
           * @param {string} data
           * @param {number} offset
           * @param {number} fn
           * @return {?}
           */
          function apply(c, data, offset, fn) {
            return indexOf(get(data), c, offset, fn);
          }
          /**
           * @param {?} key
           * @param {string} data
           * @param {number} offset
           * @param {number} selector
           * @return {?}
           */
          function expand(key, data, offset, selector) {
            return indexOf(function(bytes, dataAndEvents) {
              var position;
              var copies;
              var offset;
              /** @type {Array} */
              var out = [];
              /** @type {number} */
              var pos = 0;
              for (;pos < bytes.length && !((dataAndEvents -= 2) < 0);++pos) {
                /** @type {number} */
                copies = (position = bytes.charCodeAt(pos)) >> 8;
                /** @type {number} */
                offset = position % 256;
                out.push(offset);
                out.push(copies);
              }
              return out;
            }(data, key.length - offset), key, offset, selector);
          }
          /**
           * @param {Array} c
           * @param {number} a
           * @param {?} position
           * @return {?}
           */
          function startsWith(c, a, position) {
            return 0 === a && position === c.length ? _this.fromByteArray(c) : _this.fromByteArray(c.slice(a, position));
          }
          /**
           * @param {Arguments} context
           * @param {number} count
           * @param {number} n
           * @return {?}
           */
          function bind(context, count, n) {
            /** @type {number} */
            n = Math.min(context.length, n);
            /** @type {Array} */
            var out = [];
            /** @type {number} */
            var i = count;
            for (;i < n;) {
              var parentSel;
              var value;
              var partial;
              var cc;
              var old = context[i];
              /** @type {null} */
              var copies = null;
              /** @type {number} */
              var segmentLength = old > 239 ? 4 : old > 223 ? 3 : old > 191 ? 2 : 1;
              if (i + segmentLength <= n) {
                switch(segmentLength) {
                  case 1:
                    if (old < 128) {
                      copies = old;
                    }
                    break;
                  case 2:
                    if (128 == (192 & (parentSel = context[i + 1]))) {
                      if ((cc = (31 & old) << 6 | 63 & parentSel) > 127) {
                        /** @type {number} */
                        copies = cc;
                      }
                    }
                    break;
                  case 3:
                    parentSel = context[i + 1];
                    value = context[i + 2];
                    if (128 == (192 & parentSel)) {
                      if (128 == (192 & value)) {
                        if ((cc = (15 & old) << 12 | (63 & parentSel) << 6 | 63 & value) > 2047) {
                          if (cc < 55296 || cc > 57343) {
                            /** @type {number} */
                            copies = cc;
                          }
                        }
                      }
                    }
                    break;
                  case 4:
                    parentSel = context[i + 1];
                    value = context[i + 2];
                    partial = context[i + 3];
                    if (128 == (192 & parentSel)) {
                      if (128 == (192 & value)) {
                        if (128 == (192 & partial)) {
                          if ((cc = (15 & old) << 18 | (63 & parentSel) << 12 | (63 & value) << 6 | 63 & partial) > 65535) {
                            if (cc < 1114112) {
                              /** @type {number} */
                              copies = cc;
                            }
                          }
                        }
                      }
                    }
                  ;
                }
              }
              if (null === copies) {
                /** @type {number} */
                copies = 65533;
                /** @type {number} */
                segmentLength = 1;
              } else {
                if (copies > 65535) {
                  copies -= 65536;
                  out.push(copies >>> 10 & 1023 | 55296);
                  /** @type {number} */
                  copies = 56320 | 1023 & copies;
                }
              }
              out.push(copies);
              i += segmentLength;
            }
            return function(a) {
              /** @type {number} */
              var aLength = a.length;
              if (aLength <= 4096) {
                return String.fromCharCode.apply(String, a);
              }
              /** @type {string} */
              var str = "";
              /** @type {number} */
              var i = 0;
              for (;i < aLength;) {
                str += String.fromCharCode.apply(String, a.slice(i, i += 4096));
              }
              return str;
            }(out);
          }
          /**
           * @param {Arguments} arr1
           * @param {number} size
           * @param {number} n
           * @return {?}
           */
          function compare(arr1, size, n) {
            /** @type {string} */
            var optsData = "";
            /** @type {number} */
            n = Math.min(arr1.length, n);
            /** @type {number} */
            var i = size;
            for (;i < n;++i) {
              optsData += String.fromCharCode(127 & arr1[i]);
            }
            return optsData;
          }
          /**
           * @param {Arguments} block
           * @param {number} size
           * @param {number} n
           * @return {?}
           */
          function set(block, size, n) {
            /** @type {string} */
            var values = "";
            /** @type {number} */
            n = Math.min(block.length, n);
            /** @type {number} */
            var i = size;
            for (;i < n;++i) {
              values += String.fromCharCode(block[i]);
            }
            return values;
          }
          /**
           * @param {Arguments} array
           * @param {number} len
           * @param {?} value
           * @return {?}
           */
          function wrap(array, len, value) {
            var high = array.length;
            if (!len || len < 0) {
              /** @type {number} */
              len = 0;
            }
            if (!value || (value < 0 || value > high)) {
              value = high;
            }
            /** @type {string} */
            var str = "";
            /** @type {number} */
            var i = len;
            for (;i < value;++i) {
              str += formatArray(array[i]);
            }
            return str;
          }
          /**
           * @param {Object} source
           * @param {number} from
           * @param {number} to
           * @return {?}
           */
          function mixin(source, from, to) {
            var lines = source.slice(from, to);
            /** @type {string} */
            var obj = "";
            /** @type {number} */
            var i = 0;
            for (;i < lines.length;i += 2) {
              obj += String.fromCharCode(lines[i] + 256 * lines[i + 1]);
            }
            return obj;
          }
          /**
           * @param {number} out
           * @param {number} expectedNumberOfNonCommentArgs
           * @param {?} next
           * @return {undefined}
           */
          function transitionDone(out, expectedNumberOfNonCommentArgs, next) {
            if (out % 1 != 0 || out < 0) {
              throw new RangeError("offset is not uint");
            }
            if (out + expectedNumberOfNonCommentArgs > next) {
              throw new RangeError("Trying to access beyond buffer length");
            }
          }
          /**
           * @param {string} c
           * @param {number} data
           * @param {number} target
           * @param {number} expectedNumberOfNonCommentArgs
           * @param {number} opt_attributes
           * @param {number} outstandingDataSize
           * @return {undefined}
           */
          function reset(c, data, target, expectedNumberOfNonCommentArgs, opt_attributes, outstandingDataSize) {
            if (!self.isBuffer(c)) {
              throw new TypeError('"buffer" argument must be a Buffer instance');
            }
            if (data > opt_attributes || data < outstandingDataSize) {
              throw new RangeError('"value" argument is out of bounds');
            }
            if (target + expectedNumberOfNonCommentArgs > c.length) {
              throw new RangeError("Index out of range");
            }
          }
          /**
           * @param {Array} text
           * @param {number} d
           * @param {number} x
           * @param {boolean} recurring
           * @return {undefined}
           */
          function debug(text, d, x, recurring) {
            if (d < 0) {
              d = 65535 + d + 1;
            }
            /** @type {number} */
            var i = 0;
            /** @type {number} */
            var padLength = Math.min(text.length - x, 2);
            for (;i < padLength;++i) {
              /** @type {number} */
              text[x + i] = (d & 255 << 8 * (recurring ? i : 1 - i)) >>> 8 * (recurring ? i : 1 - i);
            }
          }
          /**
           * @param {Array} items
           * @param {number} _xhr
           * @param {number} from
           * @param {boolean} recurring
           * @return {undefined}
           */
          function render(items, _xhr, from, recurring) {
            if (_xhr < 0) {
              _xhr = 4294967295 + _xhr + 1;
            }
            /** @type {number} */
            var i = 0;
            /** @type {number} */
            var padLength = Math.min(items.length - from, 4);
            for (;i < padLength;++i) {
              /** @type {number} */
              items[from + i] = _xhr >>> 8 * (recurring ? i : 3 - i) & 255;
            }
          }
          /**
           * @param {string} expr
           * @param {number} condition
           * @param {number} length
           * @param {number} opt_attributes
           * @param {?} guard
           * @param {?} expression
           * @return {undefined}
           */
          function assert(expr, condition, length, opt_attributes, guard, expression) {
            if (length + opt_attributes > expr.length) {
              throw new RangeError("Index out of range");
            }
            if (length < 0) {
              throw new RangeError("Index out of range");
            }
          }
          /**
           * @param {string} obj
           * @param {number} val
           * @param {number} length
           * @param {boolean} data
           * @param {(Object|boolean|number|string)} type
           * @return {?}
           */
          function ondata(obj, val, length, data, type) {
            return type || assert(obj, 0, length, 4), exports.write(obj, val, length, data, 23, 4), length + 4;
          }
          /**
           * @param {string} obj
           * @param {number} result
           * @param {number} length
           * @param {boolean} cb
           * @param {(Object|boolean|number|string)} type
           * @return {?}
           */
          function log(obj, result, length, cb, type) {
            return type || assert(obj, 0, length, 8), exports.write(obj, result, length, cb, 52, 8), length + 8;
          }
          /**
           * @param {number} value
           * @return {?}
           */
          function formatArray(value) {
            return value < 16 ? "0" + value.toString(16) : value.toString(16);
          }
          /**
           * @param {string} m
           * @param {number} eq
           * @return {?}
           */
          function encode(m, eq) {
            var i;
            eq = eq || 1 / 0;
            var b = m.length;
            /** @type {null} */
            var element = null;
            /** @type {Array} */
            var bProperties = [];
            /** @type {number} */
            var a = 0;
            for (;a < b;++a) {
              if ((i = m.charCodeAt(a)) > 55295 && i < 57344) {
                if (!element) {
                  if (i > 56319) {
                    if ((eq -= 3) > -1) {
                      bProperties.push(239, 191, 189);
                    }
                    continue;
                  }
                  if (a + 1 === b) {
                    if ((eq -= 3) > -1) {
                      bProperties.push(239, 191, 189);
                    }
                    continue;
                  }
                  element = i;
                  continue;
                }
                if (i < 56320) {
                  if ((eq -= 3) > -1) {
                    bProperties.push(239, 191, 189);
                  }
                  element = i;
                  continue;
                }
                /** @type {number} */
                i = 65536 + (element - 55296 << 10 | i - 56320);
              } else {
                if (element) {
                  if ((eq -= 3) > -1) {
                    bProperties.push(239, 191, 189);
                  }
                }
              }
              if (element = null, i < 128) {
                if ((eq -= 1) < 0) {
                  break;
                }
                bProperties.push(i);
              } else {
                if (i < 2048) {
                  if ((eq -= 2) < 0) {
                    break;
                  }
                  bProperties.push(i >> 6 | 192, 63 & i | 128);
                } else {
                  if (i < 65536) {
                    if ((eq -= 3) < 0) {
                      break;
                    }
                    bProperties.push(i >> 12 | 224, i >> 6 & 63 | 128, 63 & i | 128);
                  } else {
                    if (!(i < 1114112)) {
                      throw new Error("Invalid code point");
                    }
                    if ((eq -= 4) < 0) {
                      break;
                    }
                    bProperties.push(i >> 18 | 240, i >> 12 & 63 | 128, i >> 6 & 63 | 128, 63 & i | 128);
                  }
                }
              }
            }
            return bProperties;
          }
          /**
           * @param {string} key
           * @return {?}
           */
          function get(key) {
            return _this.toByteArray(function(result) {
              if ((result = function(str) {
                return str.trim ? str.trim() : str.replace(/^\s+|\s+$/g, "");
              }(result).replace(r20, "")).length < 2) {
                return "";
              }
              for (;result.length % 4 != 0;) {
                result += "=";
              }
              return result;
            }(key));
          }
          /**
           * @param {Array} arr
           * @param {?} obj
           * @param {number} offset
           * @param {number} array
           * @return {?}
           */
          function indexOf(arr, obj, offset, array) {
            /** @type {number} */
            var i = 0;
            for (;i < array && !(i + offset >= obj.length || i >= arr.length);++i) {
              obj[i + offset] = arr[i];
            }
            return i;
          }
          var _this = $(121);
          var exports = $(122);
          var size = $(123);
          /** @type {function (number, string, string): ?} */
          global.Buffer = self;
          /**
           * @param {number} dataAndEvents
           * @return {?}
           */
          global.SlowBuffer = function(dataAndEvents) {
            return+dataAndEvents != dataAndEvents && (dataAndEvents = 0), self.alloc(+dataAndEvents);
          };
          /** @type {number} */
          global.INSPECT_MAX_BYTES = 50;
          self.TYPED_ARRAY_SUPPORT = void 0 !== dataAndEvents.TYPED_ARRAY_SUPPORT ? dataAndEvents.TYPED_ARRAY_SUPPORT : function() {
            try {
              /** @type {Uint8Array} */
              var data = new Uint8Array(1);
              return data.__proto__ = {
                __proto__ : Uint8Array.prototype,
                /**
                 * @return {?}
                 */
                foo : function() {
                  return 42;
                }
              }, 42 === data.foo() && ("function" == typeof data.subarray && 0 === data.subarray(1, 1).byteLength);
            } catch (t) {
              return false;
            }
          }();
          global.kMaxLength = fn();
          /** @type {number} */
          self.poolSize = 8192;
          /**
           * @param {Object} obj
           * @return {?}
           */
          self._augment = function(obj) {
            return obj.__proto__ = self.prototype, obj;
          };
          /**
           * @param {number} callback
           * @param {Object} num
           * @param {?} message
           * @return {?}
           */
          self.from = function(callback, num, message) {
            return write(null, callback, num, message);
          };
          if (self.TYPED_ARRAY_SUPPORT) {
            self.prototype.__proto__ = Uint8Array.prototype;
            /** @type {function (new:Uint8Array, (Array.<number>|ArrayBuffer|ArrayBufferView|null|number), number=, number=): ?} */
            self.__proto__ = Uint8Array;
            if ("undefined" != typeof Symbol) {
              if (Symbol.species) {
                if (self[Symbol.species] === self) {
                  Object.defineProperty(self, Symbol.species, {
                    value : null,
                    configurable : true
                  });
                }
              }
            }
          }
          /**
           * @param {number} until
           * @param {number} type
           * @param {number} isXML
           * @return {?}
           */
          self.alloc = function(until, type, isXML) {
            return function(isXML, until, string, value) {
              return forEach(until), until <= 0 ? require(isXML, until) : void 0 !== string ? "string" == typeof value ? require(isXML, until).fill(string, value) : require(isXML, until).fill(string) : require(isXML, until);
            }(null, until, type, isXML);
          };
          /**
           * @param {number} jsonString
           * @return {?}
           */
          self.allocUnsafe = function(jsonString) {
            return invoke(null, jsonString);
          };
          /**
           * @param {number} walkers
           * @return {?}
           */
          self.allocUnsafeSlow = function(walkers) {
            return invoke(null, walkers);
          };
          /**
           * @param {?} x
           * @return {?}
           */
          self.isBuffer = function(x) {
            return!(null == x || !x._isBuffer);
          };
          /**
           * @param {?} c
           * @param {?} expected
           * @return {?}
           */
          self.compare = function(c, expected) {
            if (!self.isBuffer(c) || !self.isBuffer(expected)) {
              throw new TypeError("Arguments must be Buffers");
            }
            if (c === expected) {
              return 0;
            }
            var k = c.length;
            var n = expected.length;
            /** @type {number} */
            var i = 0;
            /** @type {number} */
            var t = Math.min(k, n);
            for (;i < t;++i) {
              if (c[i] !== expected[i]) {
                k = c[i];
                n = expected[i];
                break;
              }
            }
            return k < n ? -1 : n < k ? 1 : 0;
          };
          /**
           * @param {(number|string)} row
           * @return {?}
           */
          self.isEncoding = function(row) {
            switch(String(row).toLowerCase()) {
              case "hex":
              ;
              case "utf8":
              ;
              case "utf-8":
              ;
              case "ascii":
              ;
              case "latin1":
              ;
              case "binary":
              ;
              case "base64":
              ;
              case "ucs2":
              ;
              case "ucs-2":
              ;
              case "utf16le":
              ;
              case "utf-16le":
                return true;
              default:
                return false;
            }
          };
          /**
           * @param {?} obj
           * @param {?} s
           * @return {?}
           */
          self.concat = function(obj, s) {
            if (!size(obj)) {
              throw new TypeError('"list" argument must be an Array of Buffers');
            }
            if (0 === obj.length) {
              return self.alloc(0);
            }
            var i;
            if (void 0 === s) {
              /** @type {number} */
              s = 0;
              /** @type {number} */
              i = 0;
              for (;i < obj.length;++i) {
                s += obj[i].length;
              }
            }
            var buf = self.allocUnsafe(s);
            /** @type {number} */
            var mbOff = 0;
            /** @type {number} */
            i = 0;
            for (;i < obj.length;++i) {
              var b = obj[i];
              if (!self.isBuffer(b)) {
                throw new TypeError('"list" argument must be an Array of Buffers');
              }
              b.copy(buf, mbOff);
              mbOff += b.length;
            }
            return buf;
          };
          /** @type {function (string, string): ?} */
          self.byteLength = decode;
          /** @type {boolean} */
          self.prototype._isBuffer = true;
          /**
           * @return {?}
           */
          self.prototype.swap16 = function() {
            var l = this.length;
            if (l % 2 != 0) {
              throw new RangeError("Buffer size must be a multiple of 16-bits");
            }
            /** @type {number} */
            var i = 0;
            for (;i < l;i += 2) {
              sprintf(this, i, i + 1);
            }
            return this;
          };
          /**
           * @return {?}
           */
          self.prototype.swap32 = function() {
            var l = this.length;
            if (l % 4 != 0) {
              throw new RangeError("Buffer size must be a multiple of 32-bits");
            }
            /** @type {number} */
            var i = 0;
            for (;i < l;i += 4) {
              sprintf(this, i, i + 3);
              sprintf(this, i + 1, i + 2);
            }
            return this;
          };
          /**
           * @return {?}
           */
          self.prototype.swap64 = function() {
            var l = this.length;
            if (l % 8 != 0) {
              throw new RangeError("Buffer size must be a multiple of 64-bits");
            }
            /** @type {number} */
            var i = 0;
            for (;i < l;i += 8) {
              sprintf(this, i, i + 7);
              sprintf(this, i + 1, i + 6);
              sprintf(this, i + 2, i + 5);
              sprintf(this, i + 3, i + 4);
            }
            return this;
          };
          /**
           * @return {?}
           */
          self.prototype.toString = function() {
            /** @type {number} */
            var i = 0 | this.length;
            return 0 === i ? "" : 0 === arguments.length ? bind(this, 0, i) : query.apply(this, arguments);
          };
          /**
           * @param {number} recurring
           * @return {?}
           */
          self.prototype.equals = function(recurring) {
            if (!self.isBuffer(recurring)) {
              throw new TypeError("Argument must be a Buffer");
            }
            return this === recurring || 0 === self.compare(this, recurring);
          };
          /**
           * @return {?}
           */
          self.prototype.inspect = function() {
            /** @type {string} */
            var optsData = "";
            var pdataCur = global.INSPECT_MAX_BYTES;
            return this.length > 0 && (optsData = this.toString("hex", 0, pdataCur).match(/.{2}/g).join(" "), this.length > pdataCur && (optsData += " ... ")), "<Buffer " + optsData + ">";
          };
          /**
           * @param {Array} b
           * @param {number} recurring
           * @param {number} end
           * @param {number} pos
           * @param {number} length
           * @return {?}
           */
          self.prototype.compare = function(b, recurring, end, pos, length) {
            if (!self.isBuffer(b)) {
              throw new TypeError("Argument must be a Buffer");
            }
            if (void 0 === recurring && (recurring = 0), void 0 === end && (end = b ? b.length : 0), void 0 === pos && (pos = 0), void 0 === length && (length = this.length), recurring < 0 || (end > b.length || (pos < 0 || length > this.length))) {
              throw new RangeError("out of range index");
            }
            if (pos >= length && recurring >= end) {
              return 0;
            }
            if (pos >= length) {
              return-1;
            }
            if (recurring >= end) {
              return 1;
            }
            if (this === b) {
              return 0;
            }
            /** @type {number} */
            var r = (length >>>= 0) - (pos >>>= 0);
            /** @type {number} */
            var g = (end >>>= 0) - (recurring >>>= 0);
            /** @type {number} */
            var min = Math.min(r, g);
            var result = this.slice(pos, length);
            var source = b.slice(recurring, end);
            /** @type {number} */
            var idx = 0;
            for (;idx < min;++idx) {
              if (result[idx] !== source[idx]) {
                r = result[idx];
                g = source[idx];
                break;
              }
            }
            return r < g ? -1 : g < r ? 1 : 0;
          };
          /**
           * @param {?} haystack
           * @param {number} from
           * @param {number} object
           * @return {?}
           */
          self.prototype.includes = function(haystack, from, object) {
            return-1 !== this.indexOf(haystack, from, object);
          };
          /**
           * @param {?} array
           * @param {number} from
           * @param {number} collection
           * @return {?}
           */
          self.prototype.indexOf = function(array, from, collection) {
            return _find(this, array, from, collection, true);
          };
          /**
           * @param {number} haystack
           * @param {number} from
           * @param {number} subKey
           * @return {?}
           */
          self.prototype.lastIndexOf = function(haystack, from, subKey) {
            return _find(this, haystack, from, subKey, false);
          };
          /**
           * @param {string} data
           * @param {number} value
           * @param {number} n
           * @param {number} encoding
           * @return {?}
           */
          self.prototype.write = function(data, value, n, encoding) {
            if (void 0 === value) {
              /** @type {string} */
              encoding = "utf8";
              n = this.length;
              /** @type {number} */
              value = 0;
            } else {
              if (void 0 === n && "string" == typeof value) {
                /** @type {number} */
                encoding = value;
                n = this.length;
                /** @type {number} */
                value = 0;
              } else {
                if (!isFinite(value)) {
                  throw new Error("Buffer.write(string, encoding, offset[, length]) is no longer supported");
                }
                value |= 0;
                if (isFinite(n)) {
                  n |= 0;
                  if (void 0 === encoding) {
                    /** @type {string} */
                    encoding = "utf8";
                  }
                } else {
                  /** @type {number} */
                  encoding = n;
                  n = void 0;
                }
              }
            }
            /** @type {number} */
            var m = this.length - value;
            if ((void 0 === n || n > m) && (n = m), data.length > 0 && (n < 0 || value < 0) || value > this.length) {
              throw new RangeError("Attempt to write outside buffer bounds");
            }
            if (!encoding) {
              /** @type {string} */
              encoding = "utf8";
            }
            /** @type {boolean} */
            var o = false;
            for (;;) {
              switch(encoding) {
                case "hex":
                  return parse(this, data, value, n);
                case "utf8":
                ;
                case "utf-8":
                  return store(this, data, value, n);
                case "ascii":
                  return validate(this, data, value, n);
                case "latin1":
                ;
                case "binary":
                  return push(this, data, value, n);
                case "base64":
                  return apply(this, data, value, n);
                case "ucs2":
                ;
                case "ucs-2":
                ;
                case "utf16le":
                ;
                case "utf-16le":
                  return expand(this, data, value, n);
                default:
                  if (o) {
                    throw new TypeError("Unknown encoding: " + encoding);
                  }
                  /** @type {string} */
                  encoding = ("" + encoding).toLowerCase();
                  /** @type {boolean} */
                  o = true;
              }
            }
          };
          /**
           * @return {?}
           */
          self.prototype.toJSON = function() {
            return{
              type : "Buffer",
              data : Array.prototype.slice.call(this._arr || this, 0)
            };
          };
          /**
           * @param {number} recurring
           * @param {number} value
           * @return {?}
           */
          self.prototype.slice = function(recurring, value) {
            var ret;
            var max = this.length;
            if ((recurring = ~~recurring) < 0 ? (recurring += max) < 0 && (recurring = 0) : recurring > max && (recurring = max), (value = void 0 === value ? max : ~~value) < 0 ? (value += max) < 0 && (value = 0) : value > max && (value = max), value < recurring && (value = recurring), self.TYPED_ARRAY_SUPPORT) {
              (ret = this.subarray(recurring, value)).__proto__ = self.prototype;
            } else {
              /** @type {number} */
              var arg = value - recurring;
              ret = new self(arg, void 0);
              /** @type {number} */
              var i = 0;
              for (;i < arg;++i) {
                ret[i] = this[i + recurring];
              }
            }
            return ret;
          };
          /**
           * @param {number} next
           * @param {number} expectedNumberOfNonCommentArgs
           * @param {?} dataAndEvents
           * @return {?}
           */
          self.prototype.readUIntLE = function(next, expectedNumberOfNonCommentArgs, dataAndEvents) {
            next |= 0;
            expectedNumberOfNonCommentArgs |= 0;
            if (!dataAndEvents) {
              transitionDone(next, expectedNumberOfNonCommentArgs, this.length);
            }
            var c = this[next];
            /** @type {number} */
            var m = 1;
            /** @type {number} */
            var offset = 0;
            for (;++offset < expectedNumberOfNonCommentArgs && (m *= 256);) {
              c += this[next + offset] * m;
            }
            return c;
          };
          /**
           * @param {number} next
           * @param {number} expectedNumberOfNonCommentArgs
           * @param {?} dataAndEvents
           * @return {?}
           */
          self.prototype.readUIntBE = function(next, expectedNumberOfNonCommentArgs, dataAndEvents) {
            next |= 0;
            expectedNumberOfNonCommentArgs |= 0;
            if (!dataAndEvents) {
              transitionDone(next, expectedNumberOfNonCommentArgs, this.length);
            }
            var value = this[next + --expectedNumberOfNonCommentArgs];
            /** @type {number} */
            var base = 1;
            for (;expectedNumberOfNonCommentArgs > 0 && (base *= 256);) {
              value += this[next + --expectedNumberOfNonCommentArgs] * base;
            }
            return value;
          };
          /**
           * @param {number} next
           * @param {undefined} noAssert
           * @return {?}
           */
          self.prototype.readUInt8 = function(next, noAssert) {
            return noAssert || transitionDone(next, 1, this.length), this[next];
          };
          /**
           * @param {number} next
           * @param {undefined} dataAndEvents
           * @return {?}
           */
          self.prototype.readUInt16LE = function(next, dataAndEvents) {
            return dataAndEvents || transitionDone(next, 2, this.length), this[next] | this[next + 1] << 8;
          };
          /**
           * @param {number} next
           * @param {undefined} noAssert
           * @return {?}
           */
          self.prototype.readUInt16BE = function(next, noAssert) {
            return noAssert || transitionDone(next, 2, this.length), this[next] << 8 | this[next + 1];
          };
          /**
           * @param {number} next
           * @param {undefined} noAssert
           * @return {?}
           */
          self.prototype.readUInt32LE = function(next, noAssert) {
            return noAssert || transitionDone(next, 4, this.length), (this[next] | this[next + 1] << 8 | this[next + 2] << 16) + 16777216 * this[next + 3];
          };
          /**
           * @param {number} next
           * @param {undefined} noAssert
           * @return {?}
           */
          self.prototype.readUInt32BE = function(next, noAssert) {
            return noAssert || transitionDone(next, 4, this.length), 16777216 * this[next] + (this[next + 1] << 16 | this[next + 2] << 8 | this[next + 3]);
          };
          /**
           * @param {number} param
           * @param {number} n
           * @param {?} dataAndEvents
           * @return {?}
           */
          self.prototype.readIntLE = function(param, n, dataAndEvents) {
            param |= 0;
            n |= 0;
            if (!dataAndEvents) {
              transitionDone(param, n, this.length);
            }
            var value = this[param];
            /** @type {number} */
            var base = 1;
            /** @type {number} */
            var index = 0;
            for (;++index < n && (base *= 256);) {
              value += this[param + index] * base;
            }
            return value >= (base *= 128) && (value -= Math.pow(2, 8 * n)), value;
          };
          /**
           * @param {number} next
           * @param {number} pos
           * @param {?} dataAndEvents
           * @return {?}
           */
          self.prototype.readIntBE = function(next, pos, dataAndEvents) {
            next |= 0;
            pos |= 0;
            if (!dataAndEvents) {
              transitionDone(next, pos, this.length);
            }
            /** @type {number} */
            var savedPos4 = pos;
            /** @type {number} */
            var base = 1;
            var value = this[next + --savedPos4];
            for (;savedPos4 > 0 && (base *= 256);) {
              value += this[next + --savedPos4] * base;
            }
            return value >= (base *= 128) && (value -= Math.pow(2, 8 * pos)), value;
          };
          /**
           * @param {number} next
           * @param {undefined} signed
           * @return {?}
           */
          self.prototype.readInt8 = function(next, signed) {
            return signed || transitionDone(next, 1, this.length), 128 & this[next] ? -1 * (255 - this[next] + 1) : this[next];
          };
          /**
           * @param {number} next
           * @param {?} noAssert
           * @return {?}
           */
          self.prototype.readInt16LE = function(next, noAssert) {
            if (!noAssert) {
              transitionDone(next, 2, this.length);
            }
            /** @type {number} */
            var unsigned = this[next] | this[next + 1] << 8;
            return 32768 & unsigned ? 4294901760 | unsigned : unsigned;
          };
          /**
           * @param {number} next
           * @param {?} noAssert
           * @return {?}
           */
          self.prototype.readInt16BE = function(next, noAssert) {
            if (!noAssert) {
              transitionDone(next, 2, this.length);
            }
            /** @type {number} */
            var unsigned = this[next + 1] | this[next] << 8;
            return 32768 & unsigned ? 4294901760 | unsigned : unsigned;
          };
          /**
           * @param {number} next
           * @param {undefined} noAssert
           * @return {?}
           */
          self.prototype.readInt32LE = function(next, noAssert) {
            return noAssert || transitionDone(next, 4, this.length), this[next] | this[next + 1] << 8 | this[next + 2] << 16 | this[next + 3] << 24;
          };
          /**
           * @param {number} next
           * @param {undefined} noAssert
           * @return {?}
           */
          self.prototype.readInt32BE = function(next, noAssert) {
            return noAssert || transitionDone(next, 4, this.length), this[next] << 24 | this[next + 1] << 16 | this[next + 2] << 8 | this[next + 3];
          };
          /**
           * @param {number} next
           * @param {undefined} dataAndEvents
           * @return {?}
           */
          self.prototype.readFloatLE = function(next, dataAndEvents) {
            return dataAndEvents || transitionDone(next, 4, this.length), exports.read(this, next, true, 23, 4);
          };
          /**
           * @param {number} next
           * @param {undefined} dataAndEvents
           * @return {?}
           */
          self.prototype.readFloatBE = function(next, dataAndEvents) {
            return dataAndEvents || transitionDone(next, 4, this.length), exports.read(this, next, false, 23, 4);
          };
          /**
           * @param {number} next
           * @param {undefined} dataAndEvents
           * @return {?}
           */
          self.prototype.readDoubleLE = function(next, dataAndEvents) {
            return dataAndEvents || transitionDone(next, 8, this.length), exports.read(this, next, true, 52, 8);
          };
          /**
           * @param {number} next
           * @param {undefined} dataAndEvents
           * @return {?}
           */
          self.prototype.readDoubleBE = function(next, dataAndEvents) {
            return dataAndEvents || transitionDone(next, 8, this.length), exports.read(this, next, false, 52, 8);
          };
          /**
           * @param {number} a
           * @param {number} key
           * @param {number} n
           * @param {?} dataAndEvents
           * @return {?}
           */
          self.prototype.writeUIntLE = function(a, key, n, dataAndEvents) {
            /** @type {number} */
            a = +a;
            key |= 0;
            n |= 0;
            if (!dataAndEvents) {
              reset(this, a, key, n, Math.pow(2, 8 * n) - 1, 0);
            }
            /** @type {number} */
            var b = 1;
            /** @type {number} */
            var index = 0;
            /** @type {number} */
            this[key] = 255 & a;
            for (;++index < n && (b *= 256);) {
              /** @type {number} */
              this[key + index] = a / b & 255;
            }
            return key + n;
          };
          /**
           * @param {number} x
           * @param {number} a
           * @param {number} n
           * @param {?} dataAndEvents
           * @return {?}
           */
          self.prototype.writeUIntBE = function(x, a, n, dataAndEvents) {
            /** @type {number} */
            x = +x;
            a |= 0;
            n |= 0;
            if (!dataAndEvents) {
              reset(this, x, a, n, Math.pow(2, 8 * n) - 1, 0);
            }
            /** @type {number} */
            var b = n - 1;
            /** @type {number} */
            var resolution = 1;
            /** @type {number} */
            this[a + b] = 255 & x;
            for (;--b >= 0 && (resolution *= 256);) {
              /** @type {number} */
              this[a + b] = x / resolution & 255;
            }
            return a + n;
          };
          /**
           * @param {number} value
           * @param {number} opts
           * @param {undefined} noAssert
           * @return {?}
           */
          self.prototype.writeUInt8 = function(value, opts, noAssert) {
            return value = +value, opts |= 0, noAssert || reset(this, value, opts, 1, 255, 0), self.TYPED_ARRAY_SUPPORT || (value = Math.floor(value)), this[opts] = 255 & value, opts + 1;
          };
          /**
           * @param {number} d
           * @param {number} opts
           * @param {undefined} noAssert
           * @return {?}
           */
          self.prototype.writeUInt16LE = function(d, opts, noAssert) {
            return d = +d, opts |= 0, noAssert || reset(this, d, opts, 2, 65535, 0), self.TYPED_ARRAY_SUPPORT ? (this[opts] = 255 & d, this[opts + 1] = d >>> 8) : debug(this, d, opts, true), opts + 2;
          };
          /**
           * @param {number} d
           * @param {number} opts
           * @param {undefined} value
           * @return {?}
           */
          self.prototype.writeUInt16BE = function(d, opts, value) {
            return d = +d, opts |= 0, value || reset(this, d, opts, 2, 65535, 0), self.TYPED_ARRAY_SUPPORT ? (this[opts] = d >>> 8, this[opts + 1] = 255 & d) : debug(this, d, opts, false), opts + 2;
          };
          /**
           * @param {number} data
           * @param {number} opts
           * @param {undefined} noAssert
           * @return {?}
           */
          self.prototype.writeUInt32LE = function(data, opts, noAssert) {
            return data = +data, opts |= 0, noAssert || reset(this, data, opts, 4, 4294967295, 0), self.TYPED_ARRAY_SUPPORT ? (this[opts + 3] = data >>> 24, this[opts + 2] = data >>> 16, this[opts + 1] = data >>> 8, this[opts] = 255 & data) : render(this, data, opts, true), opts + 4;
          };
          /**
           * @param {number} data
           * @param {number} opts
           * @param {undefined} value
           * @return {?}
           */
          self.prototype.writeUInt32BE = function(data, opts, value) {
            return data = +data, opts |= 0, value || reset(this, data, opts, 4, 4294967295, 0), self.TYPED_ARRAY_SUPPORT ? (this[opts] = data >>> 24, this[opts + 1] = data >>> 16, this[opts + 2] = data >>> 8, this[opts + 3] = 255 & data) : render(this, data, opts, false), opts + 4;
          };
          /**
           * @param {number} x
           * @param {number} base
           * @param {number} expectedNumberOfNonCommentArgs
           * @param {?} dataAndEvents
           * @return {?}
           */
          self.prototype.writeIntLE = function(x, base, expectedNumberOfNonCommentArgs, dataAndEvents) {
            if (x = +x, base |= 0, !dataAndEvents) {
              /** @type {number} */
              var i = Math.pow(2, 8 * expectedNumberOfNonCommentArgs - 1);
              reset(this, x, base, expectedNumberOfNonCommentArgs, i - 1, -i);
            }
            /** @type {number} */
            var post = 0;
            /** @type {number} */
            var resolution = 1;
            /** @type {number} */
            var a = 0;
            /** @type {number} */
            this[base] = 255 & x;
            for (;++post < expectedNumberOfNonCommentArgs && (resolution *= 256);) {
              if (x < 0) {
                if (0 === a) {
                  if (0 !== this[base + post - 1]) {
                    /** @type {number} */
                    a = 1;
                  }
                }
              }
              /** @type {number} */
              this[base + post] = (x / resolution >> 0) - a & 255;
            }
            return base + expectedNumberOfNonCommentArgs;
          };
          /**
           * @param {number} x
           * @param {number} base
           * @param {number} expectedNumberOfNonCommentArgs
           * @param {?} dataAndEvents
           * @return {?}
           */
          self.prototype.writeIntBE = function(x, base, expectedNumberOfNonCommentArgs, dataAndEvents) {
            if (x = +x, base |= 0, !dataAndEvents) {
              /** @type {number} */
              var i = Math.pow(2, 8 * expectedNumberOfNonCommentArgs - 1);
              reset(this, x, base, expectedNumberOfNonCommentArgs, i - 1, -i);
            }
            /** @type {number} */
            var post = expectedNumberOfNonCommentArgs - 1;
            /** @type {number} */
            var resolution = 1;
            /** @type {number} */
            var a = 0;
            /** @type {number} */
            this[base + post] = 255 & x;
            for (;--post >= 0 && (resolution *= 256);) {
              if (x < 0) {
                if (0 === a) {
                  if (0 !== this[base + post + 1]) {
                    /** @type {number} */
                    a = 1;
                  }
                }
              }
              /** @type {number} */
              this[base + post] = (x / resolution >> 0) - a & 255;
            }
            return base + expectedNumberOfNonCommentArgs;
          };
          /**
           * @param {number} value
           * @param {number} opts
           * @param {undefined} noAssert
           * @return {?}
           */
          self.prototype.writeInt8 = function(value, opts, noAssert) {
            return value = +value, opts |= 0, noAssert || reset(this, value, opts, 1, 127, -128), self.TYPED_ARRAY_SUPPORT || (value = Math.floor(value)), value < 0 && (value = 255 + value + 1), this[opts] = 255 & value, opts + 1;
          };
          /**
           * @param {number} data
           * @param {number} opts
           * @param {undefined} noAssert
           * @return {?}
           */
          self.prototype.writeInt16LE = function(data, opts, noAssert) {
            return data = +data, opts |= 0, noAssert || reset(this, data, opts, 2, 32767, -32768), self.TYPED_ARRAY_SUPPORT ? (this[opts] = 255 & data, this[opts + 1] = data >>> 8) : debug(this, data, opts, true), opts + 2;
          };
          /**
           * @param {number} data
           * @param {number} opts
           * @param {undefined} noAssert
           * @return {?}
           */
          self.prototype.writeInt16BE = function(data, opts, noAssert) {
            return data = +data, opts |= 0, noAssert || reset(this, data, opts, 2, 32767, -32768), self.TYPED_ARRAY_SUPPORT ? (this[opts] = data >>> 8, this[opts + 1] = 255 & data) : debug(this, data, opts, false), opts + 2;
          };
          /**
           * @param {number} data
           * @param {number} opts
           * @param {undefined} noAssert
           * @return {?}
           */
          self.prototype.writeInt32LE = function(data, opts, noAssert) {
            return data = +data, opts |= 0, noAssert || reset(this, data, opts, 4, 2147483647, -2147483648), self.TYPED_ARRAY_SUPPORT ? (this[opts] = 255 & data, this[opts + 1] = data >>> 8, this[opts + 2] = data >>> 16, this[opts + 3] = data >>> 24) : render(this, data, opts, true), opts + 4;
          };
          /**
           * @param {number} data
           * @param {number} opts
           * @param {undefined} noAssert
           * @return {?}
           */
          self.prototype.writeInt32BE = function(data, opts, noAssert) {
            return data = +data, opts |= 0, noAssert || reset(this, data, opts, 4, 2147483647, -2147483648), data < 0 && (data = 4294967295 + data + 1), self.TYPED_ARRAY_SUPPORT ? (this[opts] = data >>> 24, this[opts + 1] = data >>> 16, this[opts + 2] = data >>> 8, this[opts + 3] = 255 & data) : render(this, data, opts, false), opts + 4;
          };
          /**
           * @param {number} val
           * @param {number} extra
           * @param {string} fix
           * @return {?}
           */
          self.prototype.writeFloatLE = function(val, extra, fix) {
            return ondata(this, val, extra, true, fix);
          };
          /**
           * @param {number} val
           * @param {number} extra
           * @param {string} fix
           * @return {?}
           */
          self.prototype.writeFloatBE = function(val, extra, fix) {
            return ondata(this, val, extra, false, fix);
          };
          /**
           * @param {number} entry
           * @param {number} value
           * @param {string} current
           * @return {?}
           */
          self.prototype.writeDoubleLE = function(entry, value, current) {
            return log(this, entry, value, true, current);
          };
          /**
           * @param {number} entry
           * @param {number} l
           * @param {string} current
           * @return {?}
           */
          self.prototype.writeDoubleBE = function(entry, l, current) {
            return log(this, entry, l, false, current);
          };
          /**
           * @param {string} data
           * @param {number} offset
           * @param {number} start
           * @param {number} end
           * @return {?}
           */
          self.prototype.copy = function(data, offset, start, end) {
            if (start || (start = 0), end || (0 === end || (end = this.length)), offset >= data.length && (offset = data.length), offset || (offset = 0), end > 0 && (end < start && (end = start)), end === start) {
              return 0;
            }
            if (0 === data.length || 0 === this.length) {
              return 0;
            }
            if (offset < 0) {
              throw new RangeError("targetStart out of bounds");
            }
            if (start < 0 || start >= this.length) {
              throw new RangeError("sourceStart out of bounds");
            }
            if (end < 0) {
              throw new RangeError("sourceEnd out of bounds");
            }
            if (end > this.length) {
              end = this.length;
            }
            if (data.length - offset < end - start) {
              end = data.length - offset + start;
            }
            var i;
            /** @type {number} */
            var len = end - start;
            if (this === data && (start < offset && offset < end)) {
              /** @type {number} */
              i = len - 1;
              for (;i >= 0;--i) {
                data[i + offset] = this[i + start];
              }
            } else {
              if (len < 1E3 || !self.TYPED_ARRAY_SUPPORT) {
                /** @type {number} */
                i = 0;
                for (;i < len;++i) {
                  data[i + offset] = this[i + start];
                }
              } else {
                Uint8Array.prototype.set.call(data, this.subarray(start, start + len), offset);
              }
            }
            return len;
          };
          /**
           * @param {number} arg
           * @param {number} n
           * @param {number} index
           * @param {number} val
           * @return {?}
           */
          self.prototype.fill = function(arg, n, index, val) {
            if ("string" == typeof arg) {
              if ("string" == typeof n ? (val = n, n = 0, index = this.length) : "string" == typeof index && (val = index, index = this.length), 1 === arg.length) {
                /** @type {number} */
                var x = arg.charCodeAt(0);
                if (x < 256) {
                  /** @type {number} */
                  arg = x;
                }
              }
              if (void 0 !== val && "string" != typeof val) {
                throw new TypeError("encoding must be a string");
              }
              if ("string" == typeof val && !self.isEncoding(val)) {
                throw new TypeError("Unknown encoding: " + val);
              }
            } else {
              if ("number" == typeof arg) {
                arg &= 255;
              }
            }
            if (n < 0 || (this.length < n || this.length < index)) {
              throw new RangeError("Out of range index");
            }
            if (index <= n) {
              return this;
            }
            var i;
            if (n >>>= 0, index = void 0 === index ? this.length : index >>> 0, arg || (arg = 0), "number" == typeof arg) {
              /** @type {number} */
              i = n;
              for (;i < index;++i) {
                /** @type {number} */
                this[i] = arg;
              }
            } else {
              var args = self.isBuffer(arg) ? arg : encode((new self(arg, val)).toString());
              var size = args.length;
              /** @type {number} */
              i = 0;
              for (;i < index - n;++i) {
                this[i + n] = args[i % size];
              }
            }
            return this;
          };
          /** @type {RegExp} */
          var r20 = /[^+\/0-9A-Za-z-_]/g;
        }).call(this, $(23));
      }, function(dataAndEvents, b, deepDataAndEvents) {
        /**
         * @param {string} txt
         * @return {?}
         */
        function complete(txt) {
          var fx = txt.length;
          if (fx % 4 > 0) {
            throw new Error("Invalid string. Length must be a multiple of 4");
          }
          var type = txt.indexOf("=");
          return-1 === type && (type = fx), [type, type === fx ? 0 : 4 - type % 4];
        }
        /**
         * @param {Array} text
         * @param {number} str
         * @param {number} b
         * @return {?}
         */
        function log(text, str, b) {
          var fn;
          var bulk;
          /** @type {Array} */
          var tagNameArr = [];
          /** @type {number} */
          var x = str;
          for (;x < b;x += 3) {
            /** @type {number} */
            fn = (text[x] << 16 & 16711680) + (text[x + 1] << 8 & 65280) + (255 & text[x + 2]);
            tagNameArr.push(matches[(bulk = fn) >> 18 & 63] + matches[bulk >> 12 & 63] + matches[bulk >> 6 & 63] + matches[63 & bulk]);
          }
          return tagNameArr.join("");
        }
        /**
         * @param {string} data
         * @return {?}
         */
        b.byteLength = function(data) {
          var d = complete(data);
          var low = d[0];
          var high = d[1];
          return 3 * (low + high) / 4 - high;
        };
        /**
         * @param {string} data
         * @return {?}
         */
        b.toByteArray = function(data) {
          var BYTE;
          var i;
          var newState = complete(data);
          var expectationResult = newState[0];
          var udataCur = newState[1];
          var res = new MarkerArray(function(recurring, result, value) {
            return 3 * (result + value) / 4 - value;
          }(0, expectationResult, udataCur));
          /** @type {number} */
          var resLength = 0;
          var padLength = udataCur > 0 ? expectationResult - 4 : expectationResult;
          /** @type {number} */
          i = 0;
          for (;i < padLength;i += 4) {
            /** @type {number} */
            BYTE = index[data.charCodeAt(i)] << 18 | index[data.charCodeAt(i + 1)] << 12 | index[data.charCodeAt(i + 2)] << 6 | index[data.charCodeAt(i + 3)];
            /** @type {number} */
            res[resLength++] = BYTE >> 16 & 255;
            /** @type {number} */
            res[resLength++] = BYTE >> 8 & 255;
            /** @type {number} */
            res[resLength++] = 255 & BYTE;
          }
          return 2 === udataCur && (BYTE = index[data.charCodeAt(i)] << 2 | index[data.charCodeAt(i + 1)] >> 4, res[resLength++] = 255 & BYTE), 1 === udataCur && (BYTE = index[data.charCodeAt(i)] << 10 | index[data.charCodeAt(i + 1)] << 4 | index[data.charCodeAt(i + 2)] >> 2, res[resLength++] = BYTE >> 8 & 255, res[resLength++] = 255 & BYTE), res;
        };
        /**
         * @param {Array} c
         * @return {?}
         */
        b.fromByteArray = function(c) {
          var k;
          var a = c.length;
          /** @type {number} */
          var b = a % 3;
          /** @type {Array} */
          var tagNameArr = [];
          /** @type {number} */
          var y = 0;
          /** @type {number} */
          var maxY = a - b;
          for (;y < maxY;y += 16383) {
            tagNameArr.push(log(c, y, y + 16383 > maxY ? maxY : y + 16383));
          }
          return 1 === b ? (k = c[a - 1], tagNameArr.push(matches[k >> 2] + matches[k << 4 & 63] + "==")) : 2 === b && (k = (c[a - 2] << 8) + c[a - 1], tagNameArr.push(matches[k >> 10] + matches[k >> 4 & 63] + matches[k << 2 & 63] + "=")), tagNameArr.join("");
        };
        /** @type {Array} */
        var matches = [];
        /** @type {Array} */
        var index = [];
        /** @type {Function} */
        var MarkerArray = "undefined" != typeof Uint8Array ? Uint8Array : Array;
        /** @type {string} */
        var seed = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
        /** @type {number} */
        var i = 0;
        /** @type {number} */
        var len = seed.length;
        for (;i < len;++i) {
          matches[i] = seed[i];
          /** @type {number} */
          index[seed.charCodeAt(i)] = i;
        }
        /** @type {number} */
        index["-".charCodeAt(0)] = 62;
        /** @type {number} */
        index["_".charCodeAt(0)] = 63;
      }, function(dataAndEvents, gridStore) {
        /**
         * @param {?} args
         * @param {number} out
         * @param {boolean} recurring
         * @param {number} length
         * @param {number} opt_attributes
         * @return {?}
         */
        gridStore.read = function(args, out, recurring, length, opt_attributes) {
          var value;
          var m;
          /** @type {number} */
          var step = 8 * opt_attributes - length - 1;
          /** @type {number} */
          var radio = (1 << step) - 1;
          /** @type {number} */
          var increment = radio >> 1;
          /** @type {number} */
          var i = -7;
          /** @type {number} */
          var body = recurring ? opt_attributes - 1 : 0;
          /** @type {number} */
          var d = recurring ? -1 : 1;
          var pageY = args[out + body];
          body += d;
          /** @type {number} */
          value = pageY & (1 << -i) - 1;
          pageY >>= -i;
          i += step;
          for (;i > 0;value = 256 * value + args[out + body], body += d, i -= 8) {
          }
          /** @type {number} */
          m = value & (1 << -i) - 1;
          value >>= -i;
          i += length;
          for (;i > 0;m = 256 * m + args[out + body], body += d, i -= 8) {
          }
          if (0 === value) {
            /** @type {number} */
            value = 1 - increment;
          } else {
            if (value === radio) {
              return m ? NaN : 1 / 0 * (pageY ? -1 : 1);
            }
            m += Math.pow(2, length);
            value -= increment;
          }
          return(pageY ? -1 : 1) * m * Math.pow(2, value - length);
        };
        /**
         * @param {string} object
         * @param {number} value
         * @param {number} length
         * @param {boolean} content
         * @param {number} mLen
         * @param {number} opt_attributes
         * @return {undefined}
         */
        gridStore.write = function(object, value, length, content, mLen, opt_attributes) {
          var e;
          var m;
          var factor;
          /** @type {number} */
          var eLen = 8 * opt_attributes - mLen - 1;
          /** @type {number} */
          var eMax = (1 << eLen) - 1;
          /** @type {number} */
          var eBias = eMax >> 1;
          /** @type {number} */
          var x = 23 === mLen ? Math.pow(2, -24) - Math.pow(2, -77) : 0;
          /** @type {number} */
          var from = content ? 0 : opt_attributes - 1;
          /** @type {number} */
          var len = content ? 1 : -1;
          /** @type {number} */
          var v = value < 0 || 0 === value && 1 / value < 0 ? 1 : 0;
          /** @type {number} */
          value = Math.abs(value);
          if (isNaN(value) || value === 1 / 0) {
            /** @type {number} */
            m = isNaN(value) ? 1 : 0;
            /** @type {number} */
            e = eMax;
          } else {
            /** @type {number} */
            e = Math.floor(Math.log(value) / Math.LN2);
            if (value * (factor = Math.pow(2, -e)) < 1) {
              e--;
              factor *= 2;
            }
            if ((value += e + eBias >= 1 ? x / factor : x * Math.pow(2, 1 - eBias)) * factor >= 2) {
              e++;
              factor /= 2;
            }
            if (e + eBias >= eMax) {
              /** @type {number} */
              m = 0;
              /** @type {number} */
              e = eMax;
            } else {
              if (e + eBias >= 1) {
                /** @type {number} */
                m = (value * factor - 1) * Math.pow(2, mLen);
                e += eBias;
              } else {
                /** @type {number} */
                m = value * Math.pow(2, eBias - 1) * Math.pow(2, mLen);
                /** @type {number} */
                e = 0;
              }
            }
          }
          for (;mLen >= 8;object[length + from] = 255 & m, from += len, m /= 256, mLen -= 8) {
          }
          /** @type {number} */
          e = e << mLen | m;
          eLen += mLen;
          for (;eLen > 0;object[length + from] = 255 & e, from += len, e /= 256, eLen -= 8) {
          }
          object[length + from - len] |= 128 * v;
        };
      }, function(module, dataAndEvents) {
        /** @type {function (this:*): string} */
        var ostring = {}.toString;
        /** @type {function (*): boolean} */
        module.exports = Array.isArray || function(expectedNumberOfNonCommentArgs) {
          return "[object Array]" == ostring.call(expectedNumberOfNonCommentArgs);
        };
      }, function(dataAndEvents, expectedNumberOfNonCommentArgs, f) {
        f.r(expectedNumberOfNonCommentArgs);
        (function(m) {
          /**
           * @param {string} value
           * @return {?}
           */
          function stringify(value) {
            return value ? encode(value) : "undefined" != typeof navigator ? encode(navigator.userAgent) : replacer();
          }
          /**
           * @param {Object} text
           * @return {?}
           */
          function compress(text) {
            return "" !== text && arr.reduce(function(dataAndEvents, $cookies) {
              var value = $cookies[0];
              var re = $cookies[1];
              if (dataAndEvents) {
                return dataAndEvents;
              }
              var code = re.exec(text);
              return!!code && [value, code];
            }, false);
          }
          /**
           * @param {Object} content
           * @return {?}
           */
          function process(content) {
            var outputReCompress = compress(content);
            return outputReCompress ? outputReCompress[0] : null;
          }
          /**
           * @param {string} input
           * @return {?}
           */
          function encode(input) {
            var output = compress(input);
            if (!output) {
              return null;
            }
            var denom = output[0];
            var style = output[1];
            if ("searchbot" === denom) {
              return new init;
            }
            var $keys = style[1] && style[1].split(/[._]/).slice(0, 3);
            if ($keys) {
              if ($keys.length < 3) {
                $keys = function() {
                  /** @type {number} */
                  var length = 0;
                  /** @type {number} */
                  var j = 0;
                  /** @type {number} */
                  var len = arguments.length;
                  for (;j < len;j++) {
                    length += arguments[j].length;
                  }
                  /** @type {Array} */
                  var result = Array(length);
                  /** @type {number} */
                  var n = 0;
                  /** @type {number} */
                  j = 0;
                  for (;j < len;j++) {
                    var arr = arguments[j];
                    /** @type {number} */
                    var i = 0;
                    var e = arr.length;
                    for (;i < e;i++, n++) {
                      result[n] = arr[i];
                    }
                  }
                  return result;
                }($keys, function(high) {
                  /** @type {Array} */
                  var eventPath = [];
                  /** @type {number} */
                  var low = 0;
                  for (;low < high;low++) {
                    eventPath.push("0");
                  }
                  return eventPath;
                }(3 - $keys.length));
              }
            } else {
              /** @type {Array} */
              $keys = [];
            }
            var c = $keys.join(".");
            var el = createElement(input);
            /** @type {(Array.<string>|null)} */
            var elements = re.exec(input);
            return elements && elements[1] ? new parse(denom, c, el, elements[1]) : new set(denom, $keys.join("."), el);
          }
          /**
           * @param {string} string
           * @return {?}
           */
          function createElement(string) {
            /** @type {number} */
            var unlock = 0;
            /** @type {number} */
            var length = resolveValues.length;
            for (;unlock < length;unlock++) {
              var cache = resolveValues[unlock];
              var result = cache[0];
              if (cache[1].exec(string)) {
                return result;
              }
            }
            return null;
          }
          /**
           * @return {?}
           */
          function replacer() {
            return void 0 !== m && m.version ? new start(m.version.slice(1)) : null;
          }
          f.d(expectedNumberOfNonCommentArgs, "BrowserInfo", function() {
            return set;
          });
          f.d(expectedNumberOfNonCommentArgs, "NodeInfo", function() {
            return start;
          });
          f.d(expectedNumberOfNonCommentArgs, "SearchBotDeviceInfo", function() {
            return parse;
          });
          f.d(expectedNumberOfNonCommentArgs, "BotInfo", function() {
            return init;
          });
          f.d(expectedNumberOfNonCommentArgs, "detect", function() {
            return stringify;
          });
          f.d(expectedNumberOfNonCommentArgs, "browserName", function() {
            return process;
          });
          f.d(expectedNumberOfNonCommentArgs, "parseUserAgent", function() {
            return encode;
          });
          f.d(expectedNumberOfNonCommentArgs, "detectOS", function() {
            return createElement;
          });
          f.d(expectedNumberOfNonCommentArgs, "getNodeVersion", function() {
            return replacer;
          });
          /**
           * @param {string} name
           * @param {?} v
           * @param {boolean} os
           * @return {undefined}
           */
          var set = function(name, v, os) {
            /** @type {string} */
            this.name = name;
            this.version = v;
            /** @type {boolean} */
            this.os = os;
            /** @type {string} */
            this.type = "browser";
          };
          /**
           * @param {?} tag
           * @return {undefined}
           */
          var start = function(tag) {
            this.version = tag;
            /** @type {string} */
            this.type = "node";
            /** @type {string} */
            this.name = "node";
            this.os = m.platform;
          };
          /**
           * @param {string} execResult
           * @param {?} reviver
           * @param {boolean} embed_tokens
           * @param {boolean} opt_ignoreMissingFields
           * @return {undefined}
           */
          var parse = function(execResult, reviver, embed_tokens, opt_ignoreMissingFields) {
            /** @type {string} */
            this.name = execResult;
            this.version = reviver;
            /** @type {boolean} */
            this.os = embed_tokens;
            /** @type {boolean} */
            this.bot = opt_ignoreMissingFields;
            /** @type {string} */
            this.type = "bot-device";
          };
          /**
           * @return {undefined}
           */
          var init = function() {
            /** @type {string} */
            this.type = "bot";
            /** @type {boolean} */
            this.bot = true;
            /** @type {string} */
            this.name = "bot";
            /** @type {null} */
            this.version = null;
            /** @type {null} */
            this.os = null;
          };
          /** @type {RegExp} */
          var re = /(nuhk|Googlebot|Yammybot|Openbot|Slurp|MSNBot|Ask\ Jeeves\/Teoma|ia_archiver)/;
          /** @type {Array} */
          var arr = [["aol", /AOLShield\/([0-9\._]+)/], ["edge", /Edge\/([0-9\._]+)/], ["edge-ios", /EdgiOS\/([0-9\._]+)/], ["yandexbrowser", /YaBrowser\/([0-9\._]+)/], ["vivaldi", /Vivaldi\/([0-9\.]+)/], ["kakaotalk", /KAKAOTALK\s([0-9\.]+)/], ["samsung", /SamsungBrowser\/([0-9\.]+)/], ["silk", /\bSilk\/([0-9._-]+)\b/], ["miui", /MiuiBrowser\/([0-9\.]+)$/], ["beaker", /BeakerBrowser\/([0-9\.]+)/], ["edge-chromium", /Edg\/([0-9\.]+)/], ["chromium-webview", /(?!Chrom.*OPR)wv\).*Chrom(?:e|ium)\/([0-9\.]+)(:?\s|$)/], 
          ["chrome", /(?!Chrom.*OPR)Chrom(?:e|ium)\/([0-9\.]+)(:?\s|$)/], ["phantomjs", /PhantomJS\/([0-9\.]+)(:?\s|$)/], ["crios", /CriOS\/([0-9\.]+)(:?\s|$)/], ["firefox", /Firefox\/([0-9\.]+)(?:\s|$)/], ["fxios", /FxiOS\/([0-9\.]+)/], ["opera-mini", /Opera Mini.*Version\/([0-9\.]+)/], ["opera", /Opera\/([0-9\.]+)(?:\s|$)/], ["opera", /OPR\/([0-9\.]+)(:?\s|$)/], ["ie", /Trident\/7\.0.*rv\:([0-9\.]+).*\).*Gecko$/], ["ie", /MSIE\s([0-9\.]+);.*Trident\/[4-7].0/], ["ie", /MSIE\s(7\.0)/], ["bb10", /BB10;\sTouch.*Version\/([0-9\.]+)/], 
          ["android", /Android\s([0-9\.]+)/], ["ios", /Version\/([0-9\._]+).*Mobile.*Safari.*/], ["safari", /Version\/([0-9\._]+).*Safari/], ["facebook", /FBAV\/([0-9\.]+)/], ["instagram", /Instagram\s([0-9\.]+)/], ["ios-webview", /AppleWebKit\/([0-9\.]+).*Mobile/], ["ios-webview", /AppleWebKit\/([0-9\.]+).*Gecko\)$/], ["searchbot", /alexa|bot|crawl(er|ing)|facebookexternalhit|feedburner|google web preview|nagios|postrank|pingdom|slurp|spider|yahoo!|yandex/]];
          /** @type {Array} */
          var resolveValues = [["iOS", /iP(hone|od|ad)/], ["Android OS", /Android/], ["BlackBerry OS", /BlackBerry|BB10/], ["Windows Mobile", /IEMobile/], ["Amazon OS", /Kindle/], ["Windows 3.11", /Win16/], ["Windows 95", /(Windows 95)|(Win95)|(Windows_95)/], ["Windows 98", /(Windows 98)|(Win98)/], ["Windows 2000", /(Windows NT 5.0)|(Windows 2000)/], ["Windows XP", /(Windows NT 5.1)|(Windows XP)/], ["Windows Server 2003", /(Windows NT 5.2)/], ["Windows Vista", /(Windows NT 6.0)/], ["Windows 7", /(Windows NT 6.1)/], 
          ["Windows 8", /(Windows NT 6.2)/], ["Windows 8.1", /(Windows NT 6.3)/], ["Windows 10", /(Windows NT 10.0)/], ["Windows ME", /Windows ME/], ["Open BSD", /OpenBSD/], ["Sun OS", /SunOS/], ["Chrome OS", /CrOS/], ["Linux", /(Linux)|(X11)/], ["Mac OS", /(Mac_PowerPC)|(Macintosh)/], ["QNX", /QNX/], ["BeOS", /BeOS/], ["OS/2", /OS\/2/]];
        }).call(this, f(125));
      }, function(module, dataAndEvents) {
        /**
         * @return {?}
         */
        function none() {
          throw new Error("setTimeout has not been defined");
        }
        /**
         * @return {?}
         */
        function cssFloat() {
          throw new Error("clearTimeout has not been defined");
        }
        /**
         * @param {Function} x
         * @return {?}
         */
        function get(x) {
          if (display === setTimeout) {
            return setTimeout(x, 0);
          }
          if ((display === none || !display) && setTimeout) {
            return display = setTimeout, setTimeout(x, 0);
          }
          try {
            return display(x, 0);
          } catch (e) {
            try {
              return display.call(null, x, 0);
            } catch (e) {
              return display.call(this, x, 0);
            }
          }
        }
        /**
         * @return {undefined}
         */
        function next() {
          if (f) {
            if (received) {
              /** @type {boolean} */
              f = false;
              if (received.length) {
                data = received.concat(data);
              } else {
                /** @type {number} */
                index = -1;
              }
              if (data.length) {
                start();
              }
            }
          }
        }
        /**
         * @return {undefined}
         */
        function start() {
          if (!f) {
            var on = get(next);
            /** @type {boolean} */
            f = true;
            var length = data.length;
            for (;length;) {
              received = data;
              /** @type {Array} */
              data = [];
              for (;++index < length;) {
                if (received) {
                  received[index].run();
                }
              }
              /** @type {number} */
              index = -1;
              /** @type {number} */
              length = data.length;
            }
            /** @type {null} */
            received = null;
            /** @type {boolean} */
            f = false;
            (function(id) {
              if (property === clearTimeout) {
                return clearTimeout(id);
              }
              if ((property === cssFloat || !property) && clearTimeout) {
                return property = clearTimeout, clearTimeout(id);
              }
              try {
                property(id);
              } catch (e) {
                try {
                  return property.call(null, id);
                } catch (e) {
                  return property.call(this, id);
                }
              }
            })(on);
          }
        }
        /**
         * @param {Function} options
         * @param {Object} array
         * @return {undefined}
         */
        function Animation(options, array) {
          /** @type {Function} */
          this.fun = options;
          /** @type {Object} */
          this.array = array;
        }
        /**
         * @return {undefined}
         */
        function tmp() {
        }
        var display;
        var property;
        var process = module.exports = {};
        !function() {
          try {
            /** @type {Function} */
            display = "function" == typeof setTimeout ? setTimeout : none;
          } catch (t) {
            /** @type {function (): ?} */
            display = none;
          }
          try {
            /** @type {Function} */
            property = "function" == typeof clearTimeout ? clearTimeout : cssFloat;
          } catch (t) {
            /** @type {function (): ?} */
            property = cssFloat;
          }
        }();
        var received;
        /** @type {Array} */
        var data = [];
        /** @type {boolean} */
        var f = false;
        /** @type {number} */
        var index = -1;
        /**
         * @param {string} callback
         * @return {undefined}
         */
        process.nextTick = function(callback) {
          /** @type {Array} */
          var x = new Array(arguments.length - 1);
          if (arguments.length > 1) {
            /** @type {number} */
            var i = 1;
            for (;i < arguments.length;i++) {
              x[i - 1] = arguments[i];
            }
          }
          data.push(new Animation(callback, x));
          if (!(1 !== data.length)) {
            if (!f) {
              get(start);
            }
          }
        };
        /**
         * @return {undefined}
         */
        Animation.prototype.run = function() {
          this.fun.apply(null, this.array);
        };
        /** @type {string} */
        process.title = "browser";
        /** @type {boolean} */
        process.browser = true;
        process.env = {};
        /** @type {Array} */
        process.argv = [];
        /** @type {string} */
        process.version = "";
        process.versions = {};
        /** @type {function (): undefined} */
        process.on = tmp;
        /** @type {function (): undefined} */
        process.addListener = tmp;
        /** @type {function (): undefined} */
        process.once = tmp;
        /** @type {function (): undefined} */
        process.off = tmp;
        /** @type {function (): undefined} */
        process.removeListener = tmp;
        /** @type {function (): undefined} */
        process.removeAllListeners = tmp;
        /** @type {function (): undefined} */
        process.emit = tmp;
        /** @type {function (): undefined} */
        process.prependListener = tmp;
        /** @type {function (): undefined} */
        process.prependOnceListener = tmp;
        /**
         * @param {?} type
         * @return {?}
         */
        process.listeners = function(type) {
          return[];
        };
        /**
         * @param {?} name
         * @return {?}
         */
        process.binding = function(name) {
          throw new Error("process.binding is not supported");
        };
        /**
         * @return {?}
         */
        process.cwd = function() {
          return "/";
        };
        /**
         * @param {?} dir
         * @return {?}
         */
        process.chdir = function(dir) {
          throw new Error("process.chdir is not supported");
        };
        /**
         * @return {?}
         */
        process.umask = function() {
          return 0;
        };
      }, function(dataAndEvents, expectedNumberOfNonCommentArgs, out) {
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Array} obj
         * @return {undefined}
         */
        function defineProperty(expectedNumberOfNonCommentArgs, obj) {
          /** @type {number} */
          var i = 0;
          for (;i < obj.length;i++) {
            var desc = obj[i];
            desc.enumerable = desc.enumerable || false;
            /** @type {boolean} */
            desc.configurable = true;
            if ("value" in desc) {
              /** @type {boolean} */
              desc.writable = true;
            }
            Object.defineProperty(expectedNumberOfNonCommentArgs, desc.key, desc);
          }
        }
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Array} codeSegments
         * @return {undefined}
         */
        function installProperty(expectedNumberOfNonCommentArgs, codeSegments) {
          /** @type {number} */
          var i = 0;
          for (;i < codeSegments.length;i++) {
            var desc = codeSegments[i];
            desc.enumerable = desc.enumerable || false;
            /** @type {boolean} */
            desc.configurable = true;
            if ("value" in desc) {
              /** @type {boolean} */
              desc.writable = true;
            }
            Object.defineProperty(expectedNumberOfNonCommentArgs, desc.key, desc);
          }
        }
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Array} obj
         * @return {undefined}
         */
        function deepClone(expectedNumberOfNonCommentArgs, obj) {
          /** @type {number} */
          var i = 0;
          for (;i < obj.length;i++) {
            var desc = obj[i];
            desc.enumerable = desc.enumerable || false;
            /** @type {boolean} */
            desc.configurable = true;
            if ("value" in desc) {
              /** @type {boolean} */
              desc.writable = true;
            }
            Object.defineProperty(expectedNumberOfNonCommentArgs, desc.key, desc);
          }
        }
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {(Array|NodeList)} obj
         * @return {undefined}
         */
        function iterator(expectedNumberOfNonCommentArgs, obj) {
          /** @type {number} */
          var i = 0;
          for (;i < obj.length;i++) {
            var desc = obj[i];
            desc.enumerable = desc.enumerable || false;
            /** @type {boolean} */
            desc.configurable = true;
            if ("value" in desc) {
              /** @type {boolean} */
              desc.writable = true;
            }
            Object.defineProperty(expectedNumberOfNonCommentArgs, desc.key, desc);
          }
        }
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Array} obj
         * @return {undefined}
         */
        function secureKey(expectedNumberOfNonCommentArgs, obj) {
          /** @type {number} */
          var i = 0;
          for (;i < obj.length;i++) {
            var desc = obj[i];
            desc.enumerable = desc.enumerable || false;
            /** @type {boolean} */
            desc.configurable = true;
            if ("value" in desc) {
              /** @type {boolean} */
              desc.writable = true;
            }
            Object.defineProperty(expectedNumberOfNonCommentArgs, desc.key, desc);
          }
        }
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {(Array|NodeList)} obj
         * @return {undefined}
         */
        function check(expectedNumberOfNonCommentArgs, obj) {
          /** @type {number} */
          var i = 0;
          for (;i < obj.length;i++) {
            var desc = obj[i];
            desc.enumerable = desc.enumerable || false;
            /** @type {boolean} */
            desc.configurable = true;
            if ("value" in desc) {
              /** @type {boolean} */
              desc.writable = true;
            }
            Object.defineProperty(expectedNumberOfNonCommentArgs, desc.key, desc);
          }
        }
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Array} obj
         * @return {undefined}
         */
        function def(expectedNumberOfNonCommentArgs, obj) {
          /** @type {number} */
          var i = 0;
          for (;i < obj.length;i++) {
            var desc = obj[i];
            desc.enumerable = desc.enumerable || false;
            /** @type {boolean} */
            desc.configurable = true;
            if ("value" in desc) {
              /** @type {boolean} */
              desc.writable = true;
            }
            Object.defineProperty(expectedNumberOfNonCommentArgs, desc.key, desc);
          }
        }
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Array} members
         * @return {undefined}
         */
        function initializeProperties(expectedNumberOfNonCommentArgs, members) {
          /** @type {number} */
          var i = 0;
          for (;i < members.length;i++) {
            var desc = members[i];
            desc.enumerable = desc.enumerable || false;
            /** @type {boolean} */
            desc.configurable = true;
            if ("value" in desc) {
              /** @type {boolean} */
              desc.writable = true;
            }
            Object.defineProperty(expectedNumberOfNonCommentArgs, desc.key, desc);
          }
        }
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {(Array|NodeList)} regex
         * @return {undefined}
         */
        function value(expectedNumberOfNonCommentArgs, regex) {
          /** @type {number} */
          var index = 0;
          for (;index < regex.length;index++) {
            var desc = regex[index];
            desc.enumerable = desc.enumerable || false;
            /** @type {boolean} */
            desc.configurable = true;
            if ("value" in desc) {
              /** @type {boolean} */
              desc.writable = true;
            }
            Object.defineProperty(expectedNumberOfNonCommentArgs, desc.key, desc);
          }
        }
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Array} source
         * @return {undefined}
         */
        function mixin(expectedNumberOfNonCommentArgs, source) {
          /** @type {number} */
          var i = 0;
          for (;i < source.length;i++) {
            var desc = source[i];
            desc.enumerable = desc.enumerable || false;
            /** @type {boolean} */
            desc.configurable = true;
            if ("value" in desc) {
              /** @type {boolean} */
              desc.writable = true;
            }
            Object.defineProperty(expectedNumberOfNonCommentArgs, desc.key, desc);
          }
        }
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {(Array|NodeList)} b
         * @return {undefined}
         */
        function extend(expectedNumberOfNonCommentArgs, b) {
          /** @type {number} */
          var bi = 0;
          for (;bi < b.length;bi++) {
            var desc = b[bi];
            desc.enumerable = desc.enumerable || false;
            /** @type {boolean} */
            desc.configurable = true;
            if ("value" in desc) {
              /** @type {boolean} */
              desc.writable = true;
            }
            Object.defineProperty(expectedNumberOfNonCommentArgs, desc.key, desc);
          }
        }
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {Array} obj
         * @return {undefined}
         */
        function setValue(expectedNumberOfNonCommentArgs, obj) {
          /** @type {number} */
          var i = 0;
          for (;i < obj.length;i++) {
            var desc = obj[i];
            desc.enumerable = desc.enumerable || false;
            /** @type {boolean} */
            desc.configurable = true;
            if ("value" in desc) {
              /** @type {boolean} */
              desc.writable = true;
            }
            Object.defineProperty(expectedNumberOfNonCommentArgs, desc.key, desc);
          }
        }
        out.r(expectedNumberOfNonCommentArgs);
        out(70);
        out(107);
        out(114);
        out(118);
        var content = out(2);
        var base = out.n(content);
        var Bn = out(1);
        var object = out(9);
        var opts = out(0);
        var Node = function() {
          /**
           * @param {?} y
           * @param {?} domEl
           * @param {number} iterations
           * @param {number} x
           * @return {undefined}
           */
          function loop(y, domEl, iterations, x) {
            if (function(dataAndEvents, loop) {
              if (!(dataAndEvents instanceof loop)) {
                throw new TypeError("Cannot call a class as a function");
              }
            }(this, loop), !y) {
              throw new Error("Verifier must be provided");
            }
            if (y.equals(new Bn.BigInteger("0"))) {
              throw new Error("Verifier integer must not equal 0");
            }
            if (!domEl) {
              throw new Error("Salt must be provided");
            }
            if (iterations <= 0) {
              throw new Error("Iterations must be greater than 0.");
            }
            this.verifier = y;
            this.salt = domEl;
            /** @type {number} */
            this.iterations = iterations;
            /** @type {number} */
            this.x = x;
          }
          var ctor;
          var suiteView;
          return ctor = loop, (suiteView = [{
            key : "isMatchingVerifier",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs) {
              if (!expectedNumberOfNonCommentArgs) {
                throw new Error("Please provide a verifier to match against.");
              }
              var _ = expectedNumberOfNonCommentArgs.getSalt();
              var codeSegments = this.salt.filter(function(haystack) {
                return!_.includes(haystack);
              });
              var resolveValues = this.verifier.filter(function(recurring) {
                return!expectedNumberOfNonCommentArgs.verifier.equals(recurring);
              });
              return 0 === codeSegments.length && 0 === resolveValues.length;
            }
          }, {
            key : "getX",
            /**
             * @return {?}
             */
            value : function() {
              return this.x;
            }
          }]) && defineProperty(ctor.prototype, suiteView), loop;
        }();
        var _ = object.a.SubtleCrypto;
        var crypto = object.a.WebCrypto;
        var PasswordManager = function() {
          /**
           * @param {?} domEl
           * @param {?} iterations
           * @return {undefined}
           */
          function loop(domEl, iterations) {
            !function(dataAndEvents, loop) {
              if (!(dataAndEvents instanceof loop)) {
                throw new TypeError("Cannot call a class as a function");
              }
            }(this, loop);
            if (domEl) {
              this.salt = domEl;
            }
            if (iterations) {
              this.iterations = iterations;
            }
            /** @type {null} */
            this.x = null;
          }
          var Type;
          var wrapperPrototype;
          return Type = loop, (wrapperPrototype = [{
            key : "generateVerifier",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Object} object
             * @param {Object} data
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object, data) {
              var win;
              var sig;
              return regeneratorRuntime.async(function(self) {
                for (;;) {
                  switch(self.prev = self.next) {
                    case 0:
                      if (expectedNumberOfNonCommentArgs) {
                        /** @type {number} */
                        self.next = 2;
                        break;
                      }
                      throw new Error("HMAC SHA-256 encoded username is required.");;
                    case 2:
                      if (object) {
                        /** @type {number} */
                        self.next = 4;
                        break;
                      }
                      throw new Error("Password must be provided.");;
                    case 4:
                      if (!(object.length < data.minPasswordLength)) {
                        /** @type {number} */
                        self.next = 6;
                        break;
                      }
                      throw new Error("The password length is below the minimum length requirement.");;
                    case 6:
                      if (!(object.length > data.maxPasswordLength)) {
                        /** @type {number} */
                        self.next = 8;
                        break;
                      }
                      throw new Error("The password length is above the maximum length requirement.");;
                    case 8:
                      if (data) {
                        /** @type {number} */
                        self.next = 10;
                        break;
                      }
                      throw new Error("Version must be provided.");;
                    case 10:
                      return this.iterations || (this.iterations = data.iterations), this.salt || (this.salt = this.generateSalt(data.saltSize)), self.next = 14, regeneratorRuntime.awrap(this.computeX(this.salt, expectedNumberOfNonCommentArgs, object, data));
                    case 14:
                      if (this.x = self.sent, win = this.computeVerifier(this.x, data.getParameters().g, data.getParameters().N), sig = win.toByteArray(), 1 !== this.x.signum() || 0 !== sig[0]) {
                        /** @type {number} */
                        self.next = 20;
                        break;
                      }
                      return sig.shift(), self.abrupt("return", new Node(new Bn.BigInteger(sig), opts.a.hexStringToArrayBufferBE(this.salt), this.iterations, this.x));
                    case 20:
                      return self.abrupt("return", new Node(win, opts.a.hexStringToArrayBufferBE(this.salt), this.iterations, this.x));
                    case 21:
                    ;
                    case "end":
                      return self.stop();
                  }
                }
              }, null, this);
            }
          }, {
            key : "computeX",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @param {boolean} data
             * @param {?} thisValue
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object, data, thisValue) {
              return regeneratorRuntime.async(function(current) {
                for (;;) {
                  switch(current.prev = current.next) {
                    case 0:
                      return current.abrupt("return", thisValue.xRoutineVersion.computeX(expectedNumberOfNonCommentArgs, object, data));
                    case 1:
                    ;
                    case "end":
                      return current.stop();
                  }
                }
              });
            }
          }, {
            key : "computeVerifier",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @param {number} value
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object, value) {
              return expectedNumberOfNonCommentArgs.signum() < 0 ? object.modPow(expectedNumberOfNonCommentArgs.negate(), value).modInverse(value) : object.modPow(expectedNumberOfNonCommentArgs, value);
            }
          }, {
            key : "generateSalt",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs) {
              if (expectedNumberOfNonCommentArgs <= 0) {
                throw new Error("Salt size must be greater than 0");
              }
              /** @type {Uint8Array} */
              var pool = new Uint8Array(expectedNumberOfNonCommentArgs);
              return crypto.getRandomValues(pool), opts.a.arrayBufferToHexString(pool);
            }
          }, {
            key : "getSrpUsername",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object) {
              var restoreScript;
              var height;
              var width;
              return regeneratorRuntime.async(function(self) {
                for (;;) {
                  switch(self.prev = self.next) {
                    case 0:
                      if (!(expectedNumberOfNonCommentArgs <= 0)) {
                        /** @type {number} */
                        self.next = 2;
                        break;
                      }
                      throw new Error("Invalid accountId");;
                    case 2:
                      return self.prev = 2, restoreScript = opts.a.bigNumbertoBufferLE(expectedNumberOfNonCommentArgs, 8), self.next = 6, regeneratorRuntime.awrap(_.importKey("raw", object, {
                        name : "HMAC",
                        hash : "SHA-256"
                      }, false, ["sign", "verify"]));
                    case 6:
                      return height = self.sent, self.next = 9, regeneratorRuntime.awrap(_.sign({
                        name : "HMAC",
                        hash : "SHA-256"
                      }, height, restoreScript));
                    case 9:
                      return width = self.sent, self.abrupt("return", opts.a.arrayBufferToHexString(new Uint8Array(width)).toUpperCase());
                    case 13:
                      throw self.prev = 13, self.t0 = self.catch(2), new Error("Unable to generate the SRP username", self.t0);;
                    case 16:
                    ;
                    case "end":
                      return self.stop();
                  }
                }
              }, null, null, [[2, 13]]);
            }
          }, {
            key : "encodeVerifier",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object) {
              return object.encodeVerifier(expectedNumberOfNonCommentArgs);
            }
          }, {
            key : "encodeVerifierHex",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object) {
              var table = this.encodeVerifier(expectedNumberOfNonCommentArgs, object);
              return opts.a.arrayBufferToHexString(new Uint8Array(table)).toUpperCase();
            }
          }, {
            key : "decodeVerifier",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object) {
              return object.decodeVerifier(expectedNumberOfNonCommentArgs);
            }
          }]) && installProperty(Type.prototype, wrapperPrototype), loop;
        }();
        var callback = object.a.SubtleCrypto;
        var Tokenator = function() {
          /**
           * @param {Blob} ctxt
           * @return {undefined}
           */
          function core(ctxt) {
            !function(dataAndEvents, core) {
              if (!(dataAndEvents instanceof core)) {
                throw new TypeError("Cannot call a class as a function");
              }
            }(this, core);
            /** @type {Blob} */
            this.saltSize = ctxt;
            /** @type {string} */
            this.digest = "SHA-256";
          }
          var me;
          var suiteView;
          return me = core, (suiteView = [{
            key : "computeX",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Object} object
             * @param {string} value
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object, value) {
              var nodes;
              var array;
              var letter;
              var _args;
              var args;
              var buffer;
              var data;
              var bytes;
              var result;
              var width;
              return regeneratorRuntime.async(function(self) {
                for (;;) {
                  switch(self.prev = self.next) {
                    case 0:
                      return(nodes = opts.a.hexStringToArrayBufferBE(expectedNumberOfNonCommentArgs)).length < this.saltSize && (array = new Uint8Array(this.saltSize), nodes = array.set(nodes)), letter = value.substring(0, 16), _args = "".concat(object, ":").concat(letter.toUpperCase()), args = opts.a.stringToArrayBuffer(_args), self.next = 7, regeneratorRuntime.awrap(callback.digest(this.digest, args));
                    case 7:
                      return buffer = self.sent, data = new Uint8Array(buffer), bytes = nodes.length + data.length, (result = new Uint8Array(bytes)).set(nodes), result.set(data, nodes.length), self.next = 15, regeneratorRuntime.awrap(callback.digest(this.digest, result));
                    case 15:
                      return width = self.sent, self.abrupt("return", opts.a.bufferToBigNumberLE(new Uint8Array(width)));
                    case 17:
                    ;
                    case "end":
                      return self.stop();
                  }
                }
              }, null, this);
            }
          }]) && deepClone(me.prototype, suiteView), core;
        }();
        var PasswordVersion1 = function() {
          /**
           * @param {(number|string)} iterations
           * @return {undefined}
           */
          function core(iterations) {
            !function(dataAndEvents, core) {
              if (!(dataAndEvents instanceof core)) {
                throw new TypeError("Cannot call a class as a function");
              }
            }(this, core);
            this.iterations = iterations || 1;
            /** @type {number} */
            this.saltSize = 32;
            /** @type {number} */
            this.verifierSize = 128;
            /** @type {number} */
            this.minPasswordLength = 8;
            /** @type {number} */
            this.maxPasswordLength = 16;
            this.xRoutineVersion = new Tokenator(this.saltSize);
          }
          var x;
          var suiteView;
          var a;
          return x = core, a = [{
            key : "build",
            /**
             * @return {?}
             */
            value : function() {
              return new core;
            }
          }], (suiteView = [{
            key : "getParameters",
            /**
             * @return {?}
             */
            value : function() {
              return{
                N : new Bn.BigInteger('94558736629309251206436488916623864910444695865064772352148093707798675228170106115630190094901096401883540229236016599430725894430734991444298272129143681820273859470730877741629279425748927230996376833577406570089078823475120723855492588316592686203439138514838131581023312004481906611790561347740748686507"'),
                g : new Bn.BigInteger("2"),
                H : "SHA-256"
              };
            }
          }, {
            key : "getEncodedVerifierSize",
            /**
             * @return {?}
             */
            value : function() {
              return this.saltSize + this.verifierSize;
            }
          }, {
            key : "encodeVerifier",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs) {
              /** @type {Uint8Array} */
              var s = new Uint8Array(this.getEncodedVerifierSize());
              s.set(expectedNumberOfNonCommentArgs.salt, 0);
              var nodes = opts.a.bigNumbertoBufferBE(expectedNumberOfNonCommentArgs.verifier, this.verifierSize);
              return s.set(nodes, this.saltSize), s;
            }
          }, {
            key : "decodeVerifier",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs) {
              if (!expectedNumberOfNonCommentArgs) {
                throw new Error("Bytes must be provided.");
              }
              if (expectedNumberOfNonCommentArgs.length !== this.getEncodedVerifierSize()) {
                throw new Error("Invalid verifier length.");
              }
              var id = expectedNumberOfNonCommentArgs.slice(0, this.saltSize);
              var name = opts.a.bufferToBigNumberBE(expectedNumberOfNonCommentArgs.slice(this.saltSize, this.saltSize + this.verifierSize));
              return new Node(name, id, this.iterations);
            }
          }]) && iterator(x.prototype, suiteView), a && iterator(x, a), core;
        }();
        var $ = object.a.SubtleCrypto;
        var iterations = function() {
          /**
           * @param {number} iterations
           * @param {Blob} ctxt
           * @return {undefined}
           */
          function core(iterations, ctxt) {
            !function(dataAndEvents, core) {
              if (!(dataAndEvents instanceof core)) {
                throw new TypeError("Cannot call a class as a function");
              }
            }(this, core);
            /** @type {number} */
            this.iterations = iterations;
            /** @type {Blob} */
            this.saltSize = ctxt;
            /** @type {string} */
            this.digest = "SHA-512";
            /** @type {number} */
            this.keyLength = 512;
          }
          var component;
          var suiteView;
          return component = core, (suiteView = [{
            key : "computeX",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Object} object
             * @param {string} value
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object, value) {
              var nodes;
              var array;
              var rest;
              var data;
              var props;
              var settings;
              var width;
              return regeneratorRuntime.async(function(self) {
                for (;;) {
                  switch(self.prev = self.next) {
                    case 0:
                      return(nodes = opts.a.hexStringToArrayBufferBE(expectedNumberOfNonCommentArgs)).length < this.saltSize && (array = new Uint8Array(this.saltSize), nodes = array.set(nodes)), rest = value.substring(0, 128), data = "".concat(object, ":").concat(rest), self.next = 6, regeneratorRuntime.awrap($.importKey("raw", opts.a.stringToArrayBuffer(data), {
                        name : "PBKDF2"
                      }, false, ["deriveBits", "deriveKey"]));
                    case 6:
                      return props = self.sent, self.next = 9, regeneratorRuntime.awrap($.deriveKey({
                        name : "PBKDF2",
                        salt : nodes,
                        iterations : this.iterations,
                        hash : this.digest
                      }, props, {
                        name : "HMAC",
                        hash : this.digest,
                        length : this.keyLength
                      }, true, ["sign", "verify"]));
                    case 9:
                      return settings = self.sent, self.next = 12, regeneratorRuntime.awrap($.exportKey("raw", settings));
                    case 12:
                      return width = self.sent, self.abrupt("return", new Bn.BigInteger(new Uint8Array(width)));
                    case 14:
                    ;
                    case "end":
                      return self.stop();
                  }
                }
              }, null, this);
            }
          }]) && secureKey(component.prototype, suiteView), core;
        }();
        var PasswordVersion2 = function() {
          /**
           * @param {number} ctxt
           * @return {undefined}
           */
          function core(ctxt) {
            !function(dataAndEvents, core) {
              if (!(dataAndEvents instanceof core)) {
                throw new TypeError("Cannot call a class as a function");
              }
            }(this, core);
            this.iterations = ctxt || 15E3;
            /** @type {number} */
            this.saltSize = 32;
            /** @type {number} */
            this.verifierSize = 256;
            /** @type {number} */
            this.iterationsSize = 4;
            /** @type {number} */
            this.minPasswordLength = 8;
            /** @type {number} */
            this.maxPasswordLength = 128;
            this.xRoutineVersion = new iterations(this.iterations, this.saltSize);
          }
          var a;
          var suiteView;
          var b;
          return a = core, b = [{
            key : "build",
            /**
             * @return {?}
             */
            value : function() {
              return new core;
            }
          }], (suiteView = [{
            key : "getParameters",
            /**
             * @return {?}
             */
            value : function() {
              return{
                N : new Bn.BigInteger("21766174458617435773191008891802753781907668374255538511144643224689886235383840957210909013086056401571399717235807266581649606472148410291413364152197364477180887395655483738115072677402235101762521901569820740293149529620419333266262073471054548368736039519702486226506248861060256971802984953561121442680157668000761429988222457090413873973970171927093992114751765168063614761119615476233422096442783117971236371647333871414335895773474667308967050807005509320424799678417036867928316761272274230314067548291133582479583061439577559347101961771406173684378522703483495337037655006751328447510550299250924469288819"),
                g : new Bn.BigInteger("2"),
                H : "SHA-256"
              };
            }
          }, {
            key : "getEncodedVerifierSize",
            /**
             * @return {?}
             */
            value : function() {
              return this.saltSize + this.verifierSize + this.iterationsSize;
            }
          }, {
            key : "encodeVerifier",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs) {
              /** @type {Uint8Array} */
              var t = new Uint8Array(this.getEncodedVerifierSize());
              t.set(expectedNumberOfNonCommentArgs.salt, 0);
              var nodes = opts.a.bigNumbertoBufferBE(expectedNumberOfNonCommentArgs.verifier, this.verifierSize);
              t.set(nodes, this.saltSize);
              var results = opts.a.bigNumbertoBufferBE(new Bn.BigInteger(expectedNumberOfNonCommentArgs.iterations.toString()), this.iterationsSize);
              return t.set(results, this.saltSize + this.verifierSize), t;
            }
          }, {
            key : "decodeVerifier",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs) {
              if (!expectedNumberOfNonCommentArgs) {
                throw new Error("Bytes must be provided.");
              }
              if (expectedNumberOfNonCommentArgs.length !== this.getEncodedVerifierSize()) {
                throw new Error("Invalid verifier length.");
              }
              var id = expectedNumberOfNonCommentArgs.slice(0, this.saltSize);
              var name = opts.a.bufferToBigNumberBE(expectedNumberOfNonCommentArgs.slice(this.saltSize, this.saltSize + this.verifierSize));
              var key = this.saltSize + this.verifierSize;
              var camelKey = expectedNumberOfNonCommentArgs.slice(key, key + this.iterationsSize);
              var expr = opts.a.byteArrayToInteger(camelKey);
              return new Node(name, id, expr);
            }
          }]) && check(a.prototype, suiteView), b && check(a, b), core;
        }();
        var B = function() {
          /**
           * @param {number} domEl
           * @param {number} iterations
           * @return {undefined}
           */
          function loop(domEl, iterations) {
            !function(dataAndEvents, loop) {
              if (!(dataAndEvents instanceof loop)) {
                throw new TypeError("Cannot call a class as a function");
              }
            }(this, loop);
            if (domEl) {
              /** @type {number} */
              this.salt = domEl;
            }
            if (iterations) {
              /** @type {number} */
              this.iterations = iterations;
            }
            /** @type {null} */
            this.x = null;
          }
          var Type;
          var suiteView;
          return Type = loop, (suiteView = [{
            key : "generateVerifier",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Object} object
             * @param {Object} data
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object, data) {
              var me = this;
              if (!expectedNumberOfNonCommentArgs) {
                throw new Error("HMAC SHA-256 encoded username is required.");
              }
              if (!object) {
                throw new Error("Password must be provided.");
              }
              if (object.length < data.minPasswordLength) {
                throw new Error("The password length is below the minimum length requirement.");
              }
              if (object.length > data.maxPasswordLength) {
                throw new Error("The password length is above the maximum length requirement.");
              }
              if (!data) {
                throw new Error("Version must be provided.");
              }
              if (!this.iterations) {
                this.iterations = data.iterations;
              }
              if (!this.salt) {
                this.salt = this.generateSalt(data.saltSize);
              }
              this.x = this.computeX(this.salt, expectedNumberOfNonCommentArgs, object, data);
              var x = this.computeVerifier(this.x, data.getParameters().g, data.getParameters().N);
              return new Promise(function(it, on) {
                try {
                  var sig = x.toByteArray();
                  return 1 === me.x.signum() && 0 === sig[0] ? (sig.shift(), it(new Node(new Bn.BigInteger(sig), opts.a.hexStringToArrayBufferBE(me.salt), me.iterations, me.x))) : it(new Node(x, opts.a.hexStringToArrayBufferBE(me.salt), me.iterations, me.x));
                } catch (failuresLink) {
                  return on(failuresLink);
                }
              });
            }
          }, {
            key : "computeX",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @param {boolean} data
             * @param {?} thisValue
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object, data, thisValue) {
              return thisValue.xRoutineVersion.computeX(expectedNumberOfNonCommentArgs, object, data);
            }
          }, {
            key : "computeVerifier",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @param {number} value
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object, value) {
              return expectedNumberOfNonCommentArgs.signum() < 0 ? object.modPow(expectedNumberOfNonCommentArgs.negate(), value).modInverse(value) : object.modPow(expectedNumberOfNonCommentArgs, value);
            }
          }, {
            key : "generateSalt",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs) {
              if (expectedNumberOfNonCommentArgs <= 0) {
                throw new Error("Salt size must be greater than 0");
              }
              return base.a.codec.hex.fromBits(base.a.random.randomWords(expectedNumberOfNonCommentArgs / 2, 0));
            }
          }, {
            key : "getSrpUsername",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object) {
              return new Promise(function(fn, done) {
                if (expectedNumberOfNonCommentArgs <= 0) {
                  return done(new Error("Invalid accountId"));
                }
                try {
                  var array = opts.a.bigNumbertoBufferLE(expectedNumberOfNonCommentArgs, 8);
                  var mac = base.a.codec.hex.toBits(opts.a.arrayBufferToHexString(new Uint8Array(object)));
                  var pdataCur = (new base.a.misc.hmac(mac, base.a.hash.sha256)).mac(base.a.codec.hex.toBits(opts.a.arrayBufferToHexString(new Uint8Array(array))));
                  return fn(base.a.codec.hex.fromBits(pdataCur).toUpperCase());
                } catch (err) {
                  return done(new Error("Unable to generate the SRP username", err));
                }
              });
            }
          }, {
            key : "encodeVerifier",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object) {
              return object.encodeVerifier(expectedNumberOfNonCommentArgs);
            }
          }, {
            key : "encodeVerifierHex",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object) {
              var table = this.encodeVerifier(expectedNumberOfNonCommentArgs, object);
              return opts.a.arrayBufferToHexString(new Uint8Array(table)).toUpperCase();
            }
          }, {
            key : "decodeVerifier",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object) {
              return object.decodeVerifier(expectedNumberOfNonCommentArgs);
            }
          }]) && def(Type.prototype, suiteView), loop;
        }();
        var saltSize = function() {
          /**
           * @param {Blob} ctxt
           * @return {undefined}
           */
          function core(ctxt) {
            !function(dataAndEvents, core) {
              if (!(dataAndEvents instanceof core)) {
                throw new TypeError("Cannot call a class as a function");
              }
            }(this, core);
            /** @type {Blob} */
            this.saltSize = ctxt;
          }
          var constructor;
          var members;
          return constructor = core, (members = [{
            key : "computeX",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Object} object
             * @param {string} value
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object, value) {
              var letter = value.substring(0, 16);
              var hash = this.hash("".concat(object, ":").concat(letter.toUpperCase()));
              var node = this.hash(expectedNumberOfNonCommentArgs + hash, true);
              /** @type {string} */
              var sig = "";
              /** @type {number} */
              var offset = 0;
              for (;offset < node.length;offset += 2) {
                sig += node[node.length - offset - 2] + node[node.length - offset - 1];
              }
              return new Bn.BigInteger(sig, 16);
            }
          }, {
            key : "hash",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object) {
              /** @type {number} */
              var pdataCur = expectedNumberOfNonCommentArgs;
              if (object) {
                pdataCur = base.a.codec.hex.toBits(expectedNumberOfNonCommentArgs);
              }
              var r20 = base.a.codec.hex.fromBits(base.a.hash.sha256.hash(pdataCur));
              return this.pad(64, r20);
            }
          }, {
            key : "pad",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object) {
              /** @type {string} */
              var zeros = Array(expectedNumberOfNonCommentArgs + 1).join("0");
              return(zeros + object).slice(-zeros.length);
            }
          }]) && initializeProperties(constructor.prototype, members), core;
        }();
        var M = function() {
          /**
           * @param {(number|string)} iterations
           * @return {undefined}
           */
          function core(iterations) {
            !function(dataAndEvents, core) {
              if (!(dataAndEvents instanceof core)) {
                throw new TypeError("Cannot call a class as a function");
              }
            }(this, core);
            this.iterations = iterations || 1;
            /** @type {number} */
            this.saltSize = 32;
            /** @type {number} */
            this.verifierSize = 128;
            /** @type {number} */
            this.minPasswordLength = 8;
            /** @type {number} */
            this.maxPasswordLength = 16;
            this.xRoutineVersion = new saltSize(this.saltSize);
          }
          var context;
          var r20;
          var tag;
          return context = core, tag = [{
            key : "build",
            /**
             * @return {?}
             */
            value : function() {
              return new core;
            }
          }], (r20 = [{
            key : "getParameters",
            /**
             * @return {?}
             */
            value : function() {
              return{
                N : new Bn.BigInteger('94558736629309251206436488916623864910444695865064772352148093707798675228170106115630190094901096401883540229236016599430725894430734991444298272129143681820273859470730877741629279425748927230996376833577406570089078823475120723855492588316592686203439138514838131581023312004481906611790561347740748686507"'),
                g : new Bn.BigInteger("2"),
                H : "SHA-256"
              };
            }
          }, {
            key : "getEncodedVerifierSize",
            /**
             * @return {?}
             */
            value : function() {
              return this.saltSize + this.verifierSize;
            }
          }, {
            key : "encodeVerifier",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs) {
              /** @type {Uint8Array} */
              var s = new Uint8Array(this.getEncodedVerifierSize());
              s.set(expectedNumberOfNonCommentArgs.salt, 0);
              var nodes = opts.a.bigNumbertoBufferBE(expectedNumberOfNonCommentArgs.verifier, this.verifierSize);
              return s.set(nodes, this.saltSize), s;
            }
          }, {
            key : "decodeVerifier",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs) {
              if (!expectedNumberOfNonCommentArgs) {
                throw new Error("Bytes must be provided.");
              }
              if (expectedNumberOfNonCommentArgs.length !== this.getEncodedVerifierSize()) {
                throw new Error("Invalid verifier length.");
              }
              var id = expectedNumberOfNonCommentArgs.slice(0, this.saltSize);
              var name = opts.a.bufferToBigNumberBE(expectedNumberOfNonCommentArgs.slice(this.saltSize, this.saltSize + this.verifierSize));
              return new Node(name, id, this.iterations);
            }
          }]) && value(context.prototype, r20), tag && value(context, tag), core;
        }();
        var Handleger = function() {
          /**
           * @param {number} iterations
           * @param {Blob} ctxt
           * @return {undefined}
           */
          function core(iterations, ctxt) {
            !function(dataAndEvents, core) {
              if (!(dataAndEvents instanceof core)) {
                throw new TypeError("Cannot call a class as a function");
              }
            }(this, core);
            /** @type {number} */
            this.iterations = iterations;
            /** @type {Blob} */
            this.saltSize = ctxt;
            /** @type {string} */
            this.digest = "SHA-512";
            /** @type {number} */
            this.keyLength = 512;
          }
          var Text;
          var value;
          return Text = core, (value = [{
            key : "computeX",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Object} object
             * @param {string} value
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object, value) {
              var rest = value.substring(0, 128);
              /** @type {string} */
              var key = "".concat(object, ":").concat(rest);
              var SALT = base.a.codec.hex.toBits(expectedNumberOfNonCommentArgs);
              var camelKey = base.a.misc.pbkdf2(key, SALT, this.iterations, this.keyLength, function(dataAndEvents) {
                var newArgs = new base.a.misc.hmac(dataAndEvents, base.a.hash.sha512);
                /**
                 * @return {?}
                 */
                this.encrypt = function() {
                  return newArgs.encrypt.apply(newArgs, arguments);
                };
              });
              var r20 = base.a.codec.hex.fromBits(camelKey);
              return new Bn.BigInteger(new Uint8Array(opts.a.hexStringToArrayBufferBE(r20)));
            }
          }]) && mixin(Text.prototype, value), core;
        }();
        var R = function() {
          /**
           * @param {number} iterations
           * @return {undefined}
           */
          function core(iterations) {
            !function(dataAndEvents, core) {
              if (!(dataAndEvents instanceof core)) {
                throw new TypeError("Cannot call a class as a function");
              }
            }(this, core);
            this.iterations = iterations || 15E3;
            /** @type {number} */
            this.saltSize = 32;
            /** @type {number} */
            this.verifierSize = 256;
            /** @type {number} */
            this.iterationsSize = 4;
            /** @type {number} */
            this.minPasswordLength = 8;
            /** @type {number} */
            this.maxPasswordLength = 128;
            this.xRoutineVersion = new Handleger(this.iterations, this.saltSize);
          }
          var Benchmark;
          var QUnit;
          var oldconfig;
          return Benchmark = core, oldconfig = [{
            key : "build",
            /**
             * @return {?}
             */
            value : function() {
              return new core;
            }
          }], (QUnit = [{
            key : "getParameters",
            /**
             * @return {?}
             */
            value : function() {
              return{
                N : new Bn.BigInteger("21766174458617435773191008891802753781907668374255538511144643224689886235383840957210909013086056401571399717235807266581649606472148410291413364152197364477180887395655483738115072677402235101762521901569820740293149529620419333266262073471054548368736039519702486226506248861060256971802984953561121442680157668000761429988222457090413873973970171927093992114751765168063614761119615476233422096442783117971236371647333871414335895773474667308967050807005509320424799678417036867928316761272274230314067548291133582479583061439577559347101961771406173684378522703483495337037655006751328447510550299250924469288819"),
                g : new Bn.BigInteger("2"),
                H : "SHA-256"
              };
            }
          }, {
            key : "getEncodedVerifierSize",
            /**
             * @return {?}
             */
            value : function() {
              return this.saltSize + this.verifierSize + this.iterationsSize;
            }
          }, {
            key : "encodeVerifier",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs) {
              /** @type {Array} */
              var _ = [];
              /** @type {Array} */
              _ = _.concat(expectedNumberOfNonCommentArgs.salt);
              var args = opts.a.bigNumbertoBufferBE(expectedNumberOfNonCommentArgs.verifier);
              /** @type {Array} */
              _ = _.concat(Array.from(args));
              var a = opts.a.bigNumbertoBufferBE(new Bn.BigInteger(expectedNumberOfNonCommentArgs.iterations.toString()), this.iterationsSize);
              return _.concat(Array.from(a));
            }
          }, {
            key : "decodeVerifier",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs) {
              if (!expectedNumberOfNonCommentArgs) {
                throw new Error("Bytes must be provided.");
              }
              if (expectedNumberOfNonCommentArgs.length !== this.getEncodedVerifierSize()) {
                throw new Error("Invalid verifier length.");
              }
              var id = expectedNumberOfNonCommentArgs.split(0, this.saltSize);
              var name = opts.a.bufferToBigNumberBE(expectedNumberOfNonCommentArgs.split(0, this.verifierSize));
              var exprKey = expectedNumberOfNonCommentArgs.split(0, this.iterationsSize);
              var expr = opts.a.byteArrayToInteger(exprKey);
              return new Node(name, id, expr);
            }
          }]) && extend(Benchmark.prototype, QUnit), oldconfig && extend(Benchmark, oldconfig), core;
        }();
        var ret = function() {
          /**
           * @return {undefined}
           */
          function core() {
            !function(dataAndEvents, core) {
              if (!(dataAndEvents instanceof core)) {
                throw new TypeError("Cannot call a class as a function");
              }
            }(this, core);
          }
          var ctxStack;
          var suiteView;
          return ctxStack = core, (suiteView = [{
            key : "getBrowserExports",
            /**
             * @param {Object} expectedNumberOfNonCommentArgs
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs) {
              switch(expectedNumberOfNonCommentArgs && expectedNumberOfNonCommentArgs.name) {
                case "chrome":
                ;
                case "firefox":
                ;
                case "opera":
                  return this[expectedNumberOfNonCommentArgs.name + "Exports"](expectedNumberOfNonCommentArgs.version);
                default:
                  return this.legacyExports();
              }
            }
          }, {
            key : "chromeExports",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs) {
              var getMajorVersion = this.getMajorVersion(expectedNumberOfNonCommentArgs);
              return null != getMajorVersion && getMajorVersion <= 39 ? this.legacyExports() : this.modernExports();
            }
          }, {
            key : "firefoxExports",
            /**
             * @return {?}
             */
            value : function() {
              return this.modernExports();
            }
          }, {
            key : "operaExports",
            /**
             * @return {?}
             */
            value : function() {
              return this.modernExports();
            }
          }, {
            key : "modernExports",
            /**
             * @return {?}
             */
            value : function() {
              return{
                PasswordManager : PasswordManager,
                PasswordVersion1 : PasswordVersion1,
                PasswordVersion2 : PasswordVersion2,
                XRoutineVersion1 : Tokenator,
                XRoutineVersion2 : iterations
              };
            }
          }, {
            key : "legacyExports",
            /**
             * @return {?}
             */
            value : function() {
              return{
                PasswordManager : B,
                PasswordVersion1 : M,
                PasswordVersion2 : R,
                XRoutineVersion1 : saltSize,
                XRoutineVersion2 : Handleger
              };
            }
          }, {
            key : "getMajorVersion",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs) {
              try {
                if (!expectedNumberOfNonCommentArgs) {
                  return null;
                }
                /** @type {(Array.<string>|null)} */
                var e = /^(([0-9]+)\.([0-9]+)\.([0-9]+)(?:-([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?)(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$/gi.exec(expectedNumberOfNonCommentArgs);
                if (e) {
                  /** @type {string} */
                  var cDigit = e[2];
                  return parseInt(cDigit);
                }
                return null;
              } catch (t) {
                return null;
              }
            }
          }]) && setValue(ctxStack, suiteView), core;
        }();
        out.d(expectedNumberOfNonCommentArgs, "PasswordManager", function() {
          return i;
        });
        out.d(expectedNumberOfNonCommentArgs, "PasswordVersion1", function() {
          return t;
        });
        out.d(expectedNumberOfNonCommentArgs, "PasswordVersion2", function() {
          return handler;
        });
        out.d(expectedNumberOfNonCommentArgs, "XRoutineVersion1", function() {
          return l;
        });
        out.d(expectedNumberOfNonCommentArgs, "XRoutineVersion2", function() {
          return handleObj;
        });
        out.d(expectedNumberOfNonCommentArgs, "Crypto", function() {
          return object.a;
        });
        out.d(expectedNumberOfNonCommentArgs, "Utilities", function() {
          return opts.a;
        });
        out.d(expectedNumberOfNonCommentArgs, "sjcl", function() {
          return base.a;
        });
        var rreturn = (0, out(124).detect)();
        /** @type {null} */
        var i = null;
        /** @type {null} */
        var t = null;
        /** @type {null} */
        var handler = null;
        /** @type {null} */
        var l = null;
        /** @type {null} */
        var handleObj = null;
        var types = ret.getBrowserExports(rreturn);
        i = types.PasswordManager;
        t = types.PasswordVersion1;
        handler = types.PasswordVersion2;
        l = types.XRoutineVersion1;
        handleObj = types.XRoutineVersion2;
      }]);
    }, function(module, dataAndEvents, require) {
      var nodes = require(1);
      var flag = require(40);
      var inspect = require(5);
      var helper = require(44);
      var a = require(45);
      var Block = require(66);
      var obj = flag("wks");
      var x = nodes.Symbol;
      var recurse = Block ? x : x && x.withoutSetter || helper;
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs) {
        return inspect(obj, expectedNumberOfNonCommentArgs) || (a && inspect(x, expectedNumberOfNonCommentArgs) ? obj[expectedNumberOfNonCommentArgs] = x[expectedNumberOfNonCommentArgs] : obj[expectedNumberOfNonCommentArgs] = recurse("Symbol." + expectedNumberOfNonCommentArgs)), obj[expectedNumberOfNonCommentArgs];
      };
    }, function(module, dataAndEvents, require) {
      var getActual = require(11);
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs) {
        if (!getActual(expectedNumberOfNonCommentArgs)) {
          throw TypeError(String(expectedNumberOfNonCommentArgs) + " is not an object");
        }
        return expectedNumberOfNonCommentArgs;
      };
    }, function(module, dataAndEvents) {
      /** @type {function (this:Object, *): boolean} */
      var has = {}.hasOwnProperty;
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {Function} proto
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, proto) {
        return has.call(expectedNumberOfNonCommentArgs, proto);
      };
    }, function(module, dataAndEvents, require) {
      var Block = require(9);
      var object = require(10);
      var group = require(21);
      /** @type {function (number, Object, boolean): ?} */
      module.exports = Block ? function(expectedNumberOfNonCommentArgs, methodName, value) {
        return object.f(expectedNumberOfNonCommentArgs, methodName, group(1, value));
      } : function(expectedNumberOfNonCommentArgs, proto, value) {
        return expectedNumberOfNonCommentArgs[proto] = value, expectedNumberOfNonCommentArgs;
      };
    }, function(module, dataAndEvents) {
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs) {
        try {
          return!!expectedNumberOfNonCommentArgs();
        } catch (t) {
          return true;
        }
      };
    }, function(module, dataAndEvents, require) {
      var names = require(1);
      var inspect = require(32).f;
      var assert = require(6);
      var callback = require(13);
      var capitalize = require(27);
      var swap = require(73);
      var getActual = require(50);
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {Object} object
       * @return {undefined}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, object) {
        var obj;
        var key;
        var right;
        var value;
        var elem;
        var name = expectedNumberOfNonCommentArgs.target;
        var global = expectedNumberOfNonCommentArgs.global;
        var hasExt = expectedNumberOfNonCommentArgs.stat;
        if (obj = global ? names : hasExt ? names[name] || capitalize(name, {}) : (names[name] || {}).prototype) {
          for (key in object) {
            if (value = object[key], right = expectedNumberOfNonCommentArgs.noTargetGet ? (elem = inspect(obj, key)) && elem.value : obj[key], !getActual(global ? key : name + (hasExt ? "." : "#") + key, expectedNumberOfNonCommentArgs.forced) && void 0 !== right) {
              if (typeof value == typeof right) {
                continue;
              }
              swap(value, right);
            }
            if (expectedNumberOfNonCommentArgs.sham || right && right.sham) {
              assert(value, "sham", true);
            }
            callback(obj, key, value, expectedNumberOfNonCommentArgs);
          }
        }
      };
    }, function(module, dataAndEvents, require) {
      var getActual = require(7);
      /** @type {boolean} */
      module.exports = !getActual(function() {
        return 7 != Object.defineProperty({}, 1, {
          /**
           * @return {?}
           */
          get : function() {
            return 7;
          }
        })[1];
      });
    }, function(dataAndEvents, state, require) {
      var url = require(9);
      var Block = require(42);
      var getName = require(4);
      var expect = require(43);
      /** @type {function (Object, string, Object): Object} */
      var defineProperty = Object.defineProperty;
      /** @type {Function} */
      state.f = url ? defineProperty : function(expectedNumberOfNonCommentArgs, name, desc) {
        if (getName(expectedNumberOfNonCommentArgs), name = expect(name, true), getName(desc), Block) {
          try {
            return defineProperty(expectedNumberOfNonCommentArgs, name, desc);
          } catch (t) {
          }
        }
        if ("get" in desc || "set" in desc) {
          throw TypeError("Accessors not supported");
        }
        return "value" in desc && (expectedNumberOfNonCommentArgs[name] = desc.value), expectedNumberOfNonCommentArgs;
      };
    }, function(module, dataAndEvents) {
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs) {
        return "object" == typeof expectedNumberOfNonCommentArgs ? null !== expectedNumberOfNonCommentArgs : "function" == typeof expectedNumberOfNonCommentArgs;
      };
    }, function(module, dataAndEvents, expression) {
      var obj = expression(48);
      var modifyProps = expression(1);
      /**
       * @param {number} opts
       * @return {?}
       */
      var inspect = function(opts) {
        return "function" == typeof opts ? opts : void 0;
      };
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {Function} name
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, name) {
        return arguments.length < 2 ? inspect(obj[expectedNumberOfNonCommentArgs]) || inspect(modifyProps[expectedNumberOfNonCommentArgs]) : obj[expectedNumberOfNonCommentArgs] && obj[expectedNumberOfNonCommentArgs][name] || modifyProps[expectedNumberOfNonCommentArgs] && modifyProps[expectedNumberOfNonCommentArgs][name];
      };
    }, function(module, dataAndEvents, require) {
      var Block = require(1);
      var assert = require(6);
      var inspect = require(5);
      var next = require(27);
      var getActual = require(29);
      var nodes = require(16);
      var len = nodes.get;
      var register = nodes.enforce;
      /** @type {Array.<string>} */
      var s = String(String).split("String");
      (module.exports = function(expectedNumberOfNonCommentArgs, name, value, db) {
        /** @type {boolean} */
        var u = !!db && !!db.unsafe;
        /** @type {boolean} */
        var c = !!db && !!db.enumerable;
        /** @type {boolean} */
        var l = !!db && !!db.noTargetGet;
        if ("function" == typeof value) {
          if (!("string" != typeof name)) {
            if (!inspect(value, "name")) {
              assert(value, "name", name);
            }
          }
          /** @type {string} */
          register(value).source = s.join("string" == typeof name ? name : "");
        }
        if (expectedNumberOfNonCommentArgs !== Block) {
          if (u) {
            if (!l) {
              if (expectedNumberOfNonCommentArgs[name]) {
                /** @type {boolean} */
                c = true;
              }
            }
          } else {
            delete expectedNumberOfNonCommentArgs[name];
          }
          if (c) {
            /** @type {boolean} */
            expectedNumberOfNonCommentArgs[name] = value;
          } else {
            assert(expectedNumberOfNonCommentArgs, name, value);
          }
        } else {
          if (c) {
            /** @type {boolean} */
            expectedNumberOfNonCommentArgs[name] = value;
          } else {
            next(name, value);
          }
        }
      })(Function.prototype, "toString", function() {
        return "function" == typeof this && len(this).source || getActual(this);
      });
    }, function(module, dataAndEvents) {
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs) {
        if ("function" != typeof expectedNumberOfNonCommentArgs) {
          throw TypeError(String(expectedNumberOfNonCommentArgs) + " is not a function");
        }
        return expectedNumberOfNonCommentArgs;
      };
    }, function(module, dataAndEvents) {
      /** @type {boolean} */
      module.exports = false;
    }, function($, dataAndEvents, require) {
      var remove;
      var query;
      var walk;
      var helper = require(67);
      var Block = require(1);
      var getActual = require(11);
      var assert = require(6);
      var filter = require(5);
      var input = require(30);
      var nodes = require(31);
      var SimulatedScope = Block.WeakMap;
      if (helper) {
        var scope = new SimulatedScope;
        var getter = scope.get;
        var _$watch = scope.has;
        var fn = scope.set;
        /**
         * @param {?} first
         * @param {?} expectedHashCode
         * @return {?}
         */
        remove = function(first, expectedHashCode) {
          return fn.call(scope, first, expectedHashCode), expectedHashCode;
        };
        /**
         * @param {?} first
         * @return {?}
         */
        query = function(first) {
          return getter.call(scope, first) || {};
        };
        /**
         * @param {?} tree
         * @return {?}
         */
        walk = function(tree) {
          return _$watch.call(scope, tree);
        };
      } else {
        var i = input("state");
        /** @type {boolean} */
        nodes[i] = true;
        /**
         * @param {?} first
         * @param {?} expectedHashCode
         * @return {?}
         */
        remove = function(first, expectedHashCode) {
          return assert(first, i, expectedHashCode), expectedHashCode;
        };
        /**
         * @param {?} first
         * @return {?}
         */
        query = function(first) {
          return filter(first, i) ? first[i] : {};
        };
        /**
         * @param {?} tree
         * @return {?}
         */
        walk = function(tree) {
          return filter(tree, i);
        };
      }
      $.exports = {
        /** @type {function (?, ?): ?} */
        set : remove,
        /** @type {function (?): ?} */
        get : query,
        /** @type {function (?): ?} */
        has : walk,
        /**
         * @param {?} nodes
         * @return {?}
         */
        enforce : function(nodes) {
          return walk(nodes) ? query(nodes) : remove(nodes, {});
        },
        /**
         * @param {string} method
         * @return {?}
         */
        getterFor : function(method) {
          return function(nodes) {
            var n;
            if (!getActual(nodes) || (n = query(nodes)).type !== method) {
              throw TypeError("Incompatible receiver, " + method + " required");
            }
            return n;
          };
        }
      };
    }, function(module, dataAndEvents) {
      /** @type {function (this:*): string} */
      var ostring = {}.toString;
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs) {
        return ostring.call(expectedNumberOfNonCommentArgs).slice(8, -1);
      };
    }, function(module, dataAndEvents) {
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs) {
        if (null == expectedNumberOfNonCommentArgs) {
          throw TypeError("Can't call method on " + expectedNumberOfNonCommentArgs);
        }
        return expectedNumberOfNonCommentArgs;
      };
    }, function(module, dataAndEvents) {
      module.exports = {};
    }, function(mod, dataAndEvents, $sanitize) {
      var __bind = $sanitize(14);
      /**
       * @param {?} promise
       * @return {undefined}
       */
      var Deferred = function(promise) {
        var text;
        var doneResults;
        this.promise = new promise(function(textAlt, data) {
          if (void 0 !== text || void 0 !== doneResults) {
            throw TypeError("Bad Promise constructor");
          }
          text = textAlt;
          doneResults = data;
        });
        this.resolve = __bind(text);
        this.reject = __bind(doneResults);
      };
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {?}
       */
      mod.exports.f = function(expectedNumberOfNonCommentArgs) {
        return new Deferred(expectedNumberOfNonCommentArgs);
      };
    }, function(module, dataAndEvents) {
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {Object} object
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, object) {
        return{
          enumerable : !(1 & expectedNumberOfNonCommentArgs),
          configurable : !(2 & expectedNumberOfNonCommentArgs),
          writable : !(4 & expectedNumberOfNonCommentArgs),
          value : object
        };
      };
    }, function(module, dataAndEvents) {
      /** @type {function (*): number} */
      var ceil = Math.ceil;
      /** @type {function (*): number} */
      var floor = Math.floor;
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs) {
        return isNaN(expectedNumberOfNonCommentArgs = +expectedNumberOfNonCommentArgs) ? 0 : (expectedNumberOfNonCommentArgs > 0 ? floor : ceil)(expectedNumberOfNonCommentArgs);
      };
    }, function(module, dataAndEvents, require) {
      var format = require(72);
      var getActual = require(18);
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs) {
        return format(getActual(expectedNumberOfNonCommentArgs));
      };
    }, function(module, dataAndEvents, require) {
      var getActual = require(4);
      var objDisplay = require(93);
      var assert = require(33);
      var test = require(38);
      var inspect = require(94);
      var log = require(95);
      /**
       * @param {?} descriptor
       * @param {Object} e
       * @return {undefined}
       */
      var Promise = function(descriptor, e) {
        this.stopped = descriptor;
        /** @type {Object} */
        this.result = e;
      };
      /**
       * @param {string} gotoEnd
       * @return {?}
       */
      (module.exports = function(expectedNumberOfNonCommentArgs, proto, data, err, chai) {
        var it;
        var expected;
        var i;
        var l;
        var obj;
        var ostring;
        var current;
        var callback = test(proto, data, err ? 2 : 1);
        if (chai) {
          /** @type {number} */
          it = expectedNumberOfNonCommentArgs;
        } else {
          if ("function" != typeof(expected = inspect(expectedNumberOfNonCommentArgs))) {
            throw TypeError("Target is not iterable");
          }
          if (objDisplay(expected)) {
            /** @type {number} */
            i = 0;
            l = assert(expectedNumberOfNonCommentArgs.length);
            for (;l > i;i++) {
              if ((obj = err ? callback(getActual(current = expectedNumberOfNonCommentArgs[i])[0], current[1]) : callback(expectedNumberOfNonCommentArgs[i])) && obj instanceof Promise) {
                return obj;
              }
            }
            return new Promise(false);
          }
          it = expected.call(expectedNumberOfNonCommentArgs);
        }
        ostring = it.next;
        for (;!(current = ostring.call(it)).done;) {
          if ("object" == typeof(obj = log(it, callback, current.value, err)) && (obj && obj instanceof Promise)) {
            return obj;
          }
        }
        return new Promise(false);
      }).stop = function(gotoEnd) {
        return new Promise(true, gotoEnd);
      };
    }, function(module, dataAndEvents) {
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs) {
        try {
          return{
            error : false,
            value : expectedNumberOfNonCommentArgs()
          };
        } catch (origValue) {
          return{
            error : true,
            value : origValue
          };
        }
      };
    }, function(module, dataAndEvents, $sanitize) {
      var result = {};
      /** @type {string} */
      result[$sanitize(3)("toStringTag")] = "z";
      /** @type {boolean} */
      module.exports = "[object z]" === String(result);
    }, function(module, dataAndEvents, require) {
      var ctor = require(1);
      var getActual = require(6);
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {Function} proto
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, proto) {
        try {
          getActual(ctor, expectedNumberOfNonCommentArgs, proto);
        } catch (r) {
          /** @type {Function} */
          ctor[expectedNumberOfNonCommentArgs] = proto;
        }
        return proto;
      };
    }, function(module, dataAndEvents, require) {
      var collection = require(1);
      var assert = require(11);
      var e = collection.document;
      var s = assert(e) && assert(e.createElement);
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs) {
        return s ? e.createElement(expectedNumberOfNonCommentArgs) : {};
      };
    }, function(module, dataAndEvents, require) {
      var mod = require(41);
      var ostring = Function.toString;
      if ("function" != typeof mod.inspectSource) {
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        mod.inspectSource = function(expectedNumberOfNonCommentArgs) {
          return ostring.call(expectedNumberOfNonCommentArgs);
        };
      }
      /** @type {function (number): ?} */
      module.exports = mod.inspectSource;
    }, function(module, dataAndEvents, require) {
      var sorter = require(40);
      var getActual = require(44);
      var key = sorter("keys");
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs) {
        return key[expectedNumberOfNonCommentArgs] || (key[expectedNumberOfNonCommentArgs] = getActual(expectedNumberOfNonCommentArgs));
      };
    }, function(module, dataAndEvents) {
      module.exports = {};
    }, function(dataAndEvents, entry, format) {
      var f = format(9);
      var query = format(71);
      var buildParams = format(21);
      var msg = format(23);
      var get_mangled = format(43);
      var dataAttr = format(5);
      var cohortString = format(42);
      /** @type {function (Object, string): (ObjectPropertyDescriptor|undefined)} */
      var getOwnPropertyDescriptor = Object.getOwnPropertyDescriptor;
      /** @type {Function} */
      entry.f = f ? getOwnPropertyDescriptor : function(expectedNumberOfNonCommentArgs, name) {
        if (expectedNumberOfNonCommentArgs = msg(expectedNumberOfNonCommentArgs), name = get_mangled(name, true), cohortString) {
          try {
            return getOwnPropertyDescriptor(expectedNumberOfNonCommentArgs, name);
          } catch (t) {
          }
        }
        if (dataAttr(expectedNumberOfNonCommentArgs, name)) {
          return buildParams(!query.f.call(expectedNumberOfNonCommentArgs, name), expectedNumberOfNonCommentArgs[name]);
        }
      };
    }, function(module, dataAndEvents, require) {
      var getActual = require(22);
      /** @type {function (...[*]): number} */
      var nativeMin = Math.min;
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs) {
        return expectedNumberOfNonCommentArgs > 0 ? nativeMin(getActual(expectedNumberOfNonCommentArgs), 9007199254740991) : 0;
      };
    }, function(module, dataAndEvents) {
      /** @type {Array} */
      module.exports = ["constructor", "hasOwnProperty", "isPrototypeOf", "propertyIsEnumerable", "toLocaleString", "toString", "valueOf"];
    }, function(module, dataAndEvents, require) {
      var hasKey = require(5);
      var getActual = require(80);
      var typeOf = require(30);
      var Block = require(81);
      var type = typeOf("IE_PROTO");
      var objectProto = Object.prototype;
      /** @type {Function} */
      module.exports = Block ? Object.getPrototypeOf : function(expectedNumberOfNonCommentArgs) {
        return expectedNumberOfNonCommentArgs = getActual(expectedNumberOfNonCommentArgs), hasKey(expectedNumberOfNonCommentArgs, type) ? expectedNumberOfNonCommentArgs[type] : "function" == typeof expectedNumberOfNonCommentArgs.constructor && expectedNumberOfNonCommentArgs instanceof expectedNumberOfNonCommentArgs.constructor ? expectedNumberOfNonCommentArgs.constructor.prototype : expectedNumberOfNonCommentArgs instanceof Object ? objectProto : null;
      };
    }, function(module, dataAndEvents, require) {
      var next;
      var createObject = require(4);
      var equal = require(82);
      var tokenized = require(34);
      var json = require(31);
      var nodes = require(52);
      var query = require(28);
      var helper = require(30);
      var response = helper("IE_PROTO");
      /**
       * @return {undefined}
       */
      var ctor = function() {
      };
      /**
       * @param {string} range
       * @return {?}
       */
      var fn = function(range) {
        return "<script>" + range + "\x3c/script>";
      };
      /**
       * @return {?}
       */
      var init = function() {
        try {
          next = document.domain && new ActiveXObject("htmlfile");
        } catch (t) {
        }
        var out;
        var el;
        init = next ? function(out) {
          out.write(fn(""));
          out.close();
          var YObject = out.parentWindow.Object;
          return out = null, YObject;
        }(next) : ((el = query("iframe")).style.display = "none", nodes.appendChild(el), el.src = String("javascript:"), (out = el.contentWindow.document).open(), out.write(fn("document.F=Object")), out.close(), out.F);
        var index = tokenized.length;
        for (;index--;) {
          delete init.prototype[tokenized[index]];
        }
        return init();
      };
      /** @type {boolean} */
      json[response] = true;
      /** @type {function ((Object|null), (Object|null)=): Object} */
      module.exports = Object.create || function(expectedNumberOfNonCommentArgs, object) {
        var result;
        return null !== expectedNumberOfNonCommentArgs ? (ctor.prototype = createObject(expectedNumberOfNonCommentArgs), result = new ctor, ctor.prototype = null, result[response] = expectedNumberOfNonCommentArgs) : result = init(), void 0 === object ? result : equal(result, object);
      };
    }, function(module, dataAndEvents, require) {
      var setDescriptor = require(10).f;
      var getActual = require(5);
      var rvar = require(3)("toStringTag");
      /**
       * @param {Function} expectedNumberOfNonCommentArgs
       * @param {Object} object
       * @param {boolean} value
       * @return {undefined}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, object, value) {
        if (expectedNumberOfNonCommentArgs) {
          if (!getActual(expectedNumberOfNonCommentArgs = value ? expectedNumberOfNonCommentArgs : expectedNumberOfNonCommentArgs.prototype, rvar)) {
            setDescriptor(expectedNumberOfNonCommentArgs, rvar, {
              configurable : true,
              value : object
            });
          }
        }
      };
    }, function(module, dataAndEvents, require) {
      var getActual = require(14);
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {Function} object
       * @param {boolean} value
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, object, value) {
        if (getActual(expectedNumberOfNonCommentArgs), void 0 === object) {
          return expectedNumberOfNonCommentArgs;
        }
        switch(value) {
          case 0:
            return function() {
              return expectedNumberOfNonCommentArgs.call(object);
            };
          case 1:
            return function(mapper) {
              return expectedNumberOfNonCommentArgs.call(object, mapper);
            };
          case 2:
            return function(mapper, graphics) {
              return expectedNumberOfNonCommentArgs.call(object, mapper, graphics);
            };
          case 3:
            return function(mapper, graphics, capture) {
              return expectedNumberOfNonCommentArgs.call(object, mapper, graphics, capture);
            };
        }
        return function() {
          return expectedNumberOfNonCommentArgs.apply(object, arguments);
        };
      };
    }, function(f, dataAndEvents, require) {
      var getName = require(12);
      f.exports = getName("navigator", "userAgent") || "";
    }, function(module, dataAndEvents, require) {
      var binary = require(15);
      var Block = require(41);
      (module.exports = function(expectedNumberOfNonCommentArgs, actual) {
        return Block[expectedNumberOfNonCommentArgs] || (Block[expectedNumberOfNonCommentArgs] = void 0 !== actual ? actual : {});
      })("versions", []).push({
        version : "3.6.4",
        mode : binary ? "pure" : "global",
        copyright : "\u00c2\u00a9 2020 Denis Pushkarev (zloirock.ru)"
      });
    }, function(module, dataAndEvents, require) {
      var Block = require(1);
      var getActual = require(27);
      var JsDiff = Block["__core-js_shared__"] || getActual("__core-js_shared__", {});
      module.exports = JsDiff;
    }, function(module, dataAndEvents, func) {
      var actual = func(9);
      var lambda = func(7);
      var unwrap = func(28);
      /** @type {boolean} */
      module.exports = !actual && !lambda(function() {
        return 7 != Object.defineProperty(unwrap("div"), "a", {
          /**
           * @return {?}
           */
          get : function() {
            return 7;
          }
        }).a;
      });
    }, function(module, dataAndEvents, require) {
      var getActual = require(11);
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {boolean} proto
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, proto) {
        if (!getActual(expectedNumberOfNonCommentArgs)) {
          return expectedNumberOfNonCommentArgs;
        }
        var valueOf;
        var str;
        if (proto && ("function" == typeof(valueOf = expectedNumberOfNonCommentArgs.toString) && !getActual(str = valueOf.call(expectedNumberOfNonCommentArgs)))) {
          return str;
        }
        if ("function" == typeof(valueOf = expectedNumberOfNonCommentArgs.valueOf) && !getActual(str = valueOf.call(expectedNumberOfNonCommentArgs))) {
          return str;
        }
        if (!proto && ("function" == typeof(valueOf = expectedNumberOfNonCommentArgs.toString) && !getActual(str = valueOf.call(expectedNumberOfNonCommentArgs)))) {
          return str;
        }
        throw TypeError("Can't convert object to primitive value");
      };
    }, function(module, dataAndEvents) {
      /** @type {number} */
      var count = 0;
      /** @type {number} */
      var id = Math.random();
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs) {
        return "Symbol(" + String(void 0 === expectedNumberOfNonCommentArgs ? "" : expectedNumberOfNonCommentArgs) + ")_" + (++count + id).toString(36);
      };
    }, function(module, dataAndEvents, require) {
      var getActual = require(7);
      /** @type {boolean} */
      module.exports = !!Object.getOwnPropertySymbols && !getActual(function() {
        return!String(Symbol());
      });
    }, function(module, dataAndEvents, require) {
      var Block = require(26);
      var forOwn = require(17);
      var pdataOld = require(3)("toStringTag");
      /** @type {boolean} */
      var content = "Arguments" == forOwn(function() {
        return arguments;
      }());
      module.exports = Block ? forOwn : function(expectedNumberOfNonCommentArgs) {
        var object;
        var pdataCur;
        var idx;
        return void 0 === expectedNumberOfNonCommentArgs ? "Undefined" : null === expectedNumberOfNonCommentArgs ? "Null" : "string" == typeof(pdataCur = function(tagMap, value) {
          try {
            return tagMap[value];
          } catch (t) {
          }
        }(object = Object(expectedNumberOfNonCommentArgs), pdataOld)) ? pdataCur : content ? forOwn(object) : "Object" == (idx = forOwn(object)) && "function" == typeof object.callee ? "Arguments" : idx;
      };
    }, function(module, dataAndEvents, require) {
      var inspect = require(8);
      var flag = require(79);
      var toString = require(35);
      var parse = require(53);
      var compile = require(37);
      var getActual = require(6);
      var debug = require(13);
      var sorter = require(3);
      var Block = require(15);
      var target = require(19);
      var range = require(51);
      var ar = range.IteratorPrototype;
      var a = range.BUGGY_SAFARI_ITERATORS;
      var key = sorter("iterator");
      /**
       * @return {?}
       */
      var copy = function() {
        return this;
      };
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {Object} object
       * @param {Function} data
       * @param {?} msg
       * @param {string} method
       * @param {(Function|string)} chai
       * @param {?} includeAll
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, object, data, msg, method, chai, includeAll) {
        flag(data, object, msg);
        var str;
        var types;
        var type;
        /**
         * @param {Function} key
         * @return {?}
         */
        var callback = function(key) {
          if (key === method && fn) {
            return fn;
          }
          if (!a && key in obj) {
            return obj[key];
          }
          switch(key) {
            case "keys":
            ;
            case "values":
            ;
            case "entries":
              return function() {
                return new data(this, key);
              };
          }
          return function() {
            return new data(this);
          };
        };
        /** @type {string} */
        var name = object + " Iterator";
        /** @type {boolean} */
        var b = false;
        var obj = expectedNumberOfNonCommentArgs.prototype;
        var values = obj[key] || (obj["@@iterator"] || method && obj[method]);
        var fn = !a && values || callback(method);
        var conditional = "Array" == object && obj.entries || values;
        if (conditional && (str = toString(conditional.call(new expectedNumberOfNonCommentArgs)), ar !== Object.prototype && (str.next && (Block || (toString(str) === ar || (parse ? parse(str, ar) : "function" != typeof str[key] && getActual(str, key, copy))), compile(str, name, true, true), Block && (target[name] = copy)))), "values" == method && (values && ("values" !== values.name && (b = true, fn = function() {
          return values.call(this);
        }))), Block && !includeAll || (obj[key] === fn || getActual(obj, key, fn)), target[object] = fn, method) {
          if (types = {
            values : callback("values"),
            keys : chai ? fn : callback("keys"),
            entries : callback("entries")
          }, includeAll) {
            for (type in types) {
              if (a || (b || !(type in obj))) {
                debug(obj, type, types[type]);
              }
            }
          } else {
            inspect({
              target : object,
              proto : true,
              forced : a || b
            }, types);
          }
        }
        return types;
      };
    }, function(module, dataAndEvents, topic) {
      var out = topic(1);
      module.exports = out;
    }, function(module, dataAndEvents, require) {
      var inspect = require(5);
      var getActual = require(23);
      var callback = require(76).indexOf;
      var j = require(31);
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {Object} proto
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, proto) {
        var key;
        var actual = getActual(expectedNumberOfNonCommentArgs);
        /** @type {number} */
        var maxScanLen = 0;
        /** @type {Array} */
        var ret = [];
        for (key in actual) {
          if (!inspect(j, key)) {
            if (inspect(actual, key)) {
              ret.push(key);
            }
          }
        }
        for (;proto.length > maxScanLen;) {
          if (inspect(actual, key = proto[maxScanLen++])) {
            if (!~callback(ret, key)) {
              ret.push(key);
            }
          }
        }
        return ret;
      };
    }, function(module, dataAndEvents, require) {
      var getName = require(7);
      /** @type {RegExp} */
      var r20 = /#|\.prototype\./;
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {Function} value
       * @return {?}
       */
      var parse = function(expectedNumberOfNonCommentArgs, value) {
        var next = args[promote(expectedNumberOfNonCommentArgs)];
        return next == end || next != current && ("function" == typeof value ? getName(value) : !!value);
      };
      /** @type {function (number): ?} */
      var promote = parse.normalize = function(style) {
        return String(style).replace(r20, ".").toLowerCase();
      };
      var args = parse.data = {};
      /** @type {string} */
      var current = parse.NATIVE = "N";
      /** @type {string} */
      var end = parse.POLYFILL = "P";
      /** @type {function (number, Function): ?} */
      module.exports = parse;
    }, function(module, dataAndEvents, require) {
      var obj;
      var formatter;
      var body;
      var tighten = require(35);
      var hasOwn = require(6);
      var helper = require(5);
      var getName = require(3);
      var Block = require(15);
      var name = getName("iterator");
      /** @type {boolean} */
      var BUGGY_SAFARI_ITERATORS = false;
      if ([].keys) {
        if ("next" in (body = [].keys())) {
          if ((formatter = tighten(tighten(body))) !== Object.prototype) {
            obj = formatter;
          }
        } else {
          /** @type {boolean} */
          BUGGY_SAFARI_ITERATORS = true;
        }
      }
      if (null == obj) {
        obj = {};
      }
      if (!Block) {
        if (!helper(obj, name)) {
          hasOwn(obj, name, function() {
            return this;
          });
        }
      }
      module.exports = {
        IteratorPrototype : obj,
        BUGGY_SAFARI_ITERATORS : BUGGY_SAFARI_ITERATORS
      };
    }, function(module, dataAndEvents, require) {
      var factory = require(12);
      module.exports = factory("document", "documentElement");
    }, function(module, dataAndEvents, require) {
      var getActual = require(4);
      var getName = require(84);
      module.exports = Object.setPrototypeOf || ("__proto__" in {} ? function() {
        var set;
        /** @type {boolean} */
        var op = false;
        var xs = {};
        try {
          (set = Object.getOwnPropertyDescriptor(Object.prototype, "__proto__").set).call(xs, []);
          /** @type {boolean} */
          op = xs instanceof Array;
        } catch (t) {
        }
        return function(obj, value) {
          return getActual(obj), getName(value), op ? set.call(obj, value) : obj.__proto__ = value, obj;
        };
      }() : void 0);
    }, function(module, dataAndEvents, require) {
      var global = require(1);
      module.exports = global.Promise;
    }, function(module, dataAndEvents, require) {
      var getActual = require(4);
      var inspect = require(14);
      var prop = require(3)("species");
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {Object} proto
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, proto) {
        var a;
        var obj = getActual(expectedNumberOfNonCommentArgs).constructor;
        return void 0 === obj || null == (a = getActual(obj)[prop]) ? proto : inspect(a);
      };
    }, function(c, dataAndEvents, require) {
      var callback;
      var channel;
      var thisObj;
      var global = require(1);
      var core = require(7);
      var getActual = require(17);
      var makeIterator = require(38);
      var xml = require(52);
      var inspect = require(28);
      var Block = require(57);
      var l = global.location;
      var pass = global.setImmediate;
      var clear = global.clearImmediate;
      var process = global.process;
      var MessageChannel = global.MessageChannel;
      var _ = global.Dispatch;
      /** @type {number} */
      var a = 0;
      var data = {};
      /**
       * @param {number} property
       * @return {undefined}
       */
      var resolve = function(property) {
        if (data.hasOwnProperty(property)) {
          var fn = data[property];
          delete data[property];
          fn();
        }
      };
      /**
       * @param {number} value
       * @return {?}
       */
      var expect = function(value) {
        return function() {
          resolve(value);
        };
      };
      /**
       * @param {MessageEvent} e
       * @return {undefined}
       */
      var completed = function(e) {
        resolve(e.data);
      };
      /**
       * @param {string} o
       * @return {undefined}
       */
      var request = function(o) {
        global.postMessage(o + "", l.protocol + "//" + l.host);
      };
      if (!(pass && clear)) {
        /**
         * @param {?} fn
         * @return {?}
         */
        pass = function(fn) {
          /** @type {Array} */
          var ta = [];
          /** @type {number} */
          var i = 1;
          for (;arguments.length > i;) {
            ta.push(arguments[i++]);
          }
          return data[++a] = function() {
            ("function" == typeof fn ? fn : Function(fn)).apply(void 0, ta);
          }, callback(a), a;
        };
        /**
         * @param {?} first
         * @return {undefined}
         */
        clear = function(first) {
          delete data[first];
        };
        if ("process" == getActual(process)) {
          /**
           * @param {number} stream
           * @return {undefined}
           */
          callback = function(stream) {
            process.nextTick(expect(stream));
          };
        } else {
          if (_ && _.now) {
            /**
             * @param {number} result
             * @return {undefined}
             */
            callback = function(result) {
              _.now(expect(result));
            };
          } else {
            if (MessageChannel && !Block) {
              thisObj = (channel = new MessageChannel).port2;
              /** @type {function (MessageEvent): undefined} */
              channel.port1.onmessage = completed;
              callback = makeIterator(thisObj.postMessage, thisObj, 1);
            } else {
              if (!global.addEventListener || ("function" != typeof postMessage || (global.importScripts || core(request)))) {
                /** @type {function (number): undefined} */
                callback = "onreadystatechange" in inspect("script") ? function(result) {
                  /**
                   * @return {undefined}
                   */
                  xml.appendChild(inspect("script")).onreadystatechange = function() {
                    xml.removeChild(this);
                    resolve(result);
                  };
                } : function(func) {
                  setTimeout(expect(func), 0);
                };
              } else {
                /** @type {function (string): undefined} */
                callback = request;
                global.addEventListener("message", completed, false);
              }
            }
          }
        }
      }
      c.exports = {
        set : pass,
        clear : clear
      };
    }, function(module, dataAndEvents, getName) {
      var name = getName(39);
      /** @type {boolean} */
      module.exports = /(iphone|ipod|ipad).*applewebkit/i.test(name);
    }, function(module, dataAndEvents, require) {
      var inspect = require(4);
      var getActual = require(11);
      var argv = require(20);
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {Object} proto
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, proto) {
        if (inspect(expectedNumberOfNonCommentArgs), getActual(proto) && proto.constructor === expectedNumberOfNonCommentArgs) {
          return proto;
        }
        var invokeDfd = argv.f(expectedNumberOfNonCommentArgs);
        return(0, invokeDfd.resolve)(proto), invokeDfd.promise;
      };
    }, function(dataAndEvents, deepDataAndEvents, require) {
      var getActual = require(8);
      var fn = require(14);
      var v = require(20);
      var selector = require(25);
      var isArray = require(24);
      getActual({
        target : "Promise",
        stat : true
      }, {
        /**
         * @param {?} promises
         * @return {?}
         */
        allSettled : function(promises) {
          var self = this;
          var result = v.f(self);
          var resolve = result.resolve;
          var iterator = result.reject;
          var node = selector(function() {
            var callback = fn(self.resolve);
            /** @type {Array} */
            var expectedNumberOfNonCommentArgs = [];
            /** @type {number} */
            var rightId = 0;
            /** @type {number} */
            var u = 1;
            isArray(promises, function(node) {
              /** @type {number} */
              var id = rightId++;
              /** @type {boolean} */
              var a = false;
              expectedNumberOfNonCommentArgs.push(void 0);
              u++;
              callback.call(self, node).then(function(x) {
                if (!a) {
                  /** @type {boolean} */
                  a = true;
                  expectedNumberOfNonCommentArgs[id] = {
                    status : "fulfilled",
                    value : x
                  };
                  if (!--u) {
                    resolve(expectedNumberOfNonCommentArgs);
                  }
                }
              }, function(err) {
                if (!a) {
                  /** @type {boolean} */
                  a = true;
                  expectedNumberOfNonCommentArgs[id] = {
                    status : "rejected",
                    reason : err
                  };
                  if (!--u) {
                    resolve(expectedNumberOfNonCommentArgs);
                  }
                }
              });
            });
            if (!--u) {
              resolve(expectedNumberOfNonCommentArgs);
            }
          });
          return node.error && iterator(node.value), result.promise;
        }
      });
    }, function(dataAndEvents, expectedNumberOfNonCommentArgs, f) {
      (function(vfs) {
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @param {(Array|NodeList)} obj
         * @return {undefined}
         */
        function defineProperty(expectedNumberOfNonCommentArgs, obj) {
          /** @type {number} */
          var i = 0;
          for (;i < obj.length;i++) {
            var desc = obj[i];
            desc.enumerable = desc.enumerable || false;
            /** @type {boolean} */
            desc.configurable = true;
            if ("value" in desc) {
              /** @type {boolean} */
              desc.writable = true;
            }
            Object.defineProperty(expectedNumberOfNonCommentArgs, desc.key, desc);
          }
        }
        f.d(expectedNumberOfNonCommentArgs, "a", function() {
          return s;
        });
        var $ = f(2);
        var tree = f(0);
        var s = function() {
          /**
           * @param {string} key
           * @param {string} match
           * @param {string} p1
           * @param {Function} v
           * @param {?} iterations
           * @param {?} ampm
           * @param {Worker} inCode
           * @return {undefined}
           */
          function replacer(key, match, p1, v, iterations, ampm, inCode) {
            !function(dataAndEvents, replacer) {
              if (!(dataAndEvents instanceof replacer)) {
                throw new TypeError("Cannot call a class as a function");
              }
            }(this, replacer);
            this.N = new tree.BigInteger(key, 16);
            this.g = new tree.BigInteger(match, 16);
            this.hashFunction = p1.toUpperCase();
            /** @type {Function} */
            this.version = v;
            this.iterations = iterations;
            /** @type {null} */
            this.S = null;
            /** @type {null} */
            this.M1 = null;
            this.computeVerifier = ampm;
            /** @type {Worker} */
            this.workerUrlOverride = inCode;
          }
          var ctor;
          var suiteView;
          var current;
          return ctor = replacer, (suiteView = [{
            key : "step1",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @param {boolean} value
             * @param {string} v
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object, value, v) {
              var e;
              var millis;
              var fsElement;
              var length;
              var h;
              var ret;
              var self = this;
              return regeneratorRuntime.async(function(self) {
                for (;;) {
                  switch(self.prev = self.next) {
                    case 0:
                      return this.workerUrlOverride ? this.backgroundWorker = new Worker(this.workerUrlOverride) : this.backgroundWorker = new Worker(vfs, {}), this.check(expectedNumberOfNonCommentArgs, "username"), this.check(object, "password"), this.check(value, "salt"), e = new tree.BigInteger(v, 16), self.next = 7, regeneratorRuntime.awrap(this.checkValidPublicValue(e));
                    case 7:
                      return millis = this.randomPrivateA(), self.next = 10, regeneratorRuntime.awrap(this.computePublicA(millis));
                    case 10:
                      return fsElement = self.sent, self.next = 13, regeneratorRuntime.awrap(this.computeU(fsElement, e));
                    case 13:
                      return length = self.sent, self.next = 16, regeneratorRuntime.awrap(this.checkValidPublicValue(length));
                    case 16:
                      return self.next = 18, regeneratorRuntime.awrap(this.computeK());
                    case 18:
                      return h = self.sent, ret = new Promise(function(callback, send) {
                        /**
                         * @param {MessageEvent} messageEvent
                         * @return {?}
                         */
                        self.backgroundWorker.onmessage = function(messageEvent) {
                          var onComplete;
                          return regeneratorRuntime.async(function(stream) {
                            for (;;) {
                              switch(stream.prev = stream.next) {
                                case 0:
                                  if ("error" !== messageEvent.data[0]) {
                                    /** @type {number} */
                                    stream.next = 4;
                                    break;
                                  }
                                  return self.backgroundWorker.terminate(), send(messageEvent.data[1]), stream.abrupt("return");
                                case 4:
                                  return onComplete = new tree.BigInteger(messageEvent.data[1]), true === self.computeVerifier && (self.verifier = new tree.BigInteger(messageEvent.data[2])), stream.next = 8, regeneratorRuntime.awrap(self.computeS(onComplete, length, millis, h, e));
                                case 8:
                                  return self.S = stream.sent, stream.next = 11, regeneratorRuntime.awrap(self.computeM1(fsElement, e));
                                case 11:
                                  self.M1 = stream.sent;
                                  callback({
                                    publicA : fsElement,
                                    clientEvidenceM1 : self.M1
                                  });
                                case 13:
                                ;
                                case "end":
                                  return stream.stop();
                              }
                            }
                          });
                        };
                      }), this.backgroundWorker.postMessage(["computeX", expectedNumberOfNonCommentArgs, object, value, this.version, this.iterations, this.computeVerifier]), self.abrupt("return", ret);
                    case 22:
                    ;
                    case "end":
                      return self.stop();
                  }
                }
              }, null, this);
            }
          }, {
            key : "step2",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @param {string} value
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object, value) {
              var array;
              var children;
              var rest;
              var data;
              var width;
              var recurring;
              var node;
              return regeneratorRuntime.async(function(self) {
                for (;;) {
                  switch(self.prev = self.next) {
                    case 0:
                      return this.check(object, "M1"), this.check(value, "M2"), array = expectedNumberOfNonCommentArgs.toByteArray(), children = object.toByteArray(), rest = this.S.toByteArray(), data = new Uint8Array(array.concat(children).concat(rest)), self.next = 8, regeneratorRuntime.awrap(this.hash(this.arrayBufferToHexString(data), true));
                    case 8:
                      return width = self.sent, recurring = new tree.BigInteger(new Uint8Array(width)), node = new tree.BigInteger(value, 16), self.abrupt("return", node.equals(recurring));
                    case 12:
                    ;
                    case "end":
                      return self.stop();
                  }
                }
              }, null, this);
            }
          }, {
            key : "createRandomBigIntegerInRange",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} proto
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, proto) {
              var wrapper = expectedNumberOfNonCommentArgs.compareTo(proto);
              if (wrapper >= 0) {
                if (wrapper > 0) {
                  throw new Error('"min" may not be greater than "max"');
                }
                return expectedNumberOfNonCommentArgs;
              }
              if (expectedNumberOfNonCommentArgs.bitLength() > proto.bitLength() / 2) {
                return this.createRandomBigIntegerInRange(tree.BigInteger.ZERO, proto.subtract(expectedNumberOfNonCommentArgs)).add(expectedNumberOfNonCommentArgs);
              }
              /** @type {number} */
              var o = 0;
              for (;o < 1E3;++o) {
                var key = $.sjcl.codec.hex.fromBits($.sjcl.random.randomWords(32, 0));
                var node = new tree.BigInteger(key, 16);
                if (node.compareTo(expectedNumberOfNonCommentArgs) >= 0 && node.compareTo(proto) <= 0) {
                  return node;
                }
              }
              var rgb = $.sjcl.codec.hex.fromBits($.sjcl.random.randomWords(32, 0));
              return new tree.BigInteger(rgb, 16);
            }
          }, {
            key : "generateRandomValue",
            /**
             * @return {?}
             */
            value : function() {
              /** @type {number} */
              var shiftLeft = Math.min(256, this.N.bitLength() / 2);
              var r20 = tree.BigInteger.ONE.shiftLeft(shiftLeft - 1);
              var restoreScript = this.N.subtract(tree.BigInteger.ONE);
              return this.createRandomBigIntegerInRange(r20, restoreScript);
            }
          }, {
            key : "randomPrivateA",
            /**
             * @return {?}
             */
            value : function() {
              return this.generateRandomValue();
            }
          }, {
            key : "computePublicA",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs) {
              var e;
              return regeneratorRuntime.async(function(self) {
                for (;;) {
                  switch(self.prev = self.next) {
                    case 0:
                      return self.next = 2, regeneratorRuntime.awrap(this.unsignedModPow(this.g, expectedNumberOfNonCommentArgs, this.N));
                    case 2:
                      if (!((e = self.sent).equals(tree.BigInteger.ZERO) || (e.equals(tree.BigInteger.ONE) || e.equals(new tree.BigInteger("-1"))))) {
                        /** @type {number} */
                        self.next = 5;
                        break;
                      }
                      return self.abrupt("return", this.computePublicA(expectedNumberOfNonCommentArgs));
                    case 5:
                      return self.abrupt("return", e);
                    case 6:
                    ;
                    case "end":
                      return self.stop();
                  }
                }
              }, null, this);
            }
          }, {
            key : "computeU",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object) {
              var n;
              var v;
              return regeneratorRuntime.async(function(self) {
                for (;;) {
                  switch(self.prev = self.next) {
                    case 0:
                      return this.check(expectedNumberOfNonCommentArgs, "Public A"), self.next = 3, regeneratorRuntime.awrap(this.createPaddedPairHash(expectedNumberOfNonCommentArgs, object));
                    case 3:
                      return n = self.sent, v = new Uint8Array(n), self.abrupt("return", this.signedBigIntegerToUnsigned(new tree.BigInteger(v), 32));
                    case 6:
                    ;
                    case "end":
                      return self.stop();
                  }
                }
              }, null, this);
            }
          }, {
            key : "computeK",
            /**
             * @return {?}
             */
            value : function() {
              var width;
              return regeneratorRuntime.async(function(self) {
                for (;;) {
                  switch(self.prev = self.next) {
                    case 0:
                      return self.next = 2, regeneratorRuntime.awrap(this.createPaddedPairHash(this.N, this.g));
                    case 2:
                      return width = self.sent, self.abrupt("return", this.signedBigIntegerToUnsigned(new tree.BigInteger(new Uint8Array(width)), 32));
                    case 4:
                    ;
                    case "end":
                      return self.stop();
                  }
                }
              }, null, this);
            }
          }, {
            key : "computeS",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @param {number} dataAndEvents
             * @param {(Object|string)} event
             * @param {?} max
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object, dataAndEvents, event, max) {
              var columnWidth;
              var oldconfig;
              return regeneratorRuntime.async(function(self) {
                for (;;) {
                  switch(self.prev = self.next) {
                    case 0:
                      return columnWidth = object.multiply(expectedNumberOfNonCommentArgs).add(dataAndEvents), self.next = 3, regeneratorRuntime.awrap(this.unsignedModPow(this.g, expectedNumberOfNonCommentArgs, this.N));
                    case 3:
                      return self.t0 = event, oldconfig = self.sent.multiply(self.t0), self.abrupt("return", this.unsignedModPow(max.subtract(oldconfig), columnWidth, this.N));
                    case 6:
                    ;
                    case "end":
                      return self.stop();
                  }
                }
              }, null, this);
            }
          }, {
            key : "unsignedModPow",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @param {Object} data
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object, data) {
              var channel = this;
              return this.backgroundWorker.postMessage(["unsignedModPow", expectedNumberOfNonCommentArgs.toString(), object.toString(), data.toString()]), new Promise(function(done) {
                /**
                 * @param {MessageEvent} messageEvent
                 * @return {undefined}
                 */
                channel.backgroundWorker.onmessage = function(messageEvent) {
                  done(new tree.BigInteger(messageEvent.data[1]));
                };
              });
            }
          }, {
            key : "computeM1",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object) {
              var proto;
              return regeneratorRuntime.async(function(current) {
                for (;;) {
                  switch(current.prev = current.next) {
                    case 0:
                      return(proto = new $.sjcl.hash.sha256).update(this.bytesToBitArray(expectedNumberOfNonCommentArgs.toByteArray())), proto.update(this.bytesToBitArray(object.toByteArray())), proto.update(this.bytesToBitArray(this.S.toByteArray())), current.abrupt("return", new tree.BigInteger($.sjcl.codec.hex.fromBits(proto.finalize()), 16));
                    case 5:
                    ;
                    case "end":
                      return current.stop();
                  }
                }
              }, null, this);
            }
          }, {
            key : "checkValidPublicValue",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs) {
              var str;
              var approx;
              return regeneratorRuntime.async(function(self) {
                for (;;) {
                  switch(self.prev = self.next) {
                    case 0:
                      if (str = "Invalid public value", !expectedNumberOfNonCommentArgs) {
                        /** @type {number} */
                        self.next = 9;
                        break;
                      }
                      return self.next = 4, regeneratorRuntime.awrap(this.unsignedModPow(expectedNumberOfNonCommentArgs, tree.BigInteger.ONE, this.N));
                    case 4:
                      if (approx = self.sent, !(expectedNumberOfNonCommentArgs.equals(tree.BigInteger.ZERO) || (expectedNumberOfNonCommentArgs.equals(tree.BigInteger.ONE) || approx.equals(tree.BigInteger.ZERO)))) {
                        /** @type {number} */
                        self.next = 7;
                        break;
                      }
                      throw new Error(str);;
                    case 7:
                      /** @type {number} */
                      self.next = 10;
                      break;
                    case 9:
                      throw new Error(str);;
                    case 10:
                    ;
                    case "end":
                      return self.stop();
                  }
                }
              }, null, this);
            }
          }, {
            key : "bytesToBitArray",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs) {
              /** @type {number} */
              var conditionIndex = 0;
              /** @type {Array} */
              var buffer = [];
              /** @type {number} */
              var node = 0;
              /** @type {number} */
              conditionIndex = 0;
              for (;conditionIndex < expectedNumberOfNonCommentArgs.length;conditionIndex++) {
                /** @type {number} */
                node = node << 8 | 255 & expectedNumberOfNonCommentArgs[conditionIndex];
                if (3 == (3 & conditionIndex)) {
                  buffer.push(0 ^ node);
                  /** @type {number} */
                  node = 0;
                }
              }
              return 3 & conditionIndex && buffer.push($.sjcl.bitArray.partial(8 * (3 & conditionIndex), node)), buffer;
            }
          }, {
            key : "check",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Object} object
             * @return {undefined}
             */
            value : function(expectedNumberOfNonCommentArgs, object) {
              if (null == expectedNumberOfNonCommentArgs || ("" === expectedNumberOfNonCommentArgs || "0" === expectedNumberOfNonCommentArgs)) {
                throw new Error("".concat(object, " must not be null, empty or zero"));
              }
            }
          }, {
            key : "createPaddedPairHash",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object) {
              var curr;
              var result;
              var endtags;
              return regeneratorRuntime.async(function(current) {
                for (;;) {
                  switch(current.prev = current.next) {
                    case 0:
                      return curr = this.N.bitLength() / 4, result = expectedNumberOfNonCommentArgs.toString(16).padStart(curr, "0"), endtags = object.toString(16).padStart(curr, "0"), current.abrupt("return", this.hash(result + endtags, true));
                    case 4:
                    ;
                    case "end":
                      return current.stop();
                  }
                }
              }, null, this);
            }
          }, {
            key : "hash",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Object} object
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object) {
              var pdataCur;
              return regeneratorRuntime.async(function(current) {
                for (;;) {
                  switch(current.prev = current.next) {
                    case 0:
                      return pdataCur = $.sjcl.codec.bytes.toBits(expectedNumberOfNonCommentArgs), object && (pdataCur = $.sjcl.codec.hex.toBits(expectedNumberOfNonCommentArgs)), current.abrupt("return", new Promise(function($sanitize) {
                        $sanitize($.sjcl.codec.bytes.fromBits($.sjcl.hash.sha256.hash(pdataCur)));
                      }));
                    case 3:
                    ;
                    case "end":
                      return current.stop();
                  }
                }
              });
            }
          }, {
            key : "signedBigIntegerToUnsigned",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @param {Function} object
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs, object) {
              return expectedNumberOfNonCommentArgs.andNot((new tree.BigInteger("-1")).shiftLeft(8 * object));
            }
          }, {
            key : "arrayBufferToHexString",
            /**
             * @param {number} expectedNumberOfNonCommentArgs
             * @return {?}
             */
            value : function(expectedNumberOfNonCommentArgs) {
              var code;
              /** @type {number} */
              var resultItems = expectedNumberOfNonCommentArgs;
              /** @type {string} */
              var output = "";
              /** @type {number} */
              var i = 0;
              for (;i < resultItems.byteLength;i++) {
                if ((code = resultItems[i].toString(16)).length < 2) {
                  /** @type {string} */
                  code = "0" + code;
                }
                output += code;
              }
              return output;
            }
          }]) && defineProperty(ctor.prototype, suiteView), current && defineProperty(ctor, current), replacer;
        }();
      }).call(this, f(113));
    }, function(dataAndEvents, expectedNumberOfNonCommentArgs, d) {
      d.r(expectedNumberOfNonCommentArgs);
      d(62);
      d(105);
      d(112);
      var showChildren = d(2);
      d.d(expectedNumberOfNonCommentArgs, "PasswordManager", function() {
        return showChildren.PasswordManager;
      });
      d.d(expectedNumberOfNonCommentArgs, "PasswordVersion1", function() {
        return showChildren.PasswordVersion1;
      });
      d.d(expectedNumberOfNonCommentArgs, "PasswordVersion2", function() {
        return showChildren.PasswordVersion2;
      });
      var r = d(60);
      d.d(expectedNumberOfNonCommentArgs, "ClientSession", function() {
        return r.a;
      });
    }, function(module, dataAndEvents, topic) {
      var out = topic(63);
      topic(101);
      topic(102);
      topic(103);
      topic(104);
      module.exports = out;
    }, function(module, dataAndEvents, require) {
      require(64);
      require(69);
      require(85);
      require(89);
      require(59);
      require(100);
      var global = require(48);
      module.exports = global.Promise;
    }, function(dataAndEvents, deepDataAndEvents, func) {
      var actual = func(26);
      var makeInherit = func(13);
      var newResult = func(68);
      if (!actual) {
        makeInherit(Object.prototype, "toString", newResult, {
          unsafe : true
        });
      }
    }, function(module, dataAndEvents) {
      var dom;
      dom = function() {
        return this;
      }();
      try {
        dom = dom || (new Function("return this"))();
      } catch (t) {
        if ("object" == typeof window) {
          /** @type {Window} */
          dom = window;
        }
      }
      module.exports = dom;
    }, function(module, dataAndEvents, fun) {
      var exports = fun(45);
      module.exports = exports && (!Symbol.sham && "symbol" == typeof Symbol.iterator);
    }, function(module, dataAndEvents, require) {
      var expect = require(1);
      var next = require(29);
      var name = expect.WeakMap;
      /** @type {boolean} */
      module.exports = "function" == typeof name && /native code/.test(next(name));
    }, function(module, dataAndEvents, require) {
      var Block = require(26);
      var getActual = require(46);
      /** @type {Function} */
      module.exports = Block ? {}.toString : function() {
        return "[object " + getActual(this) + "]";
      };
    }, function(dataAndEvents, deepDataAndEvents, require) {
      var fn = require(70).charAt;
      var style = require(16);
      var is = require(47);
      var setStyle = style.set;
      var toObject = style.getterFor("String Iterator");
      is(String, "String", function(string) {
        setStyle(this, {
          type : "String Iterator",
          string : String(string),
          index : 0
        });
      }, function() {
        var result;
        var self = toObject(this);
        var list = self.string;
        var index = self.index;
        return index >= list.length ? {
          value : void 0,
          done : true
        } : (result = fn(list, index), self.index += result.length, {
          value : result,
          done : false
        });
      });
    }, function(module, dataAndEvents, getCallback) {
      var cb = getCallback(22);
      var callback = getCallback(18);
      /**
       * @param {boolean} i
       * @return {?}
       */
      var write = function(i) {
        return function(value, evt) {
          var el;
          var s;
          /** @type {string} */
          var source = String(callback(value));
          var index = cb(evt);
          /** @type {number} */
          var len = source.length;
          return index < 0 || index >= len ? i ? "" : void 0 : (el = source.charCodeAt(index)) < 55296 || (el > 56319 || (index + 1 === len || ((s = source.charCodeAt(index + 1)) < 56320 || s > 57343))) ? i ? source.charAt(index) : el : i ? source.slice(index, index + 2) : s - 56320 + (el - 55296 << 10) + 65536;
        };
      };
      module.exports = {
        codeAt : write(false),
        charAt : write(true)
      };
    }, function(dataAndEvents, entry, deepDataAndEvents) {
      /** @type {function (this:Object, string): boolean} */
      var html = {}.propertyIsEnumerable;
      /** @type {function (Object, string): (ObjectPropertyDescriptor|undefined)} */
      var getOwnPropertyDescriptor = Object.getOwnPropertyDescriptor;
      /** @type {boolean} */
      var isFunction = getOwnPropertyDescriptor && !html.call({
        1 : 2
      }, 1);
      /** @type {Function} */
      entry.f = isFunction ? function(expectedNumberOfNonCommentArgs) {
        /** @type {(ObjectPropertyDescriptor|undefined)} */
        var property = getOwnPropertyDescriptor(this, expectedNumberOfNonCommentArgs);
        return!!property && property.enumerable;
      } : html;
    }, function(module, dataAndEvents, require) {
      var getActual = require(7);
      var nodes = require(17);
      /** @type {function (this:string, *=, number=): Array.<string>} */
      var split = "".split;
      /** @type {Function} */
      module.exports = getActual(function() {
        return!Object("z").propertyIsEnumerable(0);
      }) ? function(expectedNumberOfNonCommentArgs) {
        return "String" == nodes(expectedNumberOfNonCommentArgs) ? split.call(expectedNumberOfNonCommentArgs, "") : Object(expectedNumberOfNonCommentArgs);
      } : Object;
    }, function(module, dataAndEvents, require) {
      var inspect = require(5);
      var getActual = require(74);
      var cfg = require(32);
      var options = require(10);
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {number} object
       * @return {undefined}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, object) {
        var codeSegments = getActual(object);
        var setter = options.f;
        var callback = cfg.f;
        /** @type {number} */
        var i = 0;
        for (;i < codeSegments.length;i++) {
          var key = codeSegments[i];
          if (!inspect(expectedNumberOfNonCommentArgs, key)) {
            setter(expectedNumberOfNonCommentArgs, key, callback(object, key));
          }
        }
      };
    }, function(module, dataAndEvents, require) {
      var nodes = require(12);
      var argv = require(75);
      var a = require(78);
      var getActual = require(4);
      module.exports = nodes("Reflect", "ownKeys") || function(expectedNumberOfNonCommentArgs) {
        var r = argv.f(getActual(expectedNumberOfNonCommentArgs));
        var f = a.f;
        return f ? r.concat(f(expectedNumberOfNonCommentArgs)) : r;
      };
    }, function(dataAndEvents, entry, toArray) {
      var dataAttr = toArray(49);
      var camelKey = toArray(34).concat("length", "prototype");
      /** @type {function (Object): Array.<string>} */
      entry.f = Object.getOwnPropertyNames || function(expectedNumberOfNonCommentArgs) {
        return dataAttr(expectedNumberOfNonCommentArgs, camelKey);
      };
    }, function(mod, dataAndEvents, require) {
      var getName = require(23);
      var assert = require(33);
      var Event = require(77);
      /**
       * @param {boolean} recurring
       * @return {?}
       */
      var guard = function(recurring) {
        return function(value, searchElement, type) {
          var target;
          var t = getName(value);
          var key = assert(t.length);
          var k = Event(type, key);
          if (recurring && searchElement != searchElement) {
            for (;key > k;) {
              if ((target = t[k++]) != target) {
                return true;
              }
            }
          } else {
            for (;key > k;k++) {
              if ((recurring || k in t) && t[k] === searchElement) {
                return recurring || (k || 0);
              }
            }
          }
          return!recurring && -1;
        };
      };
      mod.exports = {
        includes : guard(true),
        indexOf : guard(false)
      };
    }, function(module, dataAndEvents, require) {
      var getActual = require(22);
      /** @type {function (...[*]): number} */
      var nativeMax = Math.max;
      /** @type {function (...[*]): number} */
      var nativeMin = Math.min;
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {Function} proto
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, proto) {
        var fromIndex = getActual(expectedNumberOfNonCommentArgs);
        return fromIndex < 0 ? nativeMax(fromIndex + proto, 0) : nativeMin(fromIndex, proto);
      };
    }, function(dataAndEvents, object) {
      object.f = Object.getOwnPropertySymbols;
    }, function(module, dataAndEvents, require) {
      var basePrototype = require(51).IteratorPrototype;
      var getActual = require(36);
      var next = require(21);
      var inspect = require(37);
      var nodes = require(19);
      /**
       * @return {?}
       */
      var result = function() {
        return this;
      };
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {Function} proto
       * @param {boolean} data
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, proto, data) {
        /** @type {string} */
        var depth = proto + " Iterator";
        return expectedNumberOfNonCommentArgs.prototype = getActual(basePrototype, {
          next : next(1, data)
        }), inspect(expectedNumberOfNonCommentArgs, depth, false, true), nodes[depth] = result, expectedNumberOfNonCommentArgs;
      };
    }, function(module, dataAndEvents, require) {
      var getActual = require(18);
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs) {
        return Object(getActual(expectedNumberOfNonCommentArgs));
      };
    }, function(module, dataAndEvents, require) {
      var getActual = require(7);
      /** @type {boolean} */
      module.exports = !getActual(function() {
        /**
         * @return {undefined}
         */
        function C() {
        }
        return C.prototype.constructor = null, Object.getPrototypeOf(new C) !== C.prototype;
      });
    }, function(module, dataAndEvents, require) {
      var Block = require(9);
      var callback = require(10);
      var getActual = require(4);
      var forOwn = require(83);
      /** @type {Function} */
      module.exports = Block ? Object.defineProperties : function(expectedNumberOfNonCommentArgs, object) {
        getActual(expectedNumberOfNonCommentArgs);
        var key;
        var keys = forOwn(object);
        var len = keys.length;
        /** @type {number} */
        var j = 0;
        for (;len > j;) {
          callback.f(expectedNumberOfNonCommentArgs, key = keys[j++], object[key]);
        }
        return expectedNumberOfNonCommentArgs;
      };
    }, function(module, dataAndEvents, require) {
      var getActual = require(49);
      var args = require(34);
      /** @type {function (Object): Array.<string>} */
      module.exports = Object.keys || function(expectedNumberOfNonCommentArgs) {
        return getActual(expectedNumberOfNonCommentArgs, args);
      };
    }, function(module, dataAndEvents, require) {
      var getActual = require(11);
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs) {
        if (!getActual(expectedNumberOfNonCommentArgs) && null !== expectedNumberOfNonCommentArgs) {
          throw TypeError("Can't set " + String(expectedNumberOfNonCommentArgs) + " as a prototype");
        }
        return expectedNumberOfNonCommentArgs;
      };
    }, function(dataAndEvents, deepDataAndEvents, require) {
      var events = require(1);
      var types = require(86);
      var data = require(87);
      var debug = require(6);
      var getName = require(3);
      var k = getName("iterator");
      var name = getName("toStringTag");
      var v = data.values;
      var type;
      for (type in types) {
        var constructor = events[type];
        var cache = constructor && constructor.prototype;
        if (cache) {
          if (cache[k] !== v) {
            try {
              debug(cache, k, v);
            } catch (t) {
              cache[k] = v;
            }
          }
          if (cache[name] || debug(cache, name, type), types[type]) {
            var prop;
            for (prop in data) {
              if (cache[prop] !== data[prop]) {
                try {
                  debug(cache, prop, data[prop]);
                } catch (t) {
                  cache[prop] = data[prop];
                }
              }
            }
          }
        }
      }
    }, function(module, dataAndEvents) {
      module.exports = {
        CSSRuleList : 0,
        CSSStyleDeclaration : 0,
        CSSValueList : 0,
        ClientRectList : 0,
        DOMRectList : 0,
        DOMStringList : 0,
        DOMTokenList : 1,
        DataTransferItemList : 0,
        FileList : 0,
        HTMLAllCollection : 0,
        HTMLCollection : 0,
        HTMLFormElement : 0,
        HTMLSelectElement : 0,
        MediaList : 0,
        MimeTypeArray : 0,
        NamedNodeMap : 0,
        NodeList : 1,
        PaintRequestList : 0,
        Plugin : 0,
        PluginArray : 0,
        SVGLengthList : 0,
        SVGNumberList : 0,
        SVGPathSegList : 0,
        SVGPointList : 0,
        SVGStringList : 0,
        SVGTransformList : 0,
        SourceBufferList : 0,
        StyleSheetList : 0,
        TextTrackCueList : 0,
        TextTrackList : 0,
        TouchList : 0
      };
    }, function(module, dataAndEvents, require) {
      var getActual = require(23);
      var isArray = require(88);
      var nodes = require(19);
      var style = require(16);
      var factory = require(47);
      var setStyle = style.set;
      var ease = style.getterFor("Array Iterator");
      module.exports = factory(Array, "Array", function(obj, kind) {
        setStyle(this, {
          type : "Array Iterator",
          target : getActual(obj),
          index : 0,
          kind : kind
        });
      }, function() {
        var e = ease(this);
        var source = e.target;
        var key = e.kind;
        /** @type {number} */
        var i = e.index++;
        return!source || i >= source.length ? (e.target = void 0, {
          value : void 0,
          done : true
        }) : "keys" == key ? {
          value : i,
          done : false
        } : "values" == key ? {
          value : source[i],
          done : false
        } : {
          value : [i, source[i]],
          done : false
        };
      }, "values");
      nodes.Arguments = nodes.Array;
      isArray("keys");
      isArray("values");
      isArray("entries");
    }, function(module, dataAndEvents, _) {
      var wrapped = _(3);
      var today = _(36);
      var _this = _(10);
      var optgroup = wrapped("unscopables");
      var x = Array.prototype;
      if (null == x[optgroup]) {
        _this.f(x, optgroup, {
          configurable : true,
          value : today(null)
        });
      }
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {undefined}
       */
      module.exports = function(expectedNumberOfNonCommentArgs) {
        /** @type {boolean} */
        x[optgroup][expectedNumberOfNonCommentArgs] = true;
      };
    }, function(dataAndEvents, deepDataAndEvents, require) {
      var cb;
      var poll;
      var ret;
      var then;
      var nodes = require(8);
      var Block = require(15);
      var self = require(1);
      var factory = require(12);
      var object = require(54);
      var typeOf = require(13);
      var createObject = require(90);
      var render = require(37);
      var visitor = require(91);
      var isKind = require(11);
      var resolve = require(14);
      var group = require(92);
      var getActual = require(17);
      var getName = require(29);
      var cast = require(24);
      var helper = require(96);
      var indexOf = require(55);
      var setter = require(56).set;
      var assert = require(97);
      var callback = require(58);
      var mixIn = require(98);
      var cfg = require(20);
      var call = require(25);
      var Ember = require(16);
      var compile = require(50);
      var flag = require(3);
      var inspect = require(99);
      var obj = flag("species");
      /** @type {string} */
      var view = "Promise";
      var get = Ember.get;
      var setStyle = Ember.set;
      var toObject = Ember.getterFor(view);
      var value = object;
      var a = self.TypeError;
      var doc = self.document;
      var process = self.process;
      var handler = factory("fetch");
      var fn = cfg.f;
      var org = fn;
      /** @type {boolean} */
      var domain = "process" == getActual(process);
      /** @type {boolean} */
      var K = !!(doc && (doc.createEvent && self.dispatchEvent));
      var element = compile(view, function() {
        if (!(getName(value) !== String(value))) {
          if (66 === inspect) {
            return true;
          }
          if (!domain && "function" != typeof PromiseRejectionEvent) {
            return true;
          }
        }
        if (Block && !value.prototype.finally) {
          return true;
        }
        if (inspect >= 51 && /native code/.test(value)) {
          return false;
        }
        var me = value.resolve(1);
        /**
         * @param {?} onComplete
         * @return {undefined}
         */
        var finish = function(onComplete) {
          onComplete(function() {
          }, function() {
          });
        };
        return(me.constructor = {})[obj] = finish, !(me.then(function() {
        }) instanceof finish);
      });
      var env = element || !helper(function(isXML) {
        value.all(isXML).catch(function() {
        });
      });
      /**
       * @param {?} val
       * @return {?}
       */
      var isArray = function(val) {
        var then;
        return!(!isKind(val) || "function" != typeof(then = val.then)) && then;
      };
      /**
       * @param {number} e
       * @param {Object} data
       * @param {boolean} recurring
       * @return {undefined}
       */
      var check = function(e, data, recurring) {
        if (!data.notified) {
          /** @type {boolean} */
          data.notified = true;
          var items = data.reactions;
          assert(function() {
            var raw = data.value;
            /** @type {boolean} */
            var caseSensitive = 1 == data.state;
            /** @type {number} */
            var index = 0;
            for (;items.length > index;) {
              var expectedNumberOfNonCommentArgs;
              var cb;
              var c;
              var result = items[index++];
              var text = caseSensitive ? result.ok : result.fail;
              var resolve = result.resolve;
              var callback = result.reject;
              var domain = result.domain;
              try {
                if (text) {
                  if (!caseSensitive) {
                    if (2 === data.rejection) {
                      finish(e, data);
                    }
                    /** @type {number} */
                    data.rejection = 1;
                  }
                  if (true === text) {
                    expectedNumberOfNonCommentArgs = raw;
                  } else {
                    if (domain) {
                      domain.enter();
                    }
                    expectedNumberOfNonCommentArgs = text(raw);
                    if (domain) {
                      domain.exit();
                      /** @type {boolean} */
                      c = true;
                    }
                  }
                  if (expectedNumberOfNonCommentArgs === result.promise) {
                    callback(a("Promise-chain cycle"));
                  } else {
                    if (cb = isArray(expectedNumberOfNonCommentArgs)) {
                      cb.call(expectedNumberOfNonCommentArgs, resolve, callback);
                    } else {
                      resolve(expectedNumberOfNonCommentArgs);
                    }
                  }
                } else {
                  callback(raw);
                }
              } catch (STOP) {
                if (domain) {
                  if (!c) {
                    domain.exit();
                  }
                }
                callback(STOP);
              }
            }
            /** @type {Array} */
            data.reactions = [];
            /** @type {boolean} */
            data.notified = false;
            if (recurring) {
              if (!data.rejection) {
                done(e, data);
              }
            }
          });
        }
      };
      /**
       * @param {string} type
       * @param {number} target
       * @param {?} options
       * @return {undefined}
       */
      var triggerEvent = function(type, target, options) {
        var event;
        var getXYfromEvent;
        if (K) {
          /** @type {number} */
          (event = doc.createEvent("Event")).promise = target;
          event.reason = options;
          event.initEvent(type, false, true);
          self.dispatchEvent(event);
        } else {
          event = {
            promise : target,
            reason : options
          };
        }
        if (getXYfromEvent = self["on" + type]) {
          getXYfromEvent(event);
        } else {
          if ("unhandledrejection" === type) {
            mixIn("Unhandled promise rejection", options);
          }
        }
      };
      /**
       * @param {number} data
       * @param {Object} u
       * @return {undefined}
       */
      var done = function(data, u) {
        setter.call(self, function() {
          var object;
          var args = u.value;
          if (map(u) && (object = call(function() {
            if (domain) {
              process.emit("unhandledRejection", args, data);
            } else {
              triggerEvent("unhandledrejection", data, args);
            }
          }), u.rejection = domain || map(u) ? 2 : 1, object.error)) {
            throw object.value;
          }
        });
      };
      /**
       * @param {Object} v
       * @return {?}
       */
      var map = function(v) {
        return 1 !== v.rejection && !v.parent;
      };
      /**
       * @param {number} err
       * @param {Object} buffer
       * @return {undefined}
       */
      var finish = function(err, buffer) {
        setter.call(self, function() {
          if (domain) {
            process.emit("rejectionHandled", err);
          } else {
            triggerEvent("rejectionhandled", err, buffer.value);
          }
        });
      };
      /**
       * @param {Function} callback
       * @param {number} value
       * @param {?} obj
       * @param {Object} context
       * @return {?}
       */
      var $ = function(callback, value, obj, context) {
        return function(arg) {
          callback(value, obj, arg, context);
        };
      };
      /**
       * @param {number} event
       * @param {Object} d
       * @param {?} f
       * @param {Object} i
       * @return {undefined}
       */
      var next = function(event, d, f, i) {
        if (!d.done) {
          /** @type {boolean} */
          d.done = true;
          if (i) {
            /** @type {Object} */
            d = i;
          }
          d.value = f;
          /** @type {number} */
          d.state = 2;
          check(event, d, true);
        }
      };
      /**
       * @param {number} name
       * @param {Object} callback
       * @param {Object} t
       * @param {Object} expected
       * @return {undefined}
       */
      var test = function(name, callback, t, expected) {
        if (!callback.done) {
          /** @type {boolean} */
          callback.done = true;
          if (expected) {
            /** @type {Object} */
            callback = expected;
          }
          try {
            if (name === t) {
              throw a("Promise can't be resolved itself");
            }
            var self = isArray(t);
            if (self) {
              assert(function() {
                var res = {
                  done : false
                };
                try {
                  self.call(t, $(test, name, res, callback), $(next, name, res, callback));
                } catch (fromIndex) {
                  next(name, res, fromIndex, callback);
                }
              });
            } else {
              /** @type {Object} */
              callback.value = t;
              /** @type {number} */
              callback.state = 1;
              check(name, callback, false);
            }
          } catch (fromIndex) {
            next(name, {
              done : false
            }, fromIndex, callback);
          }
        }
      };
      if (element) {
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {undefined}
         */
        value = function(expectedNumberOfNonCommentArgs) {
          group(this, value, view);
          resolve(expectedNumberOfNonCommentArgs);
          cb.call(this);
          var res = get(this);
          try {
            expectedNumberOfNonCommentArgs($(test, this, res), $(next, this, res));
          } catch (fromIndex) {
            next(this, res, fromIndex);
          }
        };
        (cb = function(stats) {
          setStyle(this, {
            type : view,
            done : false,
            notified : false,
            parent : false,
            reactions : [],
            rejection : false,
            state : 0,
            value : void 0
          });
        }).prototype = createObject(value.prototype, {
          /**
           * @param {Object} opt_attributes
           * @param {Object} onReject
           * @return {?}
           */
          then : function(opt_attributes, onReject) {
            var self = toObject(this);
            var response = fn(indexOf(this, value));
            return response.ok = "function" != typeof opt_attributes || opt_attributes, response.fail = "function" == typeof onReject && onReject, response.domain = domain ? process.domain : void 0, self.parent = true, self.reactions.push(response), 0 != self.state && check(this, self, false), response.promise;
          },
          /**
           * @param {Object} callback
           * @return {?}
           */
          catch : function(callback) {
            return this.then(void 0, callback);
          }
        });
        /**
         * @return {undefined}
         */
        poll = function() {
          var results = new cb;
          var suiteView = get(results);
          this.promise = results;
          this.resolve = $(test, results, suiteView);
          this.reject = $(next, results, suiteView);
        };
        /** @type {function (number): ?} */
        cfg.f = fn = function(expectedNumberOfNonCommentArgs) {
          return expectedNumberOfNonCommentArgs === value || expectedNumberOfNonCommentArgs === ret ? new poll(expectedNumberOfNonCommentArgs) : org(expectedNumberOfNonCommentArgs);
        };
        if (!Block) {
          if (!("function" != typeof object)) {
            then = object.prototype.then;
            typeOf(object.prototype, "then", function(attributes, onReject) {
              var self = this;
              return(new value(function(mapper, reject) {
                then.call(self, mapper, reject);
              })).then(attributes, onReject);
            }, {
              unsafe : true
            });
            if ("function" == typeof handler) {
              nodes({
                global : true,
                enumerable : true,
                forced : true
              }, {
                /**
                 * @param {?} pool
                 * @return {?}
                 */
                fetch : function(pool) {
                  return callback(value, handler.apply(self, arguments));
                }
              });
            }
          }
        }
      }
      nodes({
        global : true,
        wrap : true,
        forced : element
      }, {
        Promise : value
      });
      render(value, view, false, true);
      visitor(view);
      ret = factory(view);
      nodes({
        target : view,
        stat : true,
        forced : element
      }, {
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        reject : function(expectedNumberOfNonCommentArgs) {
          var callback = fn(this);
          return callback.reject.call(void 0, expectedNumberOfNonCommentArgs), callback.promise;
        }
      });
      nodes({
        target : view,
        stat : true,
        forced : Block || element
      }, {
        /**
         * @param {number} expectedNumberOfNonCommentArgs
         * @return {?}
         */
        resolve : function(expectedNumberOfNonCommentArgs) {
          return callback(Block && this === ret ? value : this, expectedNumberOfNonCommentArgs);
        }
      });
      nodes({
        target : view,
        stat : true,
        forced : env
      }, {
        /**
         * @param {?} value
         * @return {?}
         */
        all : function(value) {
          var view = this;
          var result = fn(view);
          var resolver = result.resolve;
          var reject = result.reject;
          var r = call(function() {
            var fn = resolve(view.resolve);
            /** @type {Array} */
            var expectedNumberOfNonCommentArgs = [];
            /** @type {number} */
            var s = 0;
            /** @type {number} */
            var a = 1;
            cast(value, function(locals) {
              /** @type {number} */
              var unlock = s++;
              /** @type {boolean} */
              var c = false;
              expectedNumberOfNonCommentArgs.push(void 0);
              a++;
              fn.call(view, locals).then(function(data) {
                if (!c) {
                  /** @type {boolean} */
                  c = true;
                  expectedNumberOfNonCommentArgs[unlock] = data;
                  if (!--a) {
                    resolver(expectedNumberOfNonCommentArgs);
                  }
                }
              }, reject);
            });
            if (!--a) {
              resolver(expectedNumberOfNonCommentArgs);
            }
          });
          return r.error && reject(r.value), result.promise;
        },
        /**
         * @param {?} array
         * @return {?}
         */
        race : function(array) {
          var expectedNumberOfNonCommentArgs = this;
          var result = fn(expectedNumberOfNonCommentArgs);
          var reject = result.reject;
          var r = call(function() {
            var fn = resolve(expectedNumberOfNonCommentArgs.resolve);
            cast(array, function(locals) {
              fn.call(expectedNumberOfNonCommentArgs, locals).then(result.resolve, reject);
            });
          });
          return r.error && reject(r.value), result.promise;
        }
      });
    }, function(module, dataAndEvents, $sanitize) {
      var dataAttr = $sanitize(13);
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {Object} object
       * @param {boolean} data
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, object, data) {
        var name;
        for (name in object) {
          dataAttr(expectedNumberOfNonCommentArgs, name, object[name], data);
        }
        return expectedNumberOfNonCommentArgs;
      };
    }, function(module, dataAndEvents, require) {
      var inspect = require(12);
      var argv = require(10);
      var getActual = require(3);
      var Block = require(9);
      var optgroup = getActual("species");
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {undefined}
       */
      module.exports = function(expectedNumberOfNonCommentArgs) {
        var str = inspect(expectedNumberOfNonCommentArgs);
        var len = argv.f;
        if (Block) {
          if (str) {
            if (!str[optgroup]) {
              len(str, optgroup, {
                configurable : true,
                /**
                 * @return {?}
                 */
                get : function() {
                  return this;
                }
              });
            }
          }
        }
      };
    }, function(module, dataAndEvents) {
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {Function} proto
       * @param {string} value
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, proto, value) {
        if (!(expectedNumberOfNonCommentArgs instanceof proto)) {
          throw TypeError("Incorrect " + (value ? value + " " : "") + "invocation");
        }
        return expectedNumberOfNonCommentArgs;
      };
    }, function(module, dataAndEvents, require) {
      var getName = require(3);
      var nodes = require(19);
      var name = getName("iterator");
      var ap = Array.prototype;
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs) {
        return void 0 !== expectedNumberOfNonCommentArgs && (nodes.Array === expectedNumberOfNonCommentArgs || ap[name] === expectedNumberOfNonCommentArgs);
      };
    }, function(module, dataAndEvents, require) {
      var getActual = require(46);
      var Block = require(19);
      var expression = require(3)("iterator");
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs) {
        if (null != expectedNumberOfNonCommentArgs) {
          return expectedNumberOfNonCommentArgs[expression] || (expectedNumberOfNonCommentArgs["@@iterator"] || Block[getActual(expectedNumberOfNonCommentArgs)]);
        }
      };
    }, function(module, dataAndEvents, require) {
      var inspect = require(4);
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {Function} callback
       * @param {Object} data
       * @param {?} err
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, callback, data, err) {
        try {
          return err ? callback(inspect(data)[0], data[1]) : callback(data);
        } catch (e) {
          var closure = expectedNumberOfNonCommentArgs.return;
          throw void 0 !== closure && inspect(closure.call(expectedNumberOfNonCommentArgs)), e;
        }
      };
    }, function(module, dataAndEvents, $sanitize) {
      var prop = $sanitize(3)("iterator");
      /** @type {boolean} */
      var property = false;
      try {
        /** @type {number} */
        var o = 0;
        var args = {
          /**
           * @return {?}
           */
          next : function() {
            return{
              done : !!o++
            };
          },
          /**
           * @return {undefined}
           */
          return : function() {
            /** @type {boolean} */
            property = true;
          }
        };
        /**
         * @return {?}
         */
        args[prop] = function() {
          return this;
        };
        Array.from(args, function() {
          throw 2;
        });
      } catch (t) {
      }
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {Function} object
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, object) {
        if (!object && !property) {
          return false;
        }
        /** @type {boolean} */
        var str = false;
        try {
          var originalEvent = {};
          /**
           * @return {?}
           */
          originalEvent[prop] = function() {
            return{
              /**
               * @return {?}
               */
              next : function() {
                return{
                  done : str = true
                };
              }
            };
          };
          expectedNumberOfNonCommentArgs(originalEvent);
        } catch (t) {
        }
        return str;
      };
    }, function(module, dataAndEvents, require) {
      var flush;
      var current;
      var ret;
      var send;
      var iterations;
      var node;
      var value;
      var then;
      var expectedNumberOfNonCommentArgs = require(1);
      var requestAnimationFrame = require(32).f;
      var getActual = require(17);
      var setter = require(56).set;
      var Block = require(57);
      var BrowserMutationObserver = expectedNumberOfNonCommentArgs.MutationObserver || expectedNumberOfNonCommentArgs.WebKitMutationObserver;
      var process = expectedNumberOfNonCommentArgs.process;
      var p1 = expectedNumberOfNonCommentArgs.Promise;
      /** @type {boolean} */
      var b = "process" == getActual(process);
      var id = requestAnimationFrame(expectedNumberOfNonCommentArgs, "queueMicrotask");
      var exports = id && id.value;
      if (!exports) {
        /**
         * @return {undefined}
         */
        flush = function() {
          var d;
          var keys;
          if (b) {
            if (d = process.domain) {
              d.exit();
            }
          }
          for (;current;) {
            keys = current.fn;
            current = current.next;
            try {
              keys();
            } catch (t) {
              throw current ? send() : ret = void 0, t;
            }
          }
          ret = void 0;
          if (d) {
            d.enter();
          }
        };
        if (b) {
          /**
           * @return {undefined}
           */
          send = function() {
            process.nextTick(flush);
          };
        } else {
          if (BrowserMutationObserver && !Block) {
            /** @type {boolean} */
            iterations = true;
            /** @type {Text} */
            node = document.createTextNode("");
            (new BrowserMutationObserver(flush)).observe(node, {
              characterData : true
            });
            /**
             * @return {undefined}
             */
            send = function() {
              /** @type {boolean} */
              node.data = iterations = !iterations;
            };
          } else {
            if (p1 && p1.resolve) {
              value = p1.resolve(void 0);
              then = value.then;
              /**
               * @return {undefined}
               */
              send = function() {
                then.call(value, flush);
              };
            } else {
              /**
               * @return {undefined}
               */
              send = function() {
                setter.call(expectedNumberOfNonCommentArgs, flush);
              };
            }
          }
        }
      }
      module.exports = exports || function(expectedNumberOfNonCommentArgs) {
        var next = {
          fn : expectedNumberOfNonCommentArgs,
          next : void 0
        };
        if (ret) {
          ret.next = next;
        }
        if (!current) {
          current = next;
          send();
        }
        ret = next;
      };
    }, function(module, dataAndEvents, Event) {
      var self = Event(1);
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {Function} actual
       * @return {undefined}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, actual) {
        var test = self.console;
        if (test) {
          if (test.error) {
            if (1 === arguments.length) {
              test.error(expectedNumberOfNonCommentArgs);
            } else {
              test.error(expectedNumberOfNonCommentArgs, actual);
            }
          }
        }
      };
    }, function(module, dataAndEvents, require) {
      var groupedSelectors;
      var selector;
      var global = require(1);
      var expect = require(39);
      var doc = global.process;
      var docElement = doc && doc.versions;
      var uHostName = docElement && docElement.v8;
      if (uHostName) {
        selector = (groupedSelectors = uHostName.split("."))[0] + groupedSelectors[1];
      } else {
        if (expect) {
          if (!(groupedSelectors = expect.match(/Edge\/(\d+)/)) || groupedSelectors[1] >= 74) {
            if (groupedSelectors = expect.match(/Chrome\/(\d+)/)) {
              selector = groupedSelectors[1];
            }
          }
        }
      }
      module.exports = selector && +selector;
    }, function(dataAndEvents, deepDataAndEvents, require) {
      var getActual = require(8);
      var Block = require(15);
      var b = require(54);
      var nodes = require(7);
      var helper = require(12);
      var inspect = require(55);
      var expect = require(58);
      var createObject = require(13);
      getActual({
        target : "Promise",
        proto : true,
        real : true,
        forced : !!b && nodes(function() {
          b.prototype.finally.call({
            /**
             * @return {undefined}
             */
            then : function() {
            }
          }, function() {
          });
        })
      }, {
        /**
         * @param {Function} callback
         * @return {?}
         */
        finally : function(callback) {
          var str = inspect(this, helper("Promise"));
          /** @type {boolean} */
          var fn = "function" == typeof callback;
          return this.then(fn ? function(dataAndEvents) {
            return expect(str, callback()).then(function() {
              return dataAndEvents;
            });
          } : callback, fn ? function(dataAndEvents) {
            return expect(str, callback()).then(function() {
              throw dataAndEvents;
            });
          } : callback);
        }
      });
      if (!Block) {
        if (!("function" != typeof b)) {
          if (!b.prototype.finally) {
            createObject(b.prototype, "finally", helper("Promise").prototype.finally);
          }
        }
      }
    }, function(dataAndEvents, deepDataAndEvents, require) {
      var getActual = require(8);
      var Block = require(9);
      var tmpl = require(35);
      var fn = require(53);
      var create = require(36);
      var module = require(10);
      var helper = require(21);
      var Event = require(24);
      var expect = require(6);
      var request = require(16);
      var callback = request.set;
      var transferFlags = request.getterFor("AggregateError");
      /**
       * @param {string} type
       * @param {string} message
       * @return {?}
       */
      var error = function(type, message) {
        var nodes = this;
        if (!(nodes instanceof error)) {
          return new error(type, message);
        }
        if (fn) {
          nodes = fn(new Error(message), tmpl(nodes));
        }
        /** @type {Array} */
        var errors = [];
        return Event(type, errors.push, errors), Block ? callback(nodes, {
          errors : errors,
          type : "AggregateError"
        }) : nodes.errors = errors, void 0 !== message && expect(nodes, "message", String(message)), nodes;
      };
      error.prototype = create(Error.prototype, {
        constructor : helper(5, error),
        message : helper(5, ""),
        name : helper(5, "AggregateError")
      });
      if (Block) {
        module.f(error.prototype, "errors", {
          /**
           * @return {?}
           */
          get : function() {
            return transferFlags(this).errors;
          },
          configurable : true
        });
      }
      getActual({
        global : true
      }, {
        /** @type {function (string, string): ?} */
        AggregateError : error
      });
    }, function(dataAndEvents, deepDataAndEvents, $sanitize) {
      $sanitize(59);
    }, function(dataAndEvents, deepDataAndEvents, topic) {
      var throttledUpdate = topic(8);
      var out = topic(20);
      var MAP = topic(25);
      throttledUpdate({
        target : "Promise",
        stat : true
      }, {
        /**
         * @param {?} t
         * @return {?}
         */
        try : function(t) {
          var response = out.f(this);
          var g = MAP(t);
          return(g.error ? response.reject : response.resolve)(g.value), response.promise;
        }
      });
    }, function(dataAndEvents, deepDataAndEvents, require) {
      var nodes = require(8);
      var qMap = require(14);
      var helper = require(12);
      var argv = require(20);
      var f = require(25);
      var getActual = require(24);
      nodes({
        target : "Promise",
        stat : true
      }, {
        /**
         * @param {?} obj
         * @return {?}
         */
        any : function(obj) {
          var expectedNumberOfNonCommentArgs = this;
          var original = argv.f(expectedNumberOfNonCommentArgs);
          var button = original.resolve;
          var iterator = original.reject;
          var x = f(function() {
            var native_method = qMap(expectedNumberOfNonCommentArgs.resolve);
            /** @type {Array} */
            var tempData = [];
            /** @type {number} */
            var a = 0;
            /** @type {number} */
            var f = 1;
            /** @type {boolean} */
            var h = false;
            getActual(obj, function(mapper) {
              /** @type {number} */
              var unlock = a++;
              /** @type {boolean} */
              var u = false;
              tempData.push(void 0);
              f++;
              native_method.call(expectedNumberOfNonCommentArgs, mapper).then(function(expectedNumberOfNonCommentArgs) {
                if (!u) {
                  if (!h) {
                    /** @type {boolean} */
                    h = true;
                    button(expectedNumberOfNonCommentArgs);
                  }
                }
              }, function(data) {
                if (!u) {
                  if (!h) {
                    /** @type {boolean} */
                    u = true;
                    tempData[unlock] = data;
                    if (!--f) {
                      iterator(new (helper("AggregateError"))(tempData, "No one promise resolved"));
                    }
                  }
                }
              });
            });
            if (!--f) {
              iterator(new (helper("AggregateError"))(tempData, "No one promise resolved"));
            }
          });
          return x.error && iterator(x.value), original.promise;
        }
      });
    }, function(module, dataAndEvents, topic) {
      var out = topic(106);
      module.exports = out;
    }, function(module, dataAndEvents, require) {
      require(107);
      var factory = require(111);
      module.exports = factory("String", "padStart");
    }, function(dataAndEvents, deepDataAndEvents, require) {
      var getActual = require(8);
      var oldStart = require(108).start;
      getActual({
        target : "String",
        proto : true,
        forced : require(110)
      }, {
        /**
         * @param {number} curr
         * @return {?}
         */
        padStart : function(curr) {
          return oldStart(this, curr, arguments.length > 1 ? arguments[1] : void 0);
        }
      });
    }, function(module, dataAndEvents, getCallback) {
      var cb = getCallback(33);
      var ostring = getCallback(109);
      var callback = getCallback(18);
      /** @type {function (*): number} */
      var ceil = Math.ceil;
      /**
       * @param {boolean} recurring
       * @return {?}
       */
      var write = function(recurring) {
        return function(value, outErr, string) {
          var width;
          var r;
          /** @type {string} */
          var t = String(callback(value));
          /** @type {number} */
          var left = t.length;
          /** @type {string} */
          var it = void 0 === string ? " " : String(string);
          var right = cb(outErr);
          return right <= left || "" == it ? t : (width = right - left, (r = ostring.call(it, ceil(width / it.length))).length > width && (r = r.slice(0, width)), recurring ? t + r : r + t);
        };
      };
      module.exports = {
        start : write(false),
        end : write(true)
      };
    }, function(module, dataAndEvents, require) {
      var getActual = require(22);
      var nodes = require(18);
      module.exports = "".repeat || function(expectedNumberOfNonCommentArgs) {
        /** @type {string} */
        var s = String(nodes(this));
        /** @type {string} */
        var slarge = "";
        var actual = getActual(expectedNumberOfNonCommentArgs);
        if (actual < 0 || actual == 1 / 0) {
          throw RangeError("Wrong number of repetitions");
        }
        for (;actual > 0;(actual >>>= 1) && (s += s)) {
          if (1 & actual) {
            slarge += s;
          }
        }
        return slarge;
      };
    }, function(module, dataAndEvents, getName) {
      var name = getName(39);
      /** @type {boolean} */
      module.exports = /Version\/10\.\d+(\.\d+)?( Mobile\/\w+)? Safari\//.test(name);
    }, function(module, dataAndEvents, keys) {
      var props = keys(1);
      var ondata = keys(38);
      var call = Function.call;
      /**
       * @param {number} expectedNumberOfNonCommentArgs
       * @param {Function} proto
       * @param {boolean} data
       * @return {?}
       */
      module.exports = function(expectedNumberOfNonCommentArgs, proto, data) {
        return ondata(call, props[expectedNumberOfNonCommentArgs].prototype[proto], data);
      };
    }, function(module, dataAndEvents, deepDataAndEvents) {
      var html5 = function(self) {
        /**
         * @param {Function} object
         * @param {Function} recurring
         * @param {string} name
         * @param {Array} data
         * @return {?}
         */
        function cb(object, recurring, name, data) {
          var Type = recurring && recurring.prototype instanceof superClass ? recurring : superClass;
          /** @type {Object} */
          var self = Object.create(Type.prototype);
          var pdataCur = new Class(data || []);
          return self._invoke = function(config, path, data) {
            /** @type {string} */
            var arLength = "suspendedStart";
            return function(type, key) {
              if ("executing" === arLength) {
                throw new Error("Generator is already running");
              }
              if ("completed" === arLength) {
                if ("throw" === type) {
                  throw key;
                }
                return gotoNext();
              }
              /** @type {string} */
              data.method = type;
              data.arg = key;
              for (;;) {
                var name = data.delegate;
                if (name) {
                  var v = next(name, data);
                  if (v) {
                    if (v === value) {
                      continue;
                    }
                    return v;
                  }
                }
                if ("next" === data.method) {
                  data.sent = data._sent = data.arg;
                } else {
                  if ("throw" === data.method) {
                    if ("suspendedStart" === arLength) {
                      throw arLength = "completed", data.arg;
                    }
                    data.dispatchException(data.arg);
                  } else {
                    if ("return" === data.method) {
                      data.abrupt("return", data.arg);
                    }
                  }
                }
                /** @type {string} */
                arLength = "executing";
                var item = debug(config, path, data);
                if ("normal" === item.type) {
                  if (arLength = data.done ? "completed" : "suspendedYield", item.arg === value) {
                    continue;
                  }
                  return{
                    value : item.arg,
                    done : data.done
                  };
                }
                if ("throw" === item.type) {
                  /** @type {string} */
                  arLength = "completed";
                  /** @type {string} */
                  data.method = "throw";
                  data.arg = item.arg;
                }
              }
            };
          }(object, name, pdataCur), self;
        }
        /**
         * @param {Function} func
         * @param {?} arg
         * @param {Object} text
         * @return {?}
         */
        function debug(func, arg, text) {
          try {
            return{
              type : "normal",
              arg : func.call(arg, text)
            };
          } catch (param) {
            return{
              type : "throw",
              arg : param
            };
          }
        }
        /**
         * @return {undefined}
         */
        function superClass() {
        }
        /**
         * @return {undefined}
         */
        function method() {
        }
        /**
         * @return {undefined}
         */
        function prop() {
        }
        /**
         * @param {Object} obj
         * @return {undefined}
         */
        function onComplete(obj) {
          ["next", "throw", "return"].forEach(function(val) {
            /**
             * @param {Object} event
             * @return {?}
             */
            obj[val] = function(event) {
              return this._invoke(val, event);
            };
          });
        }
        /**
         * @param {Object} d
         * @return {undefined}
         */
        function handler(d) {
          var promise;
          /**
           * @param {string} ms
           * @param {Object} action
           * @return {?}
           */
          this._invoke = function(ms, action) {
            /**
             * @return {?}
             */
            function reject() {
              return new Promise(function(next_callback, opt_e) {
                !function parse(str, body, callback, val) {
                  var data = debug(d[str], d, body);
                  if ("throw" !== data.type) {
                    var a = data.arg;
                    var expectedNumberOfNonCommentArgs = a.value;
                    return expectedNumberOfNonCommentArgs && ("object" == typeof expectedNumberOfNonCommentArgs && hasOwnProperty.call(expectedNumberOfNonCommentArgs, "__await")) ? Promise.resolve(expectedNumberOfNonCommentArgs.__await).then(function(w) {
                      parse("next", w, callback, val);
                    }, function(w) {
                      parse("throw", w, callback, val);
                    }) : Promise.resolve(expectedNumberOfNonCommentArgs).then(function(e) {
                      a.value = e;
                      callback(a);
                    }, function(w) {
                      return parse("throw", w, callback, val);
                    });
                  }
                  val(data.arg);
                }(ms, action, next_callback, opt_e);
              });
            }
            return promise = promise ? promise.then(reject, reject) : reject();
          };
        }
        /**
         * @param {?} options
         * @param {Object} data
         * @return {?}
         */
        function next(options, data) {
          var obj = options.iterator[data.method];
          if (void 0 === obj) {
            if (data.delegate = null, "throw" === data.method) {
              if (options.iterator.return && (data.method = "return", data.arg = void 0, next(options, data), "throw" === data.method)) {
                return value;
              }
              /** @type {string} */
              data.method = "throw";
              /** @type {TypeError} */
              data.arg = new TypeError("The iterator does not provide a 'throw' method");
            }
            return value;
          }
          var self = debug(obj, options.iterator, data.arg);
          if ("throw" === self.type) {
            return data.method = "throw", data.arg = self.arg, data.delegate = null, value;
          }
          var v = self.arg;
          return v ? v.done ? (data[options.resultName] = v.value, data.next = options.nextLoc, "return" !== data.method && (data.method = "next", data.arg = void 0), data.delegate = null, value) : v : (data.method = "throw", data.arg = new TypeError("iterator result is not an object"), data.delegate = null, value);
        }
        /**
         * @param {Array} property
         * @return {undefined}
         */
        function addProperty(property) {
          var copies = {
            tryLoc : property[0]
          };
          if (1 in property) {
            copies.catchLoc = property[1];
          }
          if (2 in property) {
            copies.finallyLoc = property[2];
            copies.afterLoc = property[3];
          }
          this.tryEntries.push(copies);
        }
        /**
         * @param {?} httpServer
         * @return {undefined}
         */
        function start(httpServer) {
          var me = httpServer.completion || {};
          /** @type {string} */
          me.type = "normal";
          delete me.arg;
          httpServer.completion = me;
        }
        /**
         * @param {Array} key
         * @return {undefined}
         */
        function Class(key) {
          /** @type {Array} */
          this.tryEntries = [{
            tryLoc : "root"
          }];
          key.forEach(addProperty, this);
          this.reset(true);
        }
        /**
         * @param {Array} value
         * @return {?}
         */
        function val(value) {
          if (value) {
            var values = value[key];
            if (values) {
              return values.call(value);
            }
            if ("function" == typeof value.next) {
              return value;
            }
            if (!isNaN(value.length)) {
              /** @type {number} */
              var index = -1;
              /**
               * @return {?}
               */
              var list = function s() {
                for (;++index < value.length;) {
                  if (hasOwnProperty.call(value, index)) {
                    return s.value = value[index], s.done = false, s;
                  }
                }
                return s.value = void 0, s.done = true, s;
              };
              return list.next = list;
            }
          }
          return{
            /** @type {function (): ?} */
            next : gotoNext
          };
        }
        /**
         * @return {?}
         */
        function gotoNext() {
          return{
            value : void 0,
            done : true
          };
        }
        var ObjProto = Object.prototype;
        /** @type {function (this:Object, *): boolean} */
        var hasOwnProperty = ObjProto.hasOwnProperty;
        var container = "function" == typeof Symbol ? Symbol : {};
        var key = container.iterator || "@@iterator";
        var i = container.asyncIterator || "@@asyncIterator";
        var p = container.toStringTag || "@@toStringTag";
        /** @type {function (Function, Function, string, Array): ?} */
        self.wrap = cb;
        var value = {};
        var expectedNumberOfNonCommentArgs = {};
        /**
         * @return {?}
         */
        expectedNumberOfNonCommentArgs[key] = function() {
          return this;
        };
        /** @type {function (Object): (Object|null)} */
        var getPrototypeOf = Object.getPrototypeOf;
        /** @type {(Object|null)} */
        var ctor = getPrototypeOf && getPrototypeOf(getPrototypeOf(val([])));
        if (ctor) {
          if (ctor !== ObjProto) {
            if (hasOwnProperty.call(ctor, key)) {
              /** @type {Object} */
              expectedNumberOfNonCommentArgs = ctor;
            }
          }
        }
        /** @type {Object} */
        var base = prop.prototype = superClass.prototype = Object.create(expectedNumberOfNonCommentArgs);
        return method.prototype = base.constructor = prop, prop.constructor = method, prop[p] = method.displayName = "GeneratorFunction", self.isGeneratorFunction = function(recurring) {
          /** @type {(Function|boolean|null)} */
          var func = "function" == typeof recurring && recurring.constructor;
          return!!func && (func === method || "GeneratorFunction" === (func.displayName || func.name));
        }, self.mark = function(expectedNumberOfNonCommentArgs) {
          return Object.setPrototypeOf ? Object.setPrototypeOf(expectedNumberOfNonCommentArgs, prop) : (expectedNumberOfNonCommentArgs.__proto__ = prop, p in expectedNumberOfNonCommentArgs || (expectedNumberOfNonCommentArgs[p] = "GeneratorFunction")), expectedNumberOfNonCommentArgs.prototype = Object.create(base), expectedNumberOfNonCommentArgs;
        }, self.awrap = function(dataAndEvents) {
          return{
            __await : dataAndEvents
          };
        }, onComplete(handler.prototype), handler.prototype[i] = function() {
          return this;
        }, self.AsyncIterator = handler, self.async = function(str, recurring, ids, x) {
          var stream = new handler(cb(str, recurring, ids, x));
          return self.isGeneratorFunction(recurring) ? stream : stream.next().then(function(d) {
            return d.done ? d.value : stream.next();
          });
        }, onComplete(base), base[p] = "Generator", base[key] = function() {
          return this;
        }, base.toString = function() {
          return "[object Generator]";
        }, self.keys = function(expectedNumberOfNonCommentArgs) {
          /** @type {Array} */
          var eventPath = [];
          var cur;
          for (cur in expectedNumberOfNonCommentArgs) {
            eventPath.push(cur);
          }
          return eventPath.reverse(), function init() {
            for (;eventPath.length;) {
              var result = eventPath.pop();
              if (result in expectedNumberOfNonCommentArgs) {
                return init.value = result, init.done = false, init;
              }
            }
            return init.done = true, init;
          };
        }, self.values = val, Class.prototype = {
          /** @type {function (Array): undefined} */
          constructor : Class,
          /**
           * @param {boolean} dataAndEvents
           * @return {undefined}
           */
          reset : function(dataAndEvents) {
            if (this.prev = 0, this.next = 0, this.sent = this._sent = void 0, this.done = false, this.delegate = null, this.method = "next", this.arg = void 0, this.tryEntries.forEach(start), !dataAndEvents) {
              var header;
              for (header in this) {
                if ("t" === header.charAt(0)) {
                  if (hasOwnProperty.call(this, header)) {
                    if (!isNaN(+header.slice(1))) {
                      this[header] = void 0;
                    }
                  }
                }
              }
            }
          },
          /**
           * @return {?}
           */
          stop : function() {
            /** @type {boolean} */
            this.done = true;
            var me = this.tryEntries[0].completion;
            if ("throw" === me.type) {
              throw me.arg;
            }
            return this.rval;
          },
          /**
           * @param {?} arg
           * @return {?}
           */
          dispatchException : function(arg) {
            /**
             * @param {?} el
             * @param {boolean} signal_eof
             * @return {?}
             */
            function next(el, signal_eof) {
              return op.type = "throw", op.arg = arg, opts.next = el, signal_eof && (opts.method = "next", opts.arg = void 0), !!signal_eof;
            }
            if (this.done) {
              throw arg;
            }
            var opts = this;
            /** @type {number} */
            var s = this.tryEntries.length - 1;
            for (;s >= 0;--s) {
              var self = this.tryEntries[s];
              var op = self.completion;
              if ("root" === self.tryLoc) {
                return next("end");
              }
              if (self.tryLoc <= this.prev) {
                /** @type {boolean} */
                var format = hasOwnProperty.call(self, "catchLoc");
                /** @type {boolean} */
                var useFormat = hasOwnProperty.call(self, "finallyLoc");
                if (format && useFormat) {
                  if (this.prev < self.catchLoc) {
                    return next(self.catchLoc, true);
                  }
                  if (this.prev < self.finallyLoc) {
                    return next(self.finallyLoc);
                  }
                } else {
                  if (format) {
                    if (this.prev < self.catchLoc) {
                      return next(self.catchLoc, true);
                    }
                  } else {
                    if (!useFormat) {
                      throw new Error("try statement without catch or finally");
                    }
                    if (this.prev < self.finallyLoc) {
                      return next(self.finallyLoc);
                    }
                  }
                }
              }
            }
          },
          /**
           * @param {string} type
           * @param {?} arg
           * @return {?}
           */
          abrupt : function(type, arg) {
            /** @type {number} */
            var unlock = this.tryEntries.length - 1;
            for (;unlock >= 0;--unlock) {
              var cache = this.tryEntries[unlock];
              if (cache.tryLoc <= this.prev && (hasOwnProperty.call(cache, "finallyLoc") && this.prev < cache.finallyLoc)) {
                var item = cache;
                break;
              }
            }
            if (item) {
              if ("break" === type || "continue" === type) {
                if (item.tryLoc <= arg) {
                  if (arg <= item.finallyLoc) {
                    /** @type {null} */
                    item = null;
                  }
                }
              }
            }
            var data = item ? item.completion : {};
            return data.type = type, data.arg = arg, item ? (this.method = "next", this.next = item.finallyLoc, value) : this.complete(data);
          },
          /**
           * @param {?} data
           * @param {Object} next
           * @return {?}
           */
          complete : function(data, next) {
            if ("throw" === data.type) {
              throw data.arg;
            }
            return "break" === data.type || "continue" === data.type ? this.next = data.arg : "return" === data.type ? (this.rval = this.arg = data.arg, this.method = "return", this.next = "end") : "normal" === data.type && (next && (this.next = next)), value;
          },
          /**
           * @param {?} onComplete
           * @return {?}
           */
          finish : function(onComplete) {
            /** @type {number} */
            var unlock = this.tryEntries.length - 1;
            for (;unlock >= 0;--unlock) {
              var httpServer = this.tryEntries[unlock];
              if (httpServer.finallyLoc === onComplete) {
                return this.complete(httpServer.completion, httpServer.afterLoc), start(httpServer), value;
              }
            }
          },
          /**
           * @param {Function} opt_attributes
           * @return {?}
           */
          catch : function(opt_attributes) {
            /** @type {number} */
            var unlock = this.tryEntries.length - 1;
            for (;unlock >= 0;--unlock) {
              var httpServer = this.tryEntries[unlock];
              if (httpServer.tryLoc === opt_attributes) {
                var me = httpServer.completion;
                if ("throw" === me.type) {
                  var arg = me.arg;
                  start(httpServer);
                }
                return arg;
              }
            }
            throw new Error("illegal catch attempt");
          },
          /**
           * @param {Object} isXML
           * @param {string} dataAndEvents
           * @param {Function} deepDataAndEvents
           * @return {?}
           */
          delegateYield : function(isXML, dataAndEvents, deepDataAndEvents) {
            return this.delegate = {
              iterator : val(isXML),
              resultName : dataAndEvents,
              /** @type {Function} */
              nextLoc : deepDataAndEvents
            }, "next" === this.method && (this.arg = void 0), value;
          }
        }, self;
      }(module.exports);
      try {
        regeneratorRuntime = html5;
      } catch (t) {
        Function("r", "regeneratorRuntime = r")(html5);
      }
    }, function(module, dataAndEvents, point) {
      /** @type {string} */
      module.exports = point.p + "0.srp6a-routines.worker.js";
    }]);
  });
  