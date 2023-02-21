import { PageContainer, ProTable, ProColumns } from '@ant-design/pro-components';
import React from 'react';
import { Input, Button, message } from 'antd';
import { routerList, routerAdd } from '@/services/tapi/api';
import { history } from '@umijs/max';

const Index: React.FC = () => {

  const handleSave = async (data: any) => {
    const d = await routerAdd(data);
    message.info(d.msg);
    history.go(0);
  }

  const columns: ProColumns[] = [
    {
      title: "路由",
      dataIndex: "path",
    },
    {
      title: "路由名称",
      dataIndex: "name",
      render: (_, record) => [
        <Input key="routername" name="name" onChange={(e) => {
          record.name = e.target.value;
        }} />
      ]
    },
    {
      title: "操作",
      render: (text, record) => [
        <Button key="2" type="primary" onClick={() => { handleSave(record) }}>保存</Button>,
      ]
    },
  ];
  return (
    <PageContainer>
      <ProTable
        rowKey={record => record.path || ""}
        search={false}
        columns={columns}
        request={async (params) => {
          console.log(params);
          const d = await routerList()
          return Promise.resolve({
            data: d.data,
            success: true,
          });
        }}
      ></ProTable>
    </PageContainer>
  );
}

export default Index;