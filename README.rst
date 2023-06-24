reverse-sentence
===============

Reverses the words of a sentence.

Install
-------

.. code-block:: sh

   go get github.com/fgsoftware1/reverse-sentence

Example
-------

.. code-block:: go

   import(
      "github.com/fgsoftware1/reverse-sentence"
      ...
   )

.. code-block:: go

   func main() {
	   sentence := "Hello world!"
	   reversed := reverselib.Reverse(sentence)
	   fmt.Println(reversed) //world! Hello
   }

License
-------

ISC
