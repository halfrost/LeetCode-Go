'use strict';

(function () {
  const indexCfg = {{ with i18n "bookSearchConfig" }}
    {{ . }};
  {{ else }}
   {};
  {{ end }}

  indexCfg.doc = {
    id: 'id',
    field: ['title', 'content'],
    store: ['title', 'href', 'section'],
  };

  const index = FlexSearch.create('balance', indexCfg);
  window.bookSearchIndex = index;

  {{ range $index, $page := where .Site.Pages "Kind" "in" (slice "page" "section") }}
  {{ if $page.Content }}
  index.add({
    'id': {{ $index }},
    'href': '{{ $page.RelPermalink }}',
    'title': {{ (partial "docs/title" $page) | jsonify }},
    'section': {{ (partial "docs/title" $page.Parent) | jsonify }},
    'content': {{ $page.Plain | jsonify }}
  });
  {{- end -}}
  {{- end -}}
})();
