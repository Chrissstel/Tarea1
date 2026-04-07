# Tarea1
implementación apropiada para las siguientes clases: STACK (lifo), QUEUE (fifo),
TABLE/HASH/DICTIONARY (order),..


TEST CASES DE STACK

TestStackPushPeek
- Insertar elementos en el stack.
- Verificar que Peek() regresa el último elemento insertado.

TestStackPop
- Insertar múltiples elementos.
- Verificar que Pop() regresa el elemento insertado más recientemente (LIFO).

TestStackIsEmpty
- Verificar que un stack nuevo esté vacío.
- Verificar que después de insertar elementos ya no esté vacío.

TestPopEmptyStack
- Intentar hacer Pop en un stack vacío.
- Verificar que la operación falle de forma segura.


TEST CASES DE QUEUE

TestQueueEnqueuePeek
- Insertar elementos en la queue.
- Verificar que Peek() regresa el primer elemento insertado.

TestQueueDequeue
- Insertar múltiples elementos.
- Verificar que Dequeue() regresa el primer elemento insertado (FIFO).

TestQueueIsEmpty
- Verificar que una queue nueva esté vacía.
- Verificar que después de insertar elementos ya no esté vacía.

TestDequeueEmptyQueue
- Intentar hacer Dequeue en una queue vacía.
- Verificar que la operación falle de forma segura.


TEST CASES DE DICTIONARY

TestDictionaryPutGet
- Insertar elementos en el dictionary usando Put.
- Verificar que Get() regrese el valor correcto asociado a la clave.

TestDictionaryUpdateValue
- Insertar un par clave-valor en el dictionary.
- Insertar la misma clave nuevamente con un valor diferente.
- Verificar que el valor asociado a la clave se actualice correctamente.

TestDictionaryRemove
- Insertar múltiples pares clave-valor en el dictionary.
- Eliminar una clave usando Remove.
- Verificar que la clave ya no exista en el dictionary.

TestDictionaryContains
- Insertar elementos en el dictionary.
- Verificar que Contains() regrese true para claves existentes y false para claves que no existen.

TestDictionarySize
- Insertar múltiples elementos en el dictionary.
- Verificar que Size() regrese el número correcto de elementos almacenados en el dictionary.