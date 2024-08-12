import { Workspace } from 'shared/model';

export type Card = {
  id: string;
  answer: string;
  question: string;
  workspace: Workspace;
};
