var prev = function(node, selector) {
  if (selector && document.querySelector(selector) !== node.previousElementSibling) {
    return null;
  }

  return node.previousElementSibling;
}