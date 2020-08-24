function next(node, selector) {
  if (selector && document.querySelector(selector) !== node.nextElementSibling) {
    return null;
  }

  return node.nextElementSibling;
}
