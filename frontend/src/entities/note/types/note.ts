import { Workspace } from 'shared/model';

export type Note = {
  id: string;
  title: string;
  workspace: Workspace;
  content_uri: string;
};
