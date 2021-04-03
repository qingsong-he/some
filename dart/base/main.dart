import 'dart:isolate';
import 'dart:mirrors';
import 'dart:io';

void main() async {
  // case 1
  {
    // any type, but not type check in complie
    var v;
    v = 1;
    v = 1.1;

    // aways int, type check in complie
    var v1 = 1;

    // any type, type check in complie
    Object v2 = 1;
    v2 = 1.1;

    // any type, but not type check in complie
    dynamic v3 = 1;
    v3 = 1.1;

    // type check in complie
    final v4 = v2;
    const v5 = 1; // must constant expression

    // all object default value is null, so must init it before use

    const l1 = [1, 2, 3];
    const l2 = [1, 2, 3];
    print(identical(l1, l2)); // it's true, it's false when use final

    num v6;
    v6 = 1;
    v6 = 1.1;
  }

  // case 2
  {
    // assert just enable for product complie mode
    String v7 = 'hello world';
    v7 = "hello world";
    v7 = '''hello
  world''';
    print(v7); // use 'r' prefix diable escape

    String v8 = "hello hello";
    print(v8.replaceAll("hello", "-"));
  }

  // case 3
  {
    var v9 = DateTime.now();
    print([v9.timeZoneName, v9.timeZoneOffset, v9.isUtc, v9.hour]);

    var v10 = []; // any elem type
    v10.add("hello");
    v10.add(123);
    print([v10.first, v10.last, v10.length, v10.isEmpty, v10.reversed, v10]);

    // just only one elem type
    var v11 = List.filled(3, 0xff, growable: true);
    print(v11);
    v11.add(1);

    var v12 = Map();
    v12["foo"] = "bar";
    v12[1] = 2;
    print(v12);
    var v13 = {"foo": "bar"};
    v13["b"] = "b";
    print([
      v13.isEmpty,
      v13.length,
      v13.keys,
      v13.values,
      v13.containsKey("a"),
      v13.containsValue(123)
    ]);
    v13.updateAll((key, value) => "--$value--");
    print(v13);
    v13.removeWhere((key, value) => key == "foo");

    var v14 = Set();
    v14.add("foo");
    v14.add(13);
    v14.addAll([1, 2, 3, "foo", "bar"]);
    print([v14, v14.isEmpty]);
  }

  // case 4
  {
    String v15 = "😀";
    print([v15.length, v15.runes.length]);

    var v16 = Runes("\u{1f596} \u6211");
    var v17 = String.fromCharCodes(v16);
    print([
      v17,
      v17.codeUnits,
      v17.codeUnitAt(0),
      v17.codeUnitAt(1),
      v17.codeUnitAt(2),
      v17.runes
    ]);
  }

  // case 5
  {
    // Symbol
    currentMirrorSystem()
        .findLibrary(Symbol("dart.core"))
        .declarations
        .forEach((key, value) {
      print('$key - $value');
    });

    // Enum
    var v18 = Status.Runing;
    if (v18 == Status.Runing) {
      print("running");
    }

    // func
    print(add(1, 2, 3));
    print(add1(z: 1, x: 1));
    int add2(int x) {
      return x + 1;
    }

    print(add2(1));

    Function add3(int x) {
      return (int x1) => x + x1;
    }

    print(add3(1)(1));
  }

  // case 6
  {
    num v19 = 1;
    int v20 = v19 as int;
    print(v20);
    print(v20 is int);
    print(v20 is num);
    print(v20 is! num);
    print([true == true ? 'yes' : 'no', null ?? false]);

    bool v21;
    v21 = true;
    v21 ??= false;
    print(v21);

    var v22 = StringBuffer();
    v22..write("hello")..write("world");
    print(v22);

    for (var i = 0; i < 0x9; i++) {
      if (i % 2 == 0) {
        continue;
      }
      print(i);
    }
    while (true) {
      print(0);
      break;
    }
    do {
      print(0);
      break;
    } while (true);

    var v23 = 2;
    switch (v23) {
      case 1:
        print("case 1");
        break;
      case 2:
        print("case 2");
        continue mark1;
      mark1:
      default:
        print("not found");
    }

    try {
      throw OutOfMemoryError();
    } on OutOfMemoryError {
      print("no Mem");
      // rethrow;
    } catch (e) {
      print(e);
    } finally {
      print("finally");
    }
  }

  // case 7
  {
    var v24 = Point(1, 2);
    print([v24.x, v24.y, v24.yes]);
    var v25 = Point.fromJson({"x": 1, "y": 2, "yes": 3});
    print([v25.x, v25.y, v25.yes]);

    var v26 = People();
    v26.pName = "foobar";
    print([
      v26.pName,
      People.age,
    ]);
    People.showAge();

    UsePeople1().printName();
    UsePeople2().printName();
    AndroidPhone(111)
      ..startup()
      ..shutdown();

    AndroidPhonePlus(222)
      ..startup()
      ..shutdown();

    c();
    fooFactory(1).msg();
    fooFactory(2).msg();

    bar(1).msg();
    bar(2).msg();
  }

  // case 8
  {
    // import '...' show Foobar
    // import '...' hide Foobar
    // import '...' as otherName
    // import '...' deferred as otherName
  }

  // case 9
  {
    var l1 = <String>[];
    l1.add("foo");

    var s1 = <String>{};
    s1.add("foo");

    var m1 = Map<String, int>();
    m1["foo"] = 1;

    print([l1, s1, m1]);

    k addCache<k, v>(k key, v val) {
      print([key, val]);
      return key;
    }

    addCache(1, 2);
    addCache(1.1, 2.2);
    addCache<String, num>("foo", 2.2);

    var c1 = Phone<AndroidPhonePlus>(AndroidPhonePlus(123));
    c1.foobar.startup();
  }

  // case 10
  if (false) {
    // ref https://zhuanlan.zhihu.com/p/83781258
    print("begin");
    Future<String> getNetworkData() {
      return Future<String>(() {
        sleep(Duration(seconds: 3));
        // throw Exception("EOF");
        return "network data";
      });
    }

    var future = getNetworkData();
    future.then((value) => print(value)).catchError((error) {
      print(error);
    });
    print(future);
    print("end");
  }

  // case 11
  if (false) {
    // ref https://zhuanlan.zhihu.com/p/83781258
    // await will return a Future immediately, then execute outer layer code
    print("begin1");
    Future<String> getNetworkData1() async {
      sleep(Duration(seconds: 3));
      return await "get data:" + "network data1";
    }

    Future<void> run1() async {
      print("run1.begin");
      print(await getNetworkData1());
      print("run1.end");
    }

    run1();
    print("end1");
  }

  // case 12
  if (false) {
    Iterable<int> func1(int n) sync* {
      print('begin');
      var k = 0;
      while (k < n) {
        yield k++; // step 2
      }
      print('end');
    }

    var it = func1(5).iterator;
    while (it.moveNext()) {
      // step 1
      print(it.current); // step 3
    }
  }

  // case 13
  if (false) {
    Stream<int> func1(int n) async* {
      print('begin');
      var k = 0;
      while (k < n) {
        yield k++;
      }
      print('end');
    }

    if (false) {
      func1(5).listen((event) {
        print(event);
      });
    }

    var sub1 = func1(5).listen(null);
    sub1.onData((data) {
      print(data);
      // sub1.pause();
    });
  }

  // case 14
  {
    Iterable<int> func1(int n) sync* {
      if (n > 0) {
        yield n; // step 3
        yield* func1(n - 1); // step 2
      }
    }

    var i1 = func1(5).iterator;
    while (i1.moveNext()) {
      // step 1
      print(i1.current); // step 4
    }
  }

  // case 15
  {
    Foobar1()('hello');
  }

  // case 16
  {
    // ref https://ithelp.ithome.com.tw/users/20129053/ironman
    var mainPort = ReceivePort();
    var newIsolate = await Isolate.spawn(newIsolateMain, mainPort.sendPort);
    mainPort.listen((msg) {
      if (msg is SendPort) {
        int n = 42;
        msg.send(n);
        print("main Isolate: send int: $n");
      } else {
        print("main Isolate: recv int: $msg");
      }
    });
  }

  // case 17
  {
    var coll = SortedCollection(sort1);
    assert(coll.compare is Function);
    assert(coll.compare is Compare);
  }
}

