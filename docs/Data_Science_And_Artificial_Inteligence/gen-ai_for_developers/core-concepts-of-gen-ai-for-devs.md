# Core Concepts of Gen AI for Developers

### Transformer

Neural network architecture designed to understand and generate sequences - like sentences, code, music, etc.

Unlike older models that process data step-by-step, Transformers look at the whole input at once.

### Self-attention

Mechanism that allows the model to decide which parts of the input to focus on, based on their relevance to each token
being generated.

### Positional encoding

Transformers don't naturally know what comes first, second, or last in a sentence.
They see all tokens at once.

Positional encoding solves this by adding a special pattern to each token based on its position.

## Modern Transformer

Modern transformer innovations:

- FlashAttention
  - This is the breakthrough that made modern models faster and more efficient
  - Longer context windows (up to 1M+ tokens)
  - Faster inference
  - Lower costs for deployments
- Grouped-query attention
  - multiple heads share keys and values
- Attention
- Mixture of Experts


