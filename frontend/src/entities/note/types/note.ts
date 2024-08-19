import { Workspace } from 'shared/types';

export type Note = {
  id: string;
  title: string;
  workspace: Workspace;
  content_uri: string;
};
