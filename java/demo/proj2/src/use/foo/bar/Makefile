all:
	javac -d ../../../../../out -cp ../../../../../proj1/src usefoobar.java
	java -cp ../../../../../out use.foo.bar.usefoobar
jar:
	javac -d ../../../../../out -cp ../../../../../proj1/src usefoobar.java
	jar cvf demo.jar -C ../../../../../out .
	java -cp demo.jar use.foo.bar.usefoobar
clean:
	-rm -rf ../../../../../out
	-rm demo.jar