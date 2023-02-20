import { ModalForm, PageContainer, ProColumns, ProTable, ProFormText, ProFormSelect } from '@ant-design/pro-components';
import React from 'react';
import { userList, roleList, userAdd, userDel } from '@/services/tapi/api';
import { Button, message } from 'antd';
import { history } from '@umijs/max';
// import styles from './index.less';

// 编辑用户
const UserEdit: React.FC<TAPI.User & { title: string, handleRersh: any }> = (props) => {
  console.log(props)
  const { title, handleRersh } = props
  return (
    <ModalForm
      title={title}
      trigger={<Button type="primary">{title}</Button>}
      onFinish={async (values) => {
        let param: TAPI.UserAddParams = {
          id: values.id || 0,
          name: values.name,
          password: values.password || '',
          role_id: values.role_id,
        }
        let res: TAPI.UserAddResponse = await userAdd(param);
        if (res.code === 200) {
          message.success(title + "成功")
          handleRersh();
          return true;
        }
        message.error(title + "失败")
        return false;
      }}
    >
      <ProFormText
        name="id"
        hidden={true}
        initialValue={props.id || 0}
      />
      <ProFormText
        width="md"
        name="name"
        label="用户名"
        initialValue={props.name || ""}
      />
      <ProFormText
        width="md"
        name="password"
        label="密码"
        placeholder="密码，如果不填入那么就不修改密码"
      />
      <ProFormSelect
        name="role_id"
        label="角色"
        // 选中
        initialValue={props.role_id || 0}
        request={async () => {
          const list = await roleList();
          let data: Array<{ label: string | undefined, value: number | undefined }> = [];
          list.data?.map(vlaue => {
            data.push({
              "label": vlaue.name,
              "value": vlaue.id
            })
            return vlaue
          });
          return data;
        }}
      />
    </ModalForm>
  )
}

const UserList: React.FC = () => {
  const handleRersh = () => {
    history.go(0);
  }
  const handleDel = async (id: number | undefined) => {
    let params = {
      id: id,
    };
    let data = await userDel(params)
    if (data?.code === 200) {
      message.info("删除成功");
    } else {
      message.info("删除失败");
    }
    handleRersh();
  }

  const columns: ProColumns<TAPI.User>[] = [
    {
      title: '用户id',
      dataIndex: 'id',
      hideInSearch: true,
    },
    {
      title: '用户名称',
      dataIndex: 'name',
    },
    {
      title: '角色id',
      dataIndex: 'role_id',
      hideInSearch: true,
    },
    {
      title: '角色名称',
      dataIndex: 'role_name',
      hideInSearch: true,
    },
    {
      title: '更新时间',
      dataIndex: 'utime',
      hideInSearch: true,
    },
    {
      title: '创建时间',
      dataIndex: 'ctime',
      hideInSearch: true,
    },
    {
      title: "操作",
      hideInSearch: true,
      render: (text, record) => [
        <UserEdit key="edit" title="用户编辑" handleRersh={handleRersh} {...record} />,
        <Button key="2" type="primary" danger onClick={() => { handleDel(record.id) }}>删除</Button>,
      ],
    },
  ];




  return (
    <div>
      <PageContainer>
        <ProTable
          columns={columns}
          // 使用id作为key,这个key是字符串所以要转换一下
          rowKey={record => record.id?.toLocaleString() || ""}
          // 设置分页数据
          pagination={{
            pageSize: 20,
            onChange: (page) => console.log(page),
          }}
          // 数据获取，没有分页查询。分页是前端做的
          request={async (params) => {
            console.log(params)
            const UserListParams: TAPI.UserListParams = {
              name: params.name || '',
            };
            const d = await userList(UserListParams);
            return Promise.resolve({
              data: d.data,
              success: true,
            });
          }}
          toolBarRender={() => [
            <UserEdit key="edit" title="新增用户" handleRersh={handleRersh} />,
          ]}
        ></ProTable>

      </PageContainer>
    </div >
  );
}

export default UserList;
