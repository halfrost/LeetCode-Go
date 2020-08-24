var prop = function (node, name, value) {
  if (typeof value === 'undefined') {
    return node[name];
  }
  node[name] = value;
};