typedef int Compare(Object a, Object b);

class SortedCollection {
  Compare compare;
  SortedCollection(this.compare);
}

int sort1(Object a, Object b) => 0;

void newIsolateMain(SendPort mainSendPort) {
  var newPort = ReceivePort();
  mainSendPort.send(newPort.sendPort);
  newPort.listen((msg) async {
    print("sub Isolate: recv msg: $msg");
    if (msg is int) {
      final value = await slowPlusOne(msg);
      mainSendPort.send(value);
      print("sub Isolate: send msg: $value");
    }
  });
}

Future<int> slowPlusOne(int n) =>
    Future.delayed(Duration(seconds: 5), () => n + 1);

class Foobar1 {
  call(String name) {
    print("name is $name");
  }
}

class Phone<T extends AndroidPhone> {
  final T foobar;
  Phone(this.foobar);
}

abstract class bar {
  void msg() {
    print("bar.func1");
  }

  factory bar(int t) {
    switch (t) {
      case 1:
        return bar1();
      case 2:
        return bar2();
    }
    return bar2();
  }
}

class bar1 implements bar {
  @override
  void msg() {
    print("bar1.func1");
  }
}

class bar2 implements bar {
  @override
  void msg() {
    print("bar2.func1");
  }
}

foo fooFactory(int t) {
  switch (t) {
    case 1:
      return foo1();
    case 2:
      return foo2();
  }
  return foo();
}

