import { createBrowserRouter } from 'react-router-dom';
import { CollectionPage } from 'pages/collection';
import { KnowledgeBasePage } from 'pages/knowledge-base';
import { MainPage } from 'pages/main';
import { TopicPage } from 'pages/topic';

export enum Routes {
  Main = '/',
  KnowledgeBase = '/knowledge-base/:knowledgeBaseId',
  Collection = 'collection/:collectionId',
  Topic = 'topic/:topicNumber',
}

export const router = createBrowserRouter([
  {
    path: Routes.Main,
    element: <MainPage />,
  },
  {
    path: Routes.KnowledgeBase,
    element: <KnowledgeBasePage />,
    children: [
      {
        path: Routes.Collection,
        element: <CollectionPage />,
        children: [
          {
            path: Routes.Topic,
            element: <TopicPage />,
          },
        ],
      },
    ],
  },
]);
