import React from 'react';
import Editor from 'rich-markdown-editor';

const exampleText = `
# Login User

Sebagai seorang user yang terdaftar saya dapat menggunakan username dan password untuk login ke dalam sistem

## Acceptance Criteria

- Greeek Sebagai seorang user yang terdaftar
- Greeek Sebagai seorang user yang terdaftar
- Greeek Sebagai seorang user yang terdaftar

1. Acceptance
2. Acceptance

## Linked Issue

- [ ] #XDM-100 Membuat initial database
- [x] #XDM-101 Membuat basis data skema
- [x] #XDM-102 Tidak melakukan migrasi data
`;

const AppEditor = (props) => (
  <div className="app-editor">
    <Editor classNames="app-editor" defaultValue={exampleText} spellCheck={false} />
  </div>
);

export default AppEditor;
