import { Workspace } from 'shared/model';

export type CardEntity = {
  id: string;
  answer: string;
  question: string;
  workspace: Workspace;
};
