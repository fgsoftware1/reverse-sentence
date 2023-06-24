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

.. image:: https://app.fossa.com/api/projects/git%2Bgithub.com%2Ffgsoftware1%2Freverse-sentence.svg?type=large
   :alt: FOSSA Status
   :target: https://app.fossa.com/projects/git%2Bgithub.com%2Ffgsoftware1%2Freverse-sentence?ref=badge_large

