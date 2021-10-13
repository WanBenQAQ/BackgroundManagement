import Vue from 'vue'
import {
  Message,
  MessageBox,
  Button,
  Container,
  Header,
  Aside,
  Main,
  Menu,
  Submenu,
  MenuItemGroup,
  MenuItem,
  Card,
  Input,
  Row,
  Col,
  TableColumn,
  Table,
  Form,
  FormItem,
  Switch,
  Pagination,
  Dialog,
  Cascader,
  DatePicker,
  Select,
  Option,
  Loading,
  Image,
  RadioGroup,
  RadioButton,
  Tooltip,
  Calendar,
  Steps,
  Step
} from 'element-ui'

Vue.use(Button)
Vue.use(Container)
Vue.use(Header)
Vue.use(Aside)
Vue.use(Main)
Vue.use(Menu)
Vue.use(Submenu)
Vue.use(MenuItemGroup)
Vue.use(MenuItem)
Vue.use(Card)
Vue.use(Input)
Vue.use(Row)
Vue.use(Col)
Vue.use(Table)
Vue.use(TableColumn)
Vue.use(Form)
Vue.use(FormItem)
Vue.use(Switch)
Vue.use(Pagination)
Vue.use(Dialog)
Vue.use(Cascader)
Vue.use(DatePicker)
Vue.use(Select)
Vue.use(Option)
Vue.use(Loading)
Vue.use(Image)
Vue.use(RadioGroup)
Vue.use(RadioButton)
Vue.use(Tooltip)
Vue.use(Calendar)
Vue.use(Steps)
Vue.use(Step)

//全局挂载
Vue.prototype.$message = Message
Vue.prototype.$confirm = MessageBox.confirm
