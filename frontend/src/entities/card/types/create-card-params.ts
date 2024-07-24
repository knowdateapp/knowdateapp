import { Workspace } from 'shared/model';

export type CreateCardParams = {
  workspace: Workspace;
  question: string;
  answer: string;
};
