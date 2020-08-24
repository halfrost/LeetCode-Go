/**
 * Get all DOM element up the tree that contain a class, ID, or data attribute
 * @param  {Node} elem The base element
 * @param  {String} selector The class, id, data attribute, or tag to look for
 * @return {Array} Null if no match
 */
var getParents = function (elem, selector) {

  var parents = [];
  var firstChar;
  if (selector) {
    firstChar = selector.charAt(0);
  }

  // Get matches
  for (; elem && elem !== document; elem = elem.parentNode) {
    if (selector) {

      // If selector is a class
      if (firstChar === '.') {
        if (elem.classList.contains(selector.substr(1))) {
          parents.push(elem);
        }
      }

      // If selector is an ID
      if (firstChar === '#') {
        if (elem.id === selector.substr(1)) {
          parents.push(elem);
        }
      }

      // If selector is a data attribute
      if (firstChar === '[') {
        if (elem.hasAttribute(selector.substr(1, selector.length - 1))) {
          parents.push(elem);
        }
      }

      // If selector is a tag
      if (elem.tagName.toLowerCase() === selector) {
        parents.push(elem);
      }

    } else {
      parents.push(elem);
    }

  }

  // Return parents if any exist
  if (parents.length === 0) {
    return null;
  } else {
    return parents;
  }

};