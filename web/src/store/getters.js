const getters = {
  token: state => state.stu.token,
  stuId: state => state.stu.stuId,
  stuNum: state => state.stu.stuNum,
  timetable: state => state.tt.timetable,
  courses: state => state.tt.courses,
  semesterStartTime: state => state.tt.semesterStartTime,
  ver: state => state.version.ver
}

export default getters
