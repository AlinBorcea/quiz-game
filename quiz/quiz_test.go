package quiz

import "testing"

func Test_recordFromRaw_EmptyRecord(t *testing.T) {
	var rec []string
	if _, err := recordFromRaw(rec); err == nil {
		t.FailNow()
	}
}

func Test_recordFromRaw_NoCorrect(t *testing.T) {
	rec := []string{"QWe", "das,ads", ""}
	if _, err := recordFromRaw(rec); err == nil {
		t.FailNow()
	}
}

func Test_recordFromRaw_CorrectIsChar(t *testing.T) {
	rec := []string{"QWe", "das,ads", "g"}
	if _, err := recordFromRaw(rec); err == nil {
		t.FailNow()
	}
}

func Test_recordFromRaw_CorrectIsString(t *testing.T) {
	rec := []string{"QWe", "das,ads", "ghio"}
	if _, err := recordFromRaw(rec); err == nil {
		t.FailNow()
	}
}

func Test_recordFromRaw_CorrectIsSpecialChars(t *testing.T) {
	rec := []string{"QWe", "das,ads", "\n\r\t"}
	if _, err := recordFromRaw(rec); err == nil {
		t.FailNow()
	}
}

func Test_recordFromRaw_Good(t *testing.T) {
	rec := []string{"QWe", "das,ads", "2"}
	if _, err := recordFromRaw(rec); err != nil {
		t.FailNow()
	}
}

func Test_questionFromRawRecord_Empty(t *testing.T) {
	rec := []string{}
	if _, err := questionFromRawRecord(rec); err == nil {
		t.FailNow()
	}
}

func Test_questionFromRawRecord_EmptyQue(t *testing.T) {
	rec := []string{"", "ads,asd"}
	if _, err := questionFromRawRecord(rec); err == nil {
		t.FailNow()
	}
}

func Test_questionFromRawRecord_EmptyAnswers(t *testing.T) {
	rec := []string{"que", ""}
	if _, err := questionFromRawRecord(rec); err == nil {
		t.FailNow()
	}
}

func Test_questionFromRawRecord_AnswersNoComma(t *testing.T) {
	rec := []string{"que", "1234"}
	if _, err := questionFromRawRecord(rec); err != nil {
		t.FailNow()
	}
}

func Test_questionFromRawRecord_AnswersOnlyCommas(t *testing.T) {
	rec := []string{"que", ",,,"}
	if _, err := questionFromRawRecord(rec); err == nil {
		t.FailNow()
	}
}

func TestHasEmptyAnswer_Empty(t *testing.T) {
	var answers []string

	if !hasEmptyAnswer(answers) {
		t.FailNow()
	}
}

func TestHasEmptyAnswer_1Empty(t *testing.T) {
	answers := []string{"1", "", "dsa"}
	if !hasEmptyAnswer(answers) {
		t.FailNow()
	}
}

func TestHasEmptyAnswer_AllEmpty(t *testing.T) {
	answers := []string{"", "", ""}
	if !hasEmptyAnswer(answers) {
		t.FailNow()
	}
}
