console.log(typeof 1); // number
console.log(typeof 'hello'); // string
console.log(typeof false); // boolean
console.log(typeof {}); // object
console.log(typeof []); // object
console.log(typeof null); // object
console.log(typeof undefined); // undefined
console.log(typeof function () {
}); // function

if (typeof Object.beget !== 'function') {
    Object.beget = function (o) {
        var F = function () {
        };
        F.prototype = o;
        return new F();
    }
}

var stooge = {
    "foo": "bar"
};

var anther_stooge = Object.beget(stooge);

anther_stooge.foo = "bar1";
console.log(anther_stooge.foo); // bar1

delete anther_stooge.foo;
console.log(anther_stooge.foo); // bar

stooge.bar = "foo";
console.log(anther_stooge.bar); // foo

console.log(anther_stooge.hasOwnProperty('foo'));
console.log(anther_stooge.hasOwnProperty('bar'));

for (name in anther_stooge) {
    console.log(name);
}