class foo {
  void msg() {
    print("foo.msg");
  }
}

class foo1 extends foo {
  @override
  void msg() {
    print("foo1.msg");
  }
}

class foo2 extends foo {
  @override
  void msg() {
    print("foo2.msg");
  }
}

class a {
  void func1() {
    print("a.func1");
  }
}

class b {
  void func1() {
    print("b.func1");
  }
}

class c extends a with b {
  c() {
    super.func1(); // b.fuc1
  }
}

class AndroidPhonePlus extends AndroidPhone {
  AndroidPhonePlus(int num) : super(num);
  @override
  void startup() {
    print("startup by sub class");
    return super.startup();
  }
}

abstract class IPhone {
  void startup();
  void shutdown();
}

class AndroidPhone implements IPhone {
  int number;
  AndroidPhone(this.number);
  @override
  void startup() {
    print("startup");
  }

  @override
  void shutdown() {
    print("shutdown");
  }
}

abstract class People1 {
  static String name = "foobar";
  void printName() {
    print(name);
  }
}

class UsePeople1 extends People1 {}

class UsePeople2 implements People1 {
  @override
  void printName() {
    print('i am UsePeople2');
  }
}

class People {
  static int age = 22;
  static void showAge() {
    print(age);
  }

  // private member
  String _name = '';

  set pName(String v) {
    _name = v;
  }

  String get pName {
    return _name;
  }
}

// class
class Point {
  // can not modify
  // must be init by construct
  final num x;
  final num y;
  var yes;
  // Point(num x, num y) {
  //   this.x = x;
  //   this.y = y;
  // }
  Point(this.x, this.y) : yes = 3;
  // named construct
  Point.fromJson(Map json) : this(json['x'], json['y']);
  Point.fromXML(Map xml) : this(xml['x'], xml['y']);
}

// option param
int add(int x, [int y = 1, int z = 1]) {
  return x + y + z;
}

// named param
int add1({int x = 1, int y = 2, int z = 3}) {
  return x + y + z;
}

enum Status { None, Runing, Stopped, Paused }
