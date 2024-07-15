export interface INote {
  // TODO: Почему тут строковые литералы?
  id: 'string';
  title: 'string';
  workspace: 'string';
  content_uri: 'string';
}

// TODO: Почему не можем экспортировать type, а используем interface?
//  В чем отличие между type и interface?
// export type Note = {
//   id: 'string';
//   title: 'string';
//   workspace: 'string';
//   content_uri: 'string';
// }
