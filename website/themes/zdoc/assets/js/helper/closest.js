var closest = function(node, selector) {
  return (node.closest || function (_selector) {
    do {
      if ((node.matches || node.msMatchesSelector).call(node, _selector)) {
        return node;
      }
      node = node.parentElement || node.parentNode;
    } while (node !== null && node.nodeType === 1);

    return null;
  }).call(node, selector);
